// Copyright 2022 IBM Inc. All rights reserved
// Copyright © 2014 Steve Francia <spf@spf13.com>
//
// SPDX-License-Identifier: Apache2.0
package fsgo

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLstatIfPossible(t *testing.T) {
	wd, _ := os.Getwd()
	defer func() {
		os.Chdir(wd)
	}()

	osFs := &OsFs{}

	workDir, err := TempDir(osFs, "", "fsgo-lstate")
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		osFs.RemoveAll(workDir)
	}()

	memWorkDir := "/lstate"

	memFs := NewMemMapFs()
	overlayFs1 := &CopyOnWriteFs{base: osFs, layer: memFs}
	overlayFs2 := &CopyOnWriteFs{base: memFs, layer: osFs}
	overlayFsMemOnly := &CopyOnWriteFs{base: memFs, layer: NewMemMapFs()}
	basePathFs := &BasePathFs{source: osFs, path: workDir}
	basePathFsMem := &BasePathFs{source: memFs, path: memWorkDir}
	roFs := &ReadOnlyFs{source: osFs}
	roFsMem := &ReadOnlyFs{source: memFs}

	pathFileMem := filepath.Join(memWorkDir, "fsgom.txt")

	WriteFile(osFs, filepath.Join(workDir, "fsgo.txt"), []byte("Hi, FsGo!"), 0777)
	WriteFile(memFs, filepath.Join(pathFileMem), []byte("Hi, FsGo!"), 0777)

	os.Chdir(workDir)
	if err := os.Symlink("fsgo.txt", "symfsgo.txt"); err != nil {
		t.Fatal(err)
	}

	pathFile := filepath.Join(workDir, "fsgo.txt")
	pathSymlink := filepath.Join(workDir, "symfsgo.txt")

	checkLstat := func(l Lstater, name string, shouldLstat bool) os.FileInfo {
		statFile, isLstat, err := l.LstatIfPossible(name)
		if err != nil {
			t.Fatalf("Lstat check failed: %s", err)
		}
		if isLstat != shouldLstat {
			t.Fatalf("Lstat status was %t for %s", isLstat, name)
		}
		return statFile
	}

	testLstat := func(l Lstater, pathFile, pathSymlink string) {
		shouldLstat := pathSymlink != ""
		statRegular := checkLstat(l, pathFile, shouldLstat)
		statSymlink := checkLstat(l, pathSymlink, shouldLstat)
		if statRegular == nil || statSymlink == nil {
			t.Fatal("got nil FileInfo")
		}

		symSym := statSymlink.Mode()&os.ModeSymlink == os.ModeSymlink
		if symSym == (pathSymlink == "") {
			t.Fatal("expected the FileInfo to describe the symlink")
		}

		_, _, err := l.LstatIfPossible("this-should-not-exist.txt")
		if err == nil || !os.IsNotExist(err) {
			t.Fatalf("expected file to not exist, got %s", err)
		}
	}

	testLstat(osFs, pathFile, pathSymlink)
	testLstat(overlayFs1, pathFile, pathSymlink)
	testLstat(overlayFs2, pathFile, pathSymlink)
	testLstat(basePathFs, "fsgo.txt", "symfsgo.txt")
	testLstat(overlayFsMemOnly, pathFileMem, "")
	testLstat(basePathFsMem, "fsgom.txt", "")
	testLstat(roFs, pathFile, pathSymlink)
	testLstat(roFsMem, pathFileMem, "")
}
