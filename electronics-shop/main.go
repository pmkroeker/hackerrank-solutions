package main

import (
	"fmt"
	"sort"
)

/*
 * Complete the getMoneySpent function below.
 * returns the most money that could be spent
 * with 1 keyboard and 1 drive
 * return -1 if not enough money
 * I used int instead of int32
 */
func getMoneySpent(keyboards []int, drives []int, b int) int32 {
	spend := -1 // default to -1
	// Sort both arrays from smallest to largest
	sort.Ints(keyboards)
	sort.Ints(drives)
	for _, kValue := range keyboards {
		for _, dValue := range drives {
			total := kValue + dValue
			if total <= b && total > spend {
				spend = total
			}
		}
	}
	return int32(spend)
}

func main() {
	b := 10
	keyboards := []int{3, 1}
	drives := []int{5, 2, 8}
	spend := getMoneySpent(keyboards, drives, b)
	fmt.Println(spend) // should be 9
}

// 10 budget
// 3 1 keyboards
// 5 2 8 drives
