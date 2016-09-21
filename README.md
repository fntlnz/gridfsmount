# GridFS Mount

This is a tool to mount a GridFS collection as a [**Filesystem in Userspace**](https://en.wikipedia.org/wiki/Filesystem_in_Userspace) or as a **Docker volume**.

Please note that this tool is not complete yet and is subject to change.

**ANY** contribution is really appreciated.

## Features
- [ ] Docker volume driver plugin [WIP #1](https://github.com/fntlnz/gridfsmount/issues/1)
- [x] List files
- [ ] Support for directories (needs to be handled at an higher level because gridfs does not support directories)
- [x] Write files
- [ ] Overwrite files (was implemented, but does not work [#3](https://github.com/fntlnz/gridfsmount/issues/3))
- [ ] Remove files


# Usage

## Connect to a local MongoDB instance

This will connect to a local MongoDB instance listening on the default port and will mount the files under `/tmp/gridfs`

```
mkdir /tmp/gridfs
gridfsmount -db mygridfs -mountpoint /tmp/gridfs -debug -addr 127.0.0.1:27017
```

## Connect to a replicated cluster

If you have a replicated cluster you can provide multiple addresses on which to connect, for example

```
gridfsmount  -addr 10.0.20.10:27017 -addr 10.0.20.11:27017
```


## Need help?

```
gridfsmount -h
```


```
____ ____ _ ___  ____ ____ _  _ ____ _  _ _  _ ___
| __ |__/ | |  \ |___ [__  |\/| |  | |  | |\ |  |
|__] |  \ | |__/ |    ___] |  | |__| |__| | \|  |

(c) 2016 Lorenzo Fontana

version: 0.2.0
  -addr value
    	List of MongoDB database addresses
  -db string
    	The database to use to store the files collection (default "gridfsmount")
  -debug
    	Start in debug mode, provides a lot more information
  -gridfs-prefix string
    	The prefix that will be used by GridFS to create its collection (default "files")
  -mountpoint string
    	Filesystem mountpoint (default "/tmp/gridfs")
  -password string
    	Password to connect to the database
  -username string
    	Username to connect to the database

```
