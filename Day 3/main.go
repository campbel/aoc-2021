package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {

	data, err := ioutil.ReadFile("report.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")

	oxygen := process(0, lines, func(zeros, ones []string) []string {
		if len(zeros) > len(ones) {
			return zeros
		}
		return ones
	})
	oxygenRating, _ := strconv.ParseInt(oxygen, 2, 64)

	scrubber := process(0, lines, func(zeros, ones []string) []string {
		if len(zeros) > len(ones) {
			return ones
		}
		return zeros
	})
	scrubberRating, _ := strconv.ParseInt(scrubber, 2, 64)

	fmt.Println(oxygen, oxygenRating, scrubber, scrubberRating, oxygenRating*scrubberRating)
}

func process(i int, lines []string, compare func([]string, []string) []string) string {
	if len(lines) == 1 {
		return lines[0]
	}
	var (
		zeros []string
		ones  []string
	)
	for _, line := range lines {
		switch line[i] {
		case '0':
			zeros = append(zeros, line)
		case '1':
			ones = append(ones, line)
		}
	}
	return process(i+1, compare(zeros, ones), compare)
}
