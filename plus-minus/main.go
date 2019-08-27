package main

import "fmt"

func plusMinus(arr []int32) {
	arrLen := float64(len(arr))
	var posCount, negCount, zeroCount int
	for _, val := range arr {
		switch {
		case val < 0:
			negCount++
		case val > 0:
			posCount++
		default:
			zeroCount++
		}
	}

	posPct := float64(posCount) / arrLen
	negPct := float64(negCount) / arrLen
	zeroPct := float64(zeroCount) / arrLen

	fmt.Printf("%f\n", posPct)
	fmt.Printf("%f\n", negPct)
	fmt.Printf("%f\n", zeroPct)
}

func main() {
	arr := []int32{-4, 3, -9, 0, 4, 1}
	plusMinus(arr)
}
