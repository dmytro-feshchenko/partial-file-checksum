[![Build Status](https://travis-ci.org/technoboom/partial-file-checksum.svg?branch=master)](https://travis-ci.org/technoboom/partial-file-checksum)
# Utility for creating hash sum of given file
## The purpose of this utility:
When transferring files via Internet we need to take care about integrity of the files.
Common algorithms are not suitable for us, because they are commonly works with the whole file and we can't check the checksum for some part of loaded data.
## The core principle:
We split our file into chunks with fixed length (e.g. 1024KB). Starts from the last chunk we calculate his Sum256. The given hash needs to be added to the previous block (in hex representation). Do this for all other blocks.
The hash which will be in the end equals to the hash of the whole file. After this we can load the file which split into chunks and use this checksum for check in the boot process of the next block.
## Example:
To see an example run test command
```
go test
```
The result:
```
Source file length (bytes): 16927313
Count chunks: 16531
16531
03c08f4ee0b576fe319338139c045c89c3e8e9409633bea29442e21425006ea8
PASS
ok      github.com/technoboom/partial-files-checksum    0.113s
```
