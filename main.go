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
	numboards := 100
	if isTesting {
		filename = "test.txt"
		newline = []byte("\r\n")
		numboards = 3
	}

	data, err := os.ReadFile(filename)
	if nil != err {
		fmt.Fprintf(os.Stderr, "error reading from file: %v\n", err)
		return
	}
	lines := bytes.Split(data, newline)
	numbers := bytes.Split(lines[0], []byte(","))
	boards := make([][5]int, numboards)
	for iboard := 0; iboard < numboards; iboard++ {
		ioffset := (iboard * 6) + 3
		for iline := 0; iline < 5; iline++ {
			tokens := bytes.Split(lines[ioffset+iline], []byte(" "))
			icolumn := 0
			for itoken := 0; itoken < len(tokens); itoken++ {
				token := tokens[itoken]
				if 0 == len(token) {
					continue
				}
				num, err := strconv.Atoi(string(token))
				if err != nil {
					fmt.Printf("failed to parse %v: %v\n", token, err)
					return
				}
				if icolumn >= 5 {
					fmt.Println("failed to parse board: more than 5 tokens!")
                    return
				}
				boards[iboard][icolumn] = num
				icolumn++
			}
			if icolumn < 5 {
				fmt.Println("failed to parse board: less than 5 tokens!")
                return
			}
		}
	}
}
