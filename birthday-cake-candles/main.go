package main

import "fmt"

// Complete the birthdayCakeCandles function below.
func birthdayCakeCandles(ar []int32) int32 {
	var totalOut int32
	maxVal := ar[0]
	// Get arr max val
	for _, val := range ar {
		if val > maxVal {
			maxVal = val
		}
	}
	// Count max vals
	for _, val := range ar {
		if val == maxVal {
			totalOut++
		}
	}
	return totalOut
}

func main() {
	candles := []int32{3, 2, 1, 3}
	outCount := birthdayCakeCandles(candles)
	fmt.Println(outCount)
}
