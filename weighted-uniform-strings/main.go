package main

import "fmt"

// Complete the weightedUniformStrings function below.
func weightedUniformStrings(s string, queries []int32) []string {
	set := make(map[int32]struct{})
	ret := make([]string, len(queries))
	var prev string
	length := 0
	for _, ch := range s {
		weight := int(ch) - int('a') + 1
		chStr := string(ch)
		set[int32(weight)] = struct{}{}
		if chStr == prev {
			length++
			set[int32(length*weight)] = struct{}{}
		} else {
			prev = chStr
			length = 1
		}
	}
	for idx, q := range queries {
		if _, ok := set[q]; ok {
			ret[idx] = "Yes"
		} else {
			ret[idx] = "No"
		}
	}
	return ret
}

func main() {
	testStr := "abccddde"
	queries := []int32{6, 1, 3, 12, 5, 9, 10}
	ret := weightedUniformStrings(testStr, queries)
	fmt.Printf("%v", ret)
}
