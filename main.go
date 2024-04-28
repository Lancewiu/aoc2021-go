package main

import (
	"bytes"
	"fmt"
	"os"
	"slices"
)

const isTesting = false

func main() {
	filename := "input.txt"
	newline := []byte("\n")
	numbits := 12
	if isTesting {
		filename = "test.txt"
		newline = []byte("\r\n")
		numbits = 5
	}

	data, err := os.ReadFile(filename)
	if nil != err {
		fmt.Fprintf(os.Stderr, "error reading from file: %v\n", err)
		return
	}
	lines := bytes.Split(data, newline)
	numlines := len(lines) - 1 // excluding file-end newline

	linemask := make([]bool, numlines)
	for iline := 0; iline < numlines; iline++ {
		linemask[iline] = true
	}
	for ibit := 0; ibit < numbits; ibit++ {
		var count int
		for imask := 0; imask < numlines; imask++ {
			if linemask[imask] {
				count++
			}
		}
		if 2 > count {
			break
		}
		majority := count / 2
		if 0 == count%2 {
			majority--
		}
		var onescount int
		for iline := 0; iline < numlines; iline++ {
			if linemask[iline] && '1' == lines[iline][ibit] {
				onescount++
			}
		}
		majoritybit := byte('0')
		if onescount > majority {
			majoritybit = byte('1')
		}
		for iline := 0; iline < numlines; iline++ {
			linemask[iline] = linemask[iline] && lines[iline][ibit] == majoritybit
		}
	}

	var o2rating int
	iO2RatingLine := slices.Index(linemask, true)
	if -1 == iO2RatingLine {
		println("oxygen rating filter failed! zero results!")
		return
	}
	o2RatingLine := lines[iO2RatingLine]
	for ibit := 0; ibit < numbits; ibit++ {
		if o2RatingLine[ibit] == byte('1') {
			o2rating |= 1 << ((numbits - 1) - ibit)
		}
	}

	for iline := 0; iline < numlines; iline++ {
		linemask[iline] = true
	}
	for ibit := 0; ibit < numbits; ibit++ {
		var count int
		for iline := 0; iline < numlines; iline++ {
			if linemask[iline] {
				count++
			}
		}
		if 2 > count {
			break
		}
		majority := count / 2
		if 0 == count%2 {
			majority--
		}
		var onescount int
		for iline := 0; iline < numlines; iline++ {
			if linemask[iline] && '1' == lines[iline][ibit] {
				onescount++
			}
		}
		minoritybit := byte('0')
		if onescount <= majority {
			minoritybit = byte('1')
		}
		for iline := 0; iline < numlines; iline++ {
			linemask[iline] = linemask[iline] && lines[iline][ibit] == minoritybit
		}
	}

	var co2rating int
	iCo2RatingLine := slices.Index(linemask, true)
	co2RatingLine := lines[iCo2RatingLine]
	for ibit := 0; ibit < numbits; ibit++ {
		if co2RatingLine[ibit] == byte('1') {
			co2rating |= 1 << ((numbits - 1) - ibit)
		}
	}

	fmt.Printf("life support rating (o2 x co2): %d\n", o2rating*co2rating)
}
