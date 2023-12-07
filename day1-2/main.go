package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	total := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		var first, second int
		for i := range line {
			if v, ok := toDigit(line[i:]); ok {
				first = v
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			if v, ok := toDigit(line[i:]); ok {
				second = v
				break
			}
		}

		total += (first*10 + second)
	}

	fmt.Printf("%d\n", total)
}

var digits = map[string]int{
	"1":     1,
	"one":   1,
	"2":     2,
	"two":   2,
	"3":     3,
	"three": 3,
	"4":     4,
	"four":  4,
	"5":     5,
	"five":  5,
	"6":     6,
	"six":   6,
	"7":     7,
	"seven": 7,
	"8":     8,
	"eight": 8,
	"9":     9,
	"nine":  9,
}

func toDigit(s string) (int, bool) {
	for k, v := range digits {
		if strings.HasPrefix(s, k) {
			return v, true
		}
	}
	return 0, false
}
