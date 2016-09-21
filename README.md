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

