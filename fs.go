// Copyright (c) 2023, The Goki Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build js

package fs

import (
	"context"
	"syscall/js"

	"github.com/hack-pad/hackpadfs"
	"github.com/hack-pad/hackpadfs/indexeddb"
)

// FS represents a filesystem that implements the Node.js fs API.
// It is backed by an IndexedDB-based storage mechanism.
type FS struct {
	*indexeddb.FS
}

// NewFS returns a new [FS]. Most code should use [Config] instead.
func NewFS() (*FS, error) {
	ifs, err := indexeddb.NewFS(context.TODO(), "fs", indexeddb.Options{})
	if err != nil {
		return nil, err
	}
	return &FS{ifs}, nil
}

func (fs *FS) Chmod(args []js.Value) (any, error) {
	return nil, fs.FS.Chmod(args[0].String(), hackpadfs.FileMode(args[0].Int()))
}

func (fs *FS) Chown(args []js.Value) (any, error) {
	return nil, hackpadfs.Chown(fs.FS, args[0].String(), args[1].Int(), args[2].Int())
}

func (fs *FS) Close(args []js.Value) (any, error) {
	return nil, nil // TODO
}

func (fs *FS) Fchmod(args []js.Value) (any, error) {
	return fs.Chmod(args) // TODO
}

func (fs *FS) Fchown(args []js.Value) (any, error) {
	return fs.Chown(args) // TODO
}

func (fs *FS) Fstat(args []js.Value) (any, error) {
	return fs.Stat(args) // TODO
}

func (fs *FS) Fsync(args []js.Value) (any, error) {
	return nil, nil // TODO
}

func (fs *FS) Ftruncate(args []js.Value) (any, error) {
	return fs.Stat(args) // TODO
}

func (fs *FS) Stat(args []js.Value) (any, error) {
	return fs.FS.Stat(args[0].String())
}

// func (fs *FS) Truncate(args []js.Value) (any, error) {
// 	return hackpadfs.TruncateFile(args[0].String())
// }
