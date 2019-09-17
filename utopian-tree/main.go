package main

import "fmt"

// Complete the utopianTree function below.
func utopianTree(n int32) int32 {
	// starting height is one
	height := int32(1)
	// iterate over the seasons
	for i := int32(1); i <= n; i++ {
		// if season is even (summer) add one
		if i%2 == 0 {
			height++
		} else {
			// season is spring so height doubles
			height *= 2
		}
	}
	return height
}

func main() {
	one := utopianTree(0)   // 1
	two := utopianTree(1)   // 2
	three := utopianTree(4) // 7
	fmt.Println(one, two, three)
}
