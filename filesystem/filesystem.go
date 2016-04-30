package filesystem

import (
	"bazil.org/fuse/fs"
	"github.com/fntlnz/gridfsmount/datastore"
)

type Filesystem struct {
	ds *datastore.GridFSDataStore
}

func NewFilesystem(ds *datastore.GridFSDataStore) *Filesystem {
	return &Filesystem{
		ds: ds,
	}
}

func (g *Filesystem) Root() (fs.Node, error) {
	return &Dir{
		ds: g.ds,
	}, nil
}
