package utils

import "testing"

func LogResult(t *testing.T, result int) {
	t.Logf("Result: \x1b[34m%d\x1b[0m", result)
}

func LogResultBigInt(t *testing.T, result uint64) {
	t.Logf("Result: \x1b[34m%d\x1b[0m", result)
}
