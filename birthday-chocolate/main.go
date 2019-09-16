package main

import "fmt"

// Complete the birthday function below.
// returns number of segments that satisfy the conditions
func birthday(s []int32, d int32, m int32) int32 {
	// init return count to 0
	countSegments := int32(0)
	// iterate over array
	for idx := range s {
		// get upper bound. if greater than length set to length
		upper := idx + int(m)
		if upper > len(s) {
			upper = len(s)
		}
		// create sub array of avaiable sections
		sub := s[idx:upper]
		if len(sub) != int(m) {
			break
		}
		// sum segements
		sum := int32(0)
		for _, val := range sub {
			sum += val
		}
		// if sum is equal to day segement is possible
		if sum == d {
			countSegments++
		}
	}
	return countSegments
}

func main() {
	bar := []int32{2, 5, 1, 3, 4, 4, 3, 5, 1, 1, 2, 1, 4, 1, 3, 3, 4, 2, 1} // bar
	day := int32(18)                                                        // day -> sum of segement sqaures
	month := int32(7)                                                       // month -> length of segment
	segments := birthday(bar, day, month)                                   // should => 3
	fmt.Println(segments)
}
