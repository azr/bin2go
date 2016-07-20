package bin2go

import (
	"bytes"
	"os"
)

func ExampleEchoBytes_empty() {
	b := bytes.NewBuffer([]byte{})
	EchoByteSlice("echo", b, os.Stdout)
	// Output:
	// const echo = []byte{}
}

func ExampleEchoBytes_one() {
	b := bytes.NewBuffer([]byte{1})
	EchoByteSlice("echo", b, os.Stdout)
	// Output:
	// const echo = []byte{1}
}

func ExampleEchoBytes_two() {
	b := bytes.NewBuffer([]byte{1, 2})
	EchoByteSlice("echo", b, os.Stdout)
	// Output:
	// const echo = []byte{1, 2}
}

func ExampleEchoBytes_three() {
	b := bytes.NewBuffer([]byte{1, 2, 3})
	EchoByteSlice("echo", b, os.Stdout)
	// Output:
	// const echo = []byte{1, 2, 3}
}

func ExampleEchoBytes_phrase() {
	b := bytes.NewBuffer([]byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88"))
	EchoByteSlice("echo", b, os.Stdout)
	// Output:
	// const echo = []byte{31, 139, 8, 0, 0, 9, 110, 136}
}
