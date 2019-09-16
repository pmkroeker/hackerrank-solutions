package main

import (
	"fmt"
)

// Complete the migratoryBirds function below.
func migratoryBirds(arr []int32) int32 {
	// birds can only be 1 - 5
	// initialize array where index is the bird type
	birds := []int{0, 0, 0, 0, 0, 0}
	// for each bird increment that birds count
	for _, val := range arr {
		birds[val]++
	}
	var common int32
	for idx, count := range birds {
		if count > birds[common] {
			common = int32(idx)
		}
	}
	return common
}

func main() {
	birds := []int32{1, 2, 3, 2}
	common := migratoryBirds(birds)
	fmt.Println(common)
}
