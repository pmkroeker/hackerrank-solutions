package main

import (
	"fmt"
	"strings"
)

func staircase(n int32) {
	stair := make([]string, n)
	for idx := range stair {
		stair[idx] = " "
	}
	for i := int32(1); i <= n; i++ {
		stair[n-i] = "#"
		fmt.Println(strings.Join(stair, ""))
	}
}

func main() {
	staircase(6)
}
