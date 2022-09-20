// Copyright 2022 IBM Inc. All rights reserved
// Copyright © 2014 Steve Francia <spf@spf13.com>
//
// SPDX-License-Identifier: Apache2.0
package fsgo

import (
	"errors"
)

// Symlinker is an optional interface in Fsgo. It is only implemented by the
// filesystems saying so.
// It indicates support for 3 symlink related interfaces that implement the
// behaviors of the os methods:
//   - Lstat
//   - Symlink, and
//   - Readlink
type Symlinker interface {
	Lstater
	Linker
	LinkReader
}

// Linker is an optional interface in Fsgo. It is only implemented by the
// filesystems saying so.
// It will call Symlink if the filesystem itself is, or it delegates to, the os filesystem,
// or the filesystem otherwise supports Symlink's.
type Linker interface {
	SymlinkIfPossible(oldname, newname string) error
}

// ErrNoSymlink is the error that will be wrapped in an os.LinkError if a file system
// does not support Symlink's either directly or through its delegated filesystem.
// As expressed by support for the Linker interface.
var ErrNoSymlink = errors.New("symlink not supported")

// LinkReader is an optional interface in Fsgo. It is only implemented by the
// filesystems saying so.
type LinkReader interface {
	ReadlinkIfPossible(name string) (string, error)
}

// ErrNoReadlink is the error that will be wrapped in an os.Path if a file system
// does not support the readlink operation either directly or through its delegated filesystem.
// As expressed by support for the LinkReader interface.
var ErrNoReadlink = errors.New("readlink not supported")
