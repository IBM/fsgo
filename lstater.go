// Copyright 2022 IBM Inc. All rights reserved
// Copyright © 2014 Steve Francia <spf@spf13.com>
//
// SPDX-License-Identifier: Apache2.0
package fsgo

import (
	"os"
)

// Lstater is an optional interface in Fsgo. It is only implemented by the
// filesystems saying so.
// It will call Lstat if the filesystem iself is, or it delegates to, the os filesystem.
// Else it will call Stat.
// In addtion to the FileInfo, it will return a boolean telling whether Lstat was called or not.
type Lstater interface {
	LstatIfPossible(name string) (os.FileInfo, bool, error)
}
