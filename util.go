package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func logError(err error) {
	fmt.Printf("error during processing: %v\n", err)
}

// Reads in a file from disk. The "f" parameter is a transformation function which will be executed against each line.
func readCsv(filename string, f func([]byte) error, errFunc func(error)) {
	fd, e := os.Open(filename)
	if e != nil {
		fmt.Printf("failed to open file %s: %v\n", filename, e)
		return
	}
	defer fd.Close()
	fr := bufio.NewReader(fd)
	for {
		line, _, e := fr.ReadLine()
		if e == io.EOF {
			break
		}
		if e != nil {
			fmt.Printf("failed to read line: %v\n", e)
			continue
		}
		if e = f(line); e != nil {
			errFunc(e)
		}
	}
}
