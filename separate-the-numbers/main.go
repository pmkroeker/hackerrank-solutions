package main

import (
	"fmt"
	"strconv"
)

func separateNumbers(s string) {
	maxSplit := len(s)
	for sp := 1; sp <= maxSplit; sp++ {
		isOK := false
		portions := split(s, sp)
		for idx, val := range portions {
			if idx == len(portions)-1 {
				break
			}
			if (val + 1) != portions[idx+1] {
				isOK = false
				break
			} else {
				isOK = true
			}
		}
		if isOK {
			fmt.Printf("YES %d", portions[0])
			return
		}
	}
	fmt.Println("NO")
}

// func testWithSplit(s string, splitSize int) (bool, error) {
// 	if len(s)%split != 0 {
// 		return nil, "cannot be split properly"
// 	}
// 	return false, split(s, splitSize)
// }

func split(s string, splitSize int) []int {
	var ret []int
	if len(s)%splitSize != 0 {
		return ret
	}
	for len(s) > 0 {
		portion, _ := strconv.ParseInt(s[:splitSize], 10, 0)
		s = s[splitSize:]
		ret = append(ret, int(portion))
	}
	return ret
}

func main() {
	testCases := []string{"99910001001", "7891011", "9899100", "999100010001", "010203"}
	for _, c := range testCases {
		separateNumbers(c)
	}
}
