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

	var prevsum int
	for iWin := 0; iWin < 3; iWin++ {
		num, err := strconv.Atoi(string(lines[iWin]))
		if nil != err {
			fmt.Fprintf(os.Stderr, "error parsing line: failed to parse first window: %v\n", err)
			return
		}
		prevsum += num
	}
	var count uint = 0
	for iLine := 3; iLine < len(lines); iLine++ {
		if 0 == len(lines[iLine]) {
			break
		}
		var sum int
		for iWin := 0; iWin < 3; iWin++ {
			l := lines[iLine-2+iWin]
			num, err := strconv.Atoi(string(l))
			if nil != err {
				fmt.Fprintf(os.Stderr, "error parsing line: failed to parse window: %v\n", err)
				return
			}
			sum += num
		}
		if prevsum < sum {
			count++
		}
		prevsum = sum
	}
	fmt.Printf("# of larger measurements: %d\n", count)
}
