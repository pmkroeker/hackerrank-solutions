package main

import "fmt"

// Complete the countingValleys function below.
func countingValleys(n int32, s string) int32 {
	pos := int32(0)   // start at position 0 (sealevel), negative is a valley, positive is a hill
	count := int32(0) // initialize valley count at 0
	for i := int32(0); i < n; i++ {
		char := s[i]     // char is a rune at this point
		prevPos := pos   // set previous position
		if char == 'U' { // if character is U rune, increment position
			pos++
		} else if char == 'D' { // else if character is D rune, decrement position
			pos += -1
		}
		// A valley is entered when the previous position was 0, and the current position is < 0 (or -1)
		if prevPos == 0 && pos < 0 {
			count++
		}
	}
	return count
}

func main() {
	stepCount := int32(8)
	position := "UDDDUDUU" // should return 1
	valleys := countingValleys(stepCount, position)
	fmt.Println(valleys)
}
