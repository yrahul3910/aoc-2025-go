package day06_test

import (
	"testing"
	"yrahul3910/aoc-2025-go/day06"
	"yrahul3910/aoc-2025-go/utils"
)

func TestSolvePuzzle1(t *testing.T) {
	input := utils.ReadInput(t)
	result := day06.SolvePuzzle1(input)
	utils.LogResult(t, result)
}

func TestSolvePuzzle2(t *testing.T) {
	input := utils.ReadInput(t)
	result := day06.SolvePuzzle2(input)
	utils.LogResult(t, result)
}
