package main

import (
	"fmt"
	"sort"
)

func miniMaxSum(arr []int) {
	sort.Ints(arr)
	minArr := arr[:4]
	maxArr := arr[1:]
	var min, max int
	for _, val := range minArr {
		min += val
	}
	for _, val := range maxArr {
		max += val
	}
	fmt.Printf("%d %d", min, max)
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	miniMaxSum(arr)
}
