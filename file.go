package main

import (
	"bytes"

	"bazil.org/fuse"
	"github.com/fntlnz/gridfsmount/datastore"
	"golang.org/x/net/context"
)

type File struct {
	ds   *datastore.GridFSDataStore
	name string
}

const greeting = "hello, world\n"

func (f *File) Attr(ctx context.Context, a *fuse.Attr) error {
	a.Inode = 2
	a.Mode = 0444
	a.Size = uint64(len(greeting))
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
	//TODO(fntlnz): handle chunks, writing is not yet complete.
	file, _ := f.ds.Create(f.name)
	size, _ := file.Write(req.Data)

	defer file.Close()

	resp.Size = size

	return nil
}

func (f *File) Flush(ctx context.Context, req *fuse.FlushRequest) error {
	return nil
}
