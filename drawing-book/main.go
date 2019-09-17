package main

import "fmt"

/*
 * Complete the pageCount function below.
 */
func pageCount(numPages int32, page int32) int32 {
	var pageTurns int32
	// determine pageTurns from front
	fromFront := (page) / 2
	// page Turns from back
	fromBack := (numPages - page) / 2
	if fromFront < fromBack {
		pageTurns = fromFront
	} else {
		pageTurns = fromBack
	}
	// Special cases
	// if numPages is an even number and looking for the page before last page, only one turn is required
	if numPages%2 == 0 && page == numPages-1 && fromFront != 0 {
		pageTurns = 1
	}
	return pageTurns
}

func main() {
	// 6, 2 => 1
	pageTurns := pageCount(2, 1)
	fmt.Println(pageTurns)
}
