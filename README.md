# bin2go [![Build Status](https://travis-ci.org/azr/bin2go.svg?branch=master)](https://travis-ci.org/azr/bin2go)
--
    import "github.com/azr/bin2go"

Package bin2go transforms a file or a binary stream to a golang byte array.

Embed binary files into your executables.

Runners:

    go get github.com/azr/bin2go/cmd/bin2go

    bin2go binary_file var_name # to output your []byte to stdout

## Usage

#### func  EchoByteSlice

```go
func EchoByteSlice(name string, vPerline int, in io.Reader, out io.Writer) (err error)
```
EchoByteSlice reads byte by byte from *in* and creates a golang byte array named
*name*. It writes byte by byte to *out*. if vPerline != 0, add a new line every
vPerline value
