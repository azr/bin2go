# bin2go [![Build Status](https://travis-ci.org/azr/bin2go.svg?branch=master)](https://travis-ci.org/azr/bin2go)
--
    import "github.com/azr/bin2go"

Package bin2go transforms a binary stream to a golang byte array.

Embed binary files into your executables.

     func ExampleEchoBytes_empty() {
    	 b := bytes.NewBuffer([]byte{1})
    	 EchoByteSlice("echo", b, os.Stdout)
    	 // Output:
    	 // const echo = []byte{1}
     }

Runners:

    go get github.com/azr/bin2go/...

    cmd/bin2go/main.go binary_file name # to output your []byte to stdout

## Usage

#### func  EchoByteSlice

```go
func EchoByteSlice(name string, in io.Reader, out io.Writer) (err error)
```
EchoByteSlice reads byte by byte from *in* and creates a golang byte array named
*name*. It writes byte by byte to *out*.
