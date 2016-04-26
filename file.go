package main

import (
	"bytes"
	"io"
	"os"

	"bazil.org/fuse"
	"github.com/Sirupsen/logrus"
	"github.com/fntlnz/gridfsmount/datastore"
	"github.com/fntlnz/gridfsmount/util"
	"golang.org/x/net/context"
)

type File struct {
	ds       *datastore.GridFSDataStore
	name     string
	tempfile string
	synced   bool
}

func NewFile(ds *datastore.GridFSDataStore, name string) *File {
	return &File{
		ds:     ds,
		name:   name,
		synced: true,
	}
}

func (f *File) Attr(ctx context.Context, a *fuse.Attr) error {

	file, err := f.ds.FindByName(f.name)

	if err != nil {
		return fuse.ENOENT
	}

	defer file.Close()

	a.Inode = 2
	a.Mode = 0444
	a.Size = uint64(file.Size())
	return nil
}

func (f *File) ReadAll(ctx context.Context) ([]byte, error) {
	file, err := f.ds.FindByName(f.name)

	if err != nil {
		return nil, fuse.ENOENT
	}

	defer file.Close()

	buf := new(bytes.Buffer)

	_, err = buf.ReadFrom(file)

	if err != nil {
		return nil, fuse.EIO
	}
	return buf.Bytes(), nil
}

func (f *File) Write(ctx context.Context, req *fuse.WriteRequest, resp *fuse.WriteResponse) error {
	if req.Offset == 0 {
		tempname, err := util.TempFileName()

		if err != nil {
			logrus.Errorf("An error occurred while creating a temporary filename for file %s", f.name)
			return fuse.EIO
		}
		f.tempfile = tempname
	}

	logrus.Infof("Writing %s to tempfile: %s", f.name, f.tempfile)
	file, err := os.OpenFile(f.tempfile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)

	if err != nil {
		logrus.Errorf("An error occurred while opening temporary file for writing: %s", err.Error())
		return fuse.EIO
	}

	defer file.Close()

	size, err := file.Write(req.Data)

	if err != nil {
		logrus.Errorf("An error occurred while appending data to temporary file: %s", err.Error())
		return fuse.EIO
	}

	resp.Size = size
	f.synced = false

	return nil
}

func (f *File) Flush(ctx context.Context, req *fuse.FlushRequest) error {
	logrus.Infof("Flushing file %s", f.name)

	if f.synced == true {
		return nil
	}

	defer os.Remove(f.tempfile)
	tempFile, err := os.Open(f.tempfile)

	if err != nil {
		logrus.Errorf("An error occurred while opening the temporary file: %s, %s", f.tempfile, err.Error())
		return fuse.EIO
	}

	file, err := f.ds.Create(f.name)

	if err != nil {
		logrus.Errorf("An error occurred while creating the file on GridFS: %s", err.Error())
		return fuse.EIO
	}

	defer file.Close()

	_, err = io.Copy(file, tempFile)

	if err != nil {
		logrus.Errorf("An error occurred while writing to GridFS: %s", err.Error())
		return fuse.EIO
	}

	f.synced = true

	return nil
}
