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
		cardData := strings.Split(line, ":")[1]
		cardNums := strings.Split(cardData, "|")
		winners, got := strings.Fields(cardNums[0]), strings.Fields(cardNums[1])

		want := make(map[string]bool)
		for _, w := range winners {
			want[w] = true
		}

		score := 0
		for _, g := range got {
			if want[g] {
				if score == 0 {
					score++
				} else {
					score *= 2
				}
			}
		}
		total += score
	}

	fmt.Println(total)
}
