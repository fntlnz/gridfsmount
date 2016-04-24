package main

import (
	"bazil.org/fuse/fs"
	"github.com/fntlnz/gridfsmount/datastore"
)

type GridFSFuse struct {
	ds *datastore.GridFSDataStore
}

func NewGridFSFuse(ds *datastore.GridFSDataStore) *GridFSFuse {
	return &GridFSFuse{
		ds: ds,
	}
}

func (g *GridFSFuse) Root() (fs.Node, error) {
	return &Dir{
		ds: g.ds,
	}, nil
}
