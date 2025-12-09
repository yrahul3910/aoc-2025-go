package utils

import (
	"fmt"
)

func PrintArray[T any](arr []T) {
	fmt.Print("{")
	for idx, i := range arr {
		fmt.Print(i)

		if idx < len(arr)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Print("}")
}
