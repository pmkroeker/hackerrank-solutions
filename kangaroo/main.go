package main

import (
	"fmt"
	"math"
)

// Complete the kangaroo function below.
// YES when x1 + (count * v1) == x2 + (count * v2)
// count == (x1 - x2)/(v2 - v1)
func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	count := float64(x1-x2) / float64(v2-v1) // cannot use integer math, will check later for this to be an int
	if count < 0 {                           // if negative cannot be true
		return "NO"
	} else if math.IsInf(float64(count), 0) { // if Infinity will not ever happen
		return "NO"
	} else if count != float64(int32(count)) { // if count is not an int they will never meet
		return "NO"
	}
	return "YES"
}

func main() {
	resultOne := kangaroo(0, 3, 4, 2)     // Should output YES
	resultTwo := kangaroo(0, 2, 5, 3)     // Should output NO
	resultThree := kangaroo(21, 6, 47, 3) // NO
	fmt.Println(resultOne, resultTwo, resultThree)
}
