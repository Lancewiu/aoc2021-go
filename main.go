package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

const isTesting = false

func main() {
	filename := "input.txt"
	newline := []byte("\n")
	if isTesting {
		filename = "test.txt"
		newline = []byte("\r\n")
	}

	data, err := os.ReadFile(filename)
	if nil != err {
		fmt.Fprintf(os.Stderr, "error reading from file: %v\n", err)
		return
	}
	lines := bytes.Split(data, newline)

	var pos, depth int
	for iLine := 0; iLine < len(lines); iLine++ {
        if 0 == len(lines[iLine]) {
            break
        }
		tokens := bytes.Split(lines[iLine], []byte(" "))
		offset, err := strconv.Atoi(string(tokens[1]))
		if nil != err {
			fmt.Fprintf(os.Stderr, "error while converting from offset: %v\n", err)
			return
		}
		switch string(tokens[0]) {
		case "forward":
			pos += offset
			continue
		case "down":
			depth += offset
			continue
		case "up":
			depth -= offset
			continue
		}
		fmt.Fprintf(
			os.Stderr,
			"error while parsing command string: unknown command string %s\n",
			tokens[0],
		)
		return
	}

	fmt.Printf("horizontal position x final depth: %d\n", pos*depth)
}
