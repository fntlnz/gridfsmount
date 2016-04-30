package main

import (
	"bazil.org/fuse"
	"bazil.org/fuse/fs"
	"github.com/fntlnz/gridfsmount/filesystem"
)

func MountAndServe(filesystem *filesystem.Filesystem, mountPoint string) error {
	c, err := fuse.Mount(
		mountPoint,
		fuse.FSName("gridfs"),
		fuse.LocalVolume(),
	)

	if err != nil {
		return err
	}

	defer c.Close()

	err = fs.Serve(c, filesystem)

	if err != nil {
		return err
	}

	<-c.Ready

	if err := c.MountError; err != nil {
		return err
	}
	return nil
}
