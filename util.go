package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func logError(err error) {
	fmt.Printf("error during processing: %v\n", err)
}

// Reads in a file from disk. The "f" parameter is a transformation function which will be executed against each line.
func readCsv(filename string, f func(string) error, errFunc func(error)) {
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
		if e = f(string(line)); e != nil {
			errFunc(e)
		}
	}
}

func readCsvBatch(filename, delim string, f func([]string) error, errFunc func(error)) {
	fd, e := os.Open(filename)
	if e != nil {
		fmt.Printf("failed to open file %s: %v\n", filename, e)
		return
	}
	defer fd.Close()
	fr := bufio.NewReader(fd)
	batch := make([]string, 0)
	for {
		lineBytes, _, e := fr.ReadLine()
		if e == io.EOF {
			if len(batch) > 0 {
				e = f(batch)
				if e != nil {
					errFunc(e)
				}
			}
			break
		}
		if e != nil {
			fmt.Printf("failed to read line: %v\n", e)
			continue
		}
		line := strings.TrimSpace(string(lineBytes))
		if line != delim {
			batch = append(batch, line)
		} else {
			if e = f(batch); e != nil {
				errFunc(e)
			}
			batch = make([]string, 0)
		}
	}
}
