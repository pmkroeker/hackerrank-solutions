package main

import "fmt"

// Complete the sockMerchant function below.
func sockMerchant(n int32, ar []int32) int32 {
	var pairs int32
	socksOfColor := make(map[int32]int32) // make a map of key: int32 value int
	for _, sock := range ar {
		socksOfColor[sock]++ // count each sock colour
	}
	for _, val := range socksOfColor {
		// a pair is 2 socks, and integer math rounds down
		// taking the count of each colour and dividing by 2 will give us the number of pairs in that colour
		// add that count to the total pairs
		colourPairs := val / 2
		pairs += colourPairs
	}
	return pairs
}

func main() {
	sockCount := int32(9)
	socks := []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}
	paired := sockMerchant(sockCount, socks)
	fmt.Println(paired) // should be 3
}
