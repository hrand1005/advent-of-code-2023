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

	grid := make([][]rune, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}

	n := len(grid)
	m := len(grid[0])

	for i := 0; i < n; i++ {
		j := 0
		for j < m {
			if unicode.IsDigit(grid[i][j]) {
				start := j
				for j < m && unicode.IsDigit(grid[i][j]) {
					j++
				}
				saveGearNumber(grid, i, start, j)
			} else {
				j++
			}
		}
	}

	fmt.Println(sumGearRatios())
}

type point struct {
	x, y int
}

var gearNumbers = make(map[point][]string)

func saveGearNumber(grid [][]rune, row, start, end int) {
	lo := max(start-1, 0)
	hi := min(end+1, len(grid[0]))

	nString := string(grid[row][start:end])
	for i := lo; i < hi; i++ {
		if row-1 >= 0 && grid[row-1][i] == '*' {
			p := point{row - 1, i}
			gearNumbers[p] = append(gearNumbers[p], nString)
		}
		if row+1 < len(grid) && grid[row+1][i] == '*' {
			p := point{row + 1, i}
			gearNumbers[p] = append(gearNumbers[p], nString)
		}
	}

	if grid[row][lo] == '*' {
		p := point{row, lo}
		gearNumbers[p] = append(gearNumbers[p], nString)
	}

	if grid[row][hi-1] == '*' {
		p := point{row, hi - 1}
		gearNumbers[p] = append(gearNumbers[p], nString)
	}
}

func sumGearRatios() int {
	total := 0
	for _, v := range gearNumbers {
		if len(v) == 2 {
			a, err := strconv.Atoi(v[0])
			if err != nil {
				panic(err)
			}
			b, err := strconv.Atoi(v[1])
			if err != nil {
				panic(err)
			}
			total += a * b
		}
	}
	return total
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
