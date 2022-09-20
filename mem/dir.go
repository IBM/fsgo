// Copyright 2022 IBM Inc. All rights reserved
// Copyright © 2014 Steve Francia <spf@spf13.com>
//
// SPDX-License-Identifier: Apache2.0

package mem

type Dir interface {
	Len() int
	Names() []string
	Files() []*FileData
	Add(*FileData)
	Remove(*FileData)
}

func RemoveFromMemDir(dir *FileData, f *FileData) {
	dir.memDir.Remove(f)
}

func AddToMemDir(dir *FileData, f *FileData) {
	dir.memDir.Add(f)
}

func InitializeDir(d *FileData) {
	if d.memDir == nil {
		d.dir = true
		d.memDir = &DirMap{}
	}
}
