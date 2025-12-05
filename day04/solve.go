package day04

import (
	"fmt"
	"strings"
)

type index struct {
	x int
	y int
}

func ValidateIndex(i int, j int, r int, c int) bool {
	return (i >= 0 && i < r && j >= 0 && j < c)
}

func Neighbors(i int, j int, r int, c int) []index {
	neighbors := make([]index, 0, 8) // Max 8 neighbors

	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}
			if ValidateIndex(i+dx, j+dy, r, c) {
				neighbors = append(neighbors, index{i + dx, j + dy})
			}
		}
	}

	return neighbors
}

func SolvePuzzle1(input string) int {
	lines := strings.Split(input, "\n")

	r := len(lines)
	c := len(lines[0])

	count := 0

	for i := range r {
		if len(strings.TrimSpace(lines[i])) == 0 {
			continue
		}

		for j := range c {
			if lines[i][j] != '@' {
				continue
			}

			neighbors := Neighbors(i, j, r, c)
			fmt.Printf("(%d, %d): %d\n", i, j, len(neighbors))

			curCount := 0
			for _, n := range neighbors {
				if len(strings.TrimSpace(lines[n.x])) == 0 {
					continue
				}

				if lines[n.x][n.y] == '@' {
					curCount++
				}
			}

			fmt.Printf("(%d, %d) -> %d\n", i, j, curCount)
			if curCount < 4 { // Problem constraint
				count++
			}
		}
	}

	return count
}

func SolvePuzzle2(input string) int {
	// TODO: solve puzzle 2
	return 0
}
