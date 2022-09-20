// Copyright 2022 IBM Inc. All rights reserved
// Copyright © 2014 Steve Francia <spf@spf13.com>.
//
// SPDX-License-Identifier: Apache2.0

package common

import "io/fs"

// FileInfoDirEntry provides an adapter from os.FileInfo to fs.DirEntry
type FileInfoDirEntry struct {
	fs.FileInfo
}

var _ fs.DirEntry = FileInfoDirEntry{}

func (d FileInfoDirEntry) Type() fs.FileMode { return d.FileInfo.Mode().Type() }

func (d FileInfoDirEntry) Info() (fs.FileInfo, error) { return d.FileInfo, nil }
