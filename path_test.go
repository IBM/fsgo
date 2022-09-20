// Copyright 2022 IBM Inc. All rights reserved
// Copyright © 2014 Steve Francia <spf@spf13.com>
//
// SPDX-License-Identifier: Apache2.0
package fsgo

import (
	"fmt"
	"os"
	"testing"
)

func TestWalk(t *testing.T) {
	defer removeAllTestFiles(t)
	var testDir string
	for i, fs := range Fss {
		if i == 0 {
			testDir = setupTestDirRoot(t, fs)
		} else {
			setupTestDirReusePath(t, fs, testDir)
		}
	}

	outputs := make([]string, len(Fss))
	for i, fs := range Fss {
		walkFn := func(path string, info os.FileInfo, err error) error {
			if err != nil {
				t.Error("walkFn err:", err)
			}
			var size int64
			if !info.IsDir() {
				size = info.Size()
			}
			outputs[i] += fmt.Sprintln(path, info.Name(), size, info.IsDir(), err)
			return nil
		}
		err := Walk(fs, testDir, walkFn)
		if err != nil {
			t.Error(err)
		}
	}
	fail := false
	for i, o := range outputs {
		if i == 0 {
			continue
		}
		if o != outputs[i-1] {
			fail = true
			break
		}
	}
	if fail {
		t.Log("Walk outputs not equal!")
		for i, o := range outputs {
			t.Log(Fss[i].Name() + "\n" + o)
		}
		t.Fail()
	}
}
