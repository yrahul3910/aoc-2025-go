package day07

import (
	"fmt"
	"strings"
	"yrahul3910/aoc-2025-go/utils"
)

func PrintSet(set map[int]bool, lower int, upper int) {
	arr := make([]int, 0)

	for i := lower; i < upper; i++ {
		if set[i] == true {
			arr = append(arr, i)
		}
	}

	utils.PrintArray(arr)
}

func GetPositions(line string, substr string) []int {
	pos := make([]int, 0)
	runningIndex := 0
	curString := strings.Clone(line)

	for strings.Contains(curString, substr) {
		curIndex := strings.Index(curString, substr)
		pos = append(pos, curIndex+runningIndex)
		runningIndex += curIndex + 1

		curString = curString[curIndex+1:]
	}

	return pos
}

func SolvePuzzle1(input string) int {
	lines := strings.Split(input, "\n")
	initialPos := strings.Index(lines[0], "S")
	nCols := len(lines[0])

	// go has no sets :)
	beams := map[int]bool{initialPos: true}
	splitCount := 0

	for _, line := range lines {
		splitters := GetPositions(line, "^")
		fmt.Print("splitters = ")
		utils.PrintArray(splitters)
		fmt.Print("beams = ")
		PrintSet(beams, 0, nCols)
		fmt.Println()

		for _, pos := range splitters {
			// is there a beam in this position (column)?
			if beams[pos] == true {
				splitCount++

				if pos > 0 {
					beams[pos-1] = true
				}
				if pos < nCols-1 {
					beams[pos+1] = true
				}
				delete(beams, pos)
			}
		}
	}

	return splitCount
}

func SolvePuzzle2(input string) int {
	// TODO: solve puzzle 2
	return 0
}
