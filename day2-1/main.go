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
		game, data := parts[0], parts[1]

		if gamePossible(data) {
			id, err := strconv.Atoi(strings.Split(game, " ")[1])
			if err != nil {
				panic(err)
			}
			total += id
		}
	}

	fmt.Println(total)
}

func gamePossible(gData string) bool {
	sets := strings.Split(gData, ";")
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
				if count > 12 {
					return false
				}
			case "blue":
				if count > 14 {
					return false
				}
			case "green":
				if count > 13 {
					return false
				}
			default:
				panic("unkown option: " + pair[1])
			}
		}
	}
	return true
}
