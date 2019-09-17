package main

import "fmt"

// Complete the breakingRecords function below.
func breakingRecords(scores []int32) []int32 {
	// create return array
	records := []int32{0, 0} // high, low
	// set high and low scores as first index of array
	high := scores[0]
	low := scores[0]
	for _, score := range scores {
		if score > high {
			records[0]++
			high = score
		} else if score < low {
			records[1]++
			low = score
		}
	}
	return records
}

func main() {
	// 10 5 20 20 4 5 2 25 1
	// => 2, 4
	scores := []int32{10, 5, 20, 20, 4, 5, 2, 25, 1}
	records := breakingRecords(scores)
	fmt.Println(records)
}
