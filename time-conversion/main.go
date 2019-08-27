package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func timeConversion(s string) string {
	isAM := strings.HasSuffix(s, "AM")
	re := regexp.MustCompile("(?i)(am|pm)")
	s = re.ReplaceAllString(s, "")
	timeSplit := strings.FieldsFunc(s, func(r rune) bool {
		return string(r) == ":"
	})
	hour, _ := strconv.ParseInt(timeSplit[0], 10, 0)

	if !isAM && hour != 12 {
		hour += 12
	} else if hour == 12 && isAM {
		hour = 0
	}
	timeSplit[0] = fmt.Sprintf("%02d", hour)
	return strings.Join(timeSplit, ":")
}

func main() {
	input := "12:05:45PM"
	converted := timeConversion(input)
	fmt.Println(converted)
}
