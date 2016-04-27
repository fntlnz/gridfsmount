package main

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/fntlnz/gridfsmount/datastore"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
	"golang.org/x/net/context"
)

type Dir struct {
	ds *datastore.GridFSDataStore
}

func (dir *Dir) Attr(ctx context.Context, a *fuse.Attr) error {
	a.Inode = 1
	a.Mode = os.ModeDir | 0555
	return nil
}

func (dir *Dir) Lookup(ctx context.Context, name string) (fs.Node, error) {
	file, err := dir.ds.FindByName(name)

	if err != nil {
		return nil, fuse.ENOENT
	}

	defer file.Close()

	node, err := NewFile(dir.ds, file.Name())

	if err != nil {
		logrus.Errorf("Error creating file: %s", err.Error())
		return nil, fuse.EIO
	}

	return node, nil
}

func (dir *Dir) ReadDirAll(ctx context.Context) ([]fuse.Dirent, error) {

	files, err := dir.ds.ListFileNames()
	if err != nil {
		return nil, fuse.ENOENT
	}

	var de []fuse.Dirent
	for _, file := range files {
		de = append(de, fuse.Dirent{
			Inode: 2,
			Name:  file,
			Type:  fuse.DT_File,
		})
	}
	return de, nil
}

func (dir *Dir) Create(ctx context.Context, req *fuse.CreateRequest, resp *fuse.CreateResponse) (fs.Node, fs.Handle, error) {

	file, err := dir.ds.Create(req.Name)

	if err != nil {
		logrus.Errorf("An error occurred creating file into the datastore: %s", err.Error())
		return nil, nil, fuse.EIO
	}

	defer file.Close()

	node, err := NewFile(dir.ds, file.Name())

	if err != nil {
		logrus.Errorf("An error occurred creating the file: %s", err.Error())
		return nil, nil, fuse.EIO
	}

	return node, node, nil

}
