package main

import (
	"flag"
	"fmt"
	"os"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
	"github.com/Sirupsen/logrus"
	"github.com/fntlnz/gridfsmount/datastore"
	"github.com/fntlnz/gridfsmount/util"
	"gopkg.in/mgo.v2"
)

const (
	BANNER = `____ ____ _ ___  ____ ____ _  _ ____ _  _ _  _ ___
| __ |__/ | |  \ |___ [__  |\/| |  | |  | |\ |  |
|__] |  \ | |__/ |    ___] |  | |__| |__| | \|  |

(c) 2016 Lorenzo Fontana

version: %s
`
	version = "0.1.0-dev"
)

var (
	mongoUri     string
	dbName       string
	gridFSPrefix string
	mountPoint   string
	debug        bool
)

func usage() {
	fmt.Fprintf(os.Stderr, BANNER, version)
	flag.PrintDefaults()
}

func init() {
	flag.StringVar(&mongoUri, "mongouri", "127.0.0.1:27017", "Mongo endpoint")
	flag.StringVar(&dbName, "db", "gridfsmount", "The database to use to store the files collection")
	flag.StringVar(&gridFSPrefix, "gridfs-prefix", "files", "The prefix that will be used by GridFS to create its collection")
	flag.StringVar(&mountPoint, "mountpoint", "/tmp/gridfs", "Filesystem mountpoint")
	flag.BoolVar(&debug, "debug", false, "Start in debug mode, provides a lot more information")
	flag.Usage = usage
	flag.Parse()
}

func main() {
	if debug {
		util.EnableDebug()
	}

	session, err := mgo.Dial(mongoUri)
	if err != nil {
		logrus.Fatal(err)
	}

	ds := datastore.NewGridFSDataStore(session, dbName, gridFSPrefix)
	defer ds.Close()

	gridFSFuse := NewGridFSFuse(ds)

	c, err := fuse.Mount(
		mountPoint,
		fuse.FSName("gridfs"),
		fuse.LocalVolume(),
	)

	if err != nil {
		logrus.Fatal(err)
	}

	defer c.Close()

	err = fs.Serve(c, gridFSFuse)

	if err != nil {
		logrus.Fatal(err)
	}

	<-c.Ready

	if err := c.MountError; err != nil {
		logrus.Fatal(err)
	}
}
