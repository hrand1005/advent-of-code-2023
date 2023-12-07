package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
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
		line := []rune(scanner.Text())
		var first, second rune
		for _, v := range line {
			if unicode.IsDigit(v) {
				first = v
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			if unicode.IsDigit(line[i]) {
				second = line[i]
				break
			}
		}

		num, err := strconv.Atoi(fmt.Sprintf("%c%c", first, second))
		if err != nil {
			panic(err)
		}
		total += num
	}

	fmt.Printf("%v\n", total)
}
