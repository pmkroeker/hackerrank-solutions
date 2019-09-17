package main

import (
	"fmt"
)

// Complete the dayOfProgrammer function below.
// The day of the programmer is the 256th day of the year
// this function is not a real implementation solution
func dayOfProgrammer(year int32) string {
	if year == 1918 {
		return "26.09.1918"
	} else if ((year <= 1917) && (year%4 == 0)) || ((year > 1918) && (year%400 == 0 || ((year%4 == 0) && (year%100 != 0)))) {
		return fmt.Sprintf("12.09.%d", year)
	} else {
		return fmt.Sprintf("13.09.%d", year)
	}
}

func main() {
	date := dayOfProgrammer(2017) // 12.09.2016
	fmt.Println(date)
}
