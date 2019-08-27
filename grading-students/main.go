package main

import "fmt"

// This function can be pasted into the hackerrank reponse
func gradingStudents(grades []int32) []int32 {
	for idx, grade := range grades {
		if grade < 38 {
			continue
		}
		remain := grade % 5
		if remain >= 3 {
			grades[idx] = grade + (5 - remain)
		}
	}
	return grades
}
func main() {
	grades := []int32{73, 67, 38, 33} // Sample Input 0
	// Expected Output {75, 67, 40, 33}
	newGrades := gradingStudents(grades)
	fmt.Println(newGrades)
}
