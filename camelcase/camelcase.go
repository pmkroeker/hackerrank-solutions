package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	var input string
	fmt.Scanf("%s\n", &input)
	fmt.Printf("The input word is: %s\n", input)
	countPeter := countCamel(input)
	fmt.Printf("Camel Case Count (Peter): %d\n", countPeter)
	countGopher := gopherCountCamel(input)
	fmt.Printf("Camel Case Count (Gopher): %d\n", countGopher)
}

func countCamel(s string) int32 {
	re := regexp.MustCompile(`[A-Z]`)
	camelCaseCount := re.FindAllString(s, -1)
	return int32(len(camelCaseCount) + 1)
}

func gopherCountCamel(s string) int {
	answer := 1
	for _, ch := range s {
		str := string(ch)
		if strings.ToUpper(str) == str {
			answer++
		}
	}
	return answer
}
