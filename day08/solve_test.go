package day08_test

import (
	"testing"
	"yrahul3910/aoc-2025-go/day08"
	"yrahul3910/aoc-2025-go/utils"
)

func TestSolvePuzzle1(t *testing.T) {
	input := utils.ReadInput(t)
	result := day08.SolvePuzzle1(input)
	utils.LogResult(t, result)
}

func TestSolvePuzzle2(t *testing.T) {
	input := utils.ReadInput(t)
	result := day08.SolvePuzzle2(input)
	utils.LogResult(t, result)
}
