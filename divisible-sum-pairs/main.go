package main

import "fmt"

// Complete the divisibleSumPairs function below.
func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
	count := int32(0)
	for x := int32(0); x < n; x++ {
		for y := x + 1; y < n; y++ {
			if (ar[x]+ar[y])%k == 0 {
				count++
			}
		}
	}
	return count
}

func main() {
	// 6 3
	// 1 3 2 6 1 2
	// => 5
	ar := []int32{1, 3, 2, 6, 1, 2}
	dsp := divisibleSumPairs(6, 3, ar)
	fmt.Println(dsp)
}
