package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/azr/bin2go"
)

func main() {
	var Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s binary_file var_name\n", os.Args[0])
		flag.PrintDefaults()
	}
	perLine := flag.Int("per-line", 20, "max number of values per line. Set to zero to never add new line.")
	flag.Usage = Usage
	flag.Parse()
	argv := flag.Args()

	if len(argv) != 2 {
		flag.Usage()
		os.Exit(1)
	}

	file, err := os.Open(argv[0]) // For read access.
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open %s %s", argv[0], err)
		flag.Usage()
		return
	}

	err = bin2go.EchoByteSlice(argv[1], *perLine, bufio.NewReader(file), os.Stdout)
	if err != nil {
		fmt.Fprintf(os.Stderr, "EchoBytes failed: %s", err)
	}
}
