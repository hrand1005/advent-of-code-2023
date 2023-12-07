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

	total := 0
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
				if isPartNum(grid, i, start, j) {
					num, err := strconv.Atoi(string(grid[i][start:j]))
					if err != nil {
						panic(err)
					}
					total += num
				}
			} else {
				j++
			}
		}
	}

	fmt.Println(total)
}

func isPartNum(grid [][]rune, row, start, end int) bool {
	lo := max(start-1, 0)
	hi := min(end+1, len(grid[0]))

	for i := lo; i < hi; i++ {
		if row-1 >= 0 && isSymbol(grid, row-1, i) {
			return true
		}
		if row+1 < len(grid) && isSymbol(grid, row+1, i) {
			return true
		}
	}

	return isSymbol(grid, row, lo) || isSymbol(grid, row, hi-1)
}

func isSymbol(grid [][]rune, row, col int) bool {
	return !unicode.IsDigit(grid[row][col]) && grid[row][col] != '.'
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
