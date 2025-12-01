package day11_test

import (
	"testing"
	"yrahul3910/aoc-2025-go/day11"
	"yrahul3910/aoc-2025-go/utils"
)

func TestSolvePuzzle1(t *testing.T) {
	input := utils.ReadInput(t)
	result := day11.SolvePuzzle1(input)
	utils.LogResult(t, result)
}

func TestSolvePuzzle2(t *testing.T) {
	input := utils.ReadInput(t)
	result := day11.SolvePuzzle2(input)
	utils.LogResult(t, result)
}
