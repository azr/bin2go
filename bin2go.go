// Package bin2go transforms a binary stream to a golang byte array.
//
// Embed binary files into your executables.
//  func ExampleEchoBytes_empty() {
//	 b := bytes.NewBuffer([]byte{1})
//	 EchoByteSlice("echo", b, os.Stdout)
//	 // Output:
// 	 // var echo = []byte{1}
//  }
// Runners:
//
//   cmd/echo/main.go binary_file var_name # to output your []byte to stdout
//
package bin2go

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//EchoByteSlice reads byte by byte from *in* and creates a golang byte array named *name*.
//It writes byte by byte to *out*.
func EchoByteSlice(name string, in io.Reader, out io.Writer) (err error) {
	_, err = fmt.Fprintf(out, `var %s = []byte{`, name)
	if err != nil {
		return err
	}

	bfrom := bufio.NewReader(in)
	for n := 0; ; n++ {
		var rerr error
		c, rerr := bfrom.ReadByte()
		if rerr != nil {
			if rerr != io.EOF {
				fmt.Fprintf(os.Stderr, "ReadByte failed: %s ", rerr)
			}
			break
		}
		if n > 0 {
			_, err = fmt.Fprintf(out, `, `)
			if err != nil {
				return err
			}
		}
		_, err = fmt.Fprintf(out, `%d`, c)
		if err != nil {
			return err
		}
	}
	_, err = fmt.Fprintf(out, `}`)
	return err
}
