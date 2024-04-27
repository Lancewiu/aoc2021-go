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

	var count uint = 0
	prevnum, err := strconv.Atoi(string(lines[0]))
	if nil != err {
		fmt.Fprintf(os.Stderr, "error parsing line: failed to parse first line: %v\n", err)
		return
	}
	for i := 1; i < len(lines); i++ {
		if 0 == len(lines[i]) {
			break
		}
		num, err := strconv.Atoi(string(lines[i]))
		if nil != err {
			fmt.Fprintf(os.Stderr, "error parsing line: failed to parse line: %v\n", err)
			return
		}
		if prevnum < num {
			count++
		}
		prevnum = num
	}
	fmt.Printf("# of larger measurements: %d\n", count)
}
