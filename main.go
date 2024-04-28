package main

import (
	"bytes"
	"fmt"
	"os"
)

const isTesting = false

func main() {
	filename := "input.txt"
	newline := []byte("\n")
	numbits := 12
	bitmask := 0b111111111111
	if isTesting {
		filename = "test.txt"
		newline = []byte("\r\n")
		numbits = 5
		bitmask = 0b11111
	}

	data, err := os.ReadFile(filename)
	if nil != err {
		fmt.Fprintf(os.Stderr, "error reading from file: %v\n", err)
		return
	}
	lines := bytes.Split(data, newline)

	onecounts := make([]int, numbits)
	linecount := 0
	for iline := 0; iline < len(lines); iline++ {
		l := lines[iline]
		if 0 == len(l) {
			break
		}
		linecount++
		if numbits != len(l) {
			fmt.Fprintf(
				os.Stderr,
				"assumed all lines have %d characters. got %d characters instead\n",
				numbits,
				len(l),
			)
			return
		}
		for ibit := 0; ibit < numbits; ibit++ {
			if '1' == l[ibit] {
				onecounts[ibit]++
			}
		}
	}
	var gammarate int
	minority := linecount / 2
	for ibit := 0; ibit < numbits; ibit++ {
		if onecounts[ibit] > minority {
			gammarate |= 1 << ((numbits - 1) - ibit)
		}
	}
	epsilonrate := bitmask &^ gammarate
	fmt.Printf("power consumption (gamma rate x epsilon rate): %d\n", gammarate*epsilonrate)
}
