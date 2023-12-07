package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
		parts := strings.Split(line, ":")
		total += gamePower(parts[1])
	}

	fmt.Println(total)
}

func gamePower(gData string) int {
	sets := strings.Split(gData, ";")
	var maxRed, maxBlue, maxGreen int
	for _, set := range sets {
		colors := strings.Split(set, ",")
		for _, c := range colors {
			c = strings.Trim(c, " ")
			pair := strings.Split(c, " ")
			count, err := strconv.Atoi(pair[0])
			if err != nil {
				panic(err)
			}
			switch pair[1] {
			case "red":
				maxRed = max(maxRed, count)
			case "green":
				maxGreen = max(maxGreen, count)
			case "blue":
				maxBlue = max(maxBlue, count)
			default:
				panic("unkown option: " + pair[1])
			}
		}
	}
	return maxRed * maxGreen * maxBlue
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
