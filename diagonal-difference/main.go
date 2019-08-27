package main

import (
	"fmt"
	"math"
)

/*
 * Complete the 'diagonalDifference' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY arr as parameter.
 */
func diagonalDifference(arr [][]int32) int32 {
	// Write your code here
	var sum1, sum2 int32
	arrLen := len(arr[0])

	for idx, innerArr := range arr {
		sum1 = sum1 + innerArr[idx]
		sum2 = sum2 + innerArr[arrLen-idx-1]
	}

	return int32(math.Abs(float64(sum1 - sum2)))
}

func main() {
	lines := [][]int32{{11, 2, 4}, {4, 5, 6}, {10, 8, -12}}
	diff := diagonalDifference(lines)
	fmt.Println(diff)
}
