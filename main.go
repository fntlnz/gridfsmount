package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/fntlnz/gridfsmount/datastore"
	"github.com/fntlnz/gridfsmount/filesystem"
	"github.com/fntlnz/gridfsmount/util"
	"gopkg.in/mgo.v2"
)

const (
	banner = `____ ____ _ ___  ____ ____ _  _ ____ _  _ _  _ ___
| __ |__/ | |  \ |___ [__  |\/| |  | |  | |\ |  |
|__] |  \ | |__/ |    ___] |  | |__| |__| | \|  |

(c) 2016 Lorenzo Fontana

version: %s
`
	version = "0.1.0-dev"
)

var (
	mongoAddrs   util.ArrayFlags
	dbName       string
	dbUsername   string
	dbPassword   string
	gridFSPrefix string
	mountPoint   string
	debug        bool
)

func usage() {
	fmt.Fprintf(os.Stderr, banner, version)
	flag.PrintDefaults()
}

func init() {
	flag.Var(&mongoAddrs, "addr", "List of MongoDB database addresses")
	flag.StringVar(&dbName, "db", "gridfsmount", "The database to use to store the files collection")
	flag.StringVar(&dbUsername, "username", "", "Username to connect to the database")
	flag.StringVar(&dbPassword, "password", "", "Password to connect to the database")
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

	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    mongoAddrs,
		Timeout:  60 * time.Second,
		Database: dbName,
		Username: dbUsername,
		Password: dbPassword,
	}

	session, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		logrus.Fatal(err)
	}

	ds := datastore.NewGridFSDataStore(session, dbName, gridFSPrefix)
	defer ds.Close()

	fs := filesystem.NewFilesystem(ds)
	MountAndServe(fs, mountPoint)
}
