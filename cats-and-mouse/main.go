package main

import (
	"fmt"
	"math"
)

const (
	catA  = "Cat A"
	catB  = "Cat B"
	mouse = "Mouse C"
)

// Complete the catAndMouse function below.
func catAndMouse(x int32, y int32, mouseC int32) string {
	distA := math.Abs(float64(x) - float64(mouseC))
	distB := math.Abs(float64(y) - float64(mouseC))
	switch {
	case distA < distB:
		return catA
	case distB < distA:
		return catB
	default:
		return mouse
	}
}

func main() {
	x, y, z := int32(1), int32(2), int32(3) // => Cat B
	a, b, c := int32(1), int32(3), int32(2) // Return Mouse C
	one := catAndMouse(x, y, z)
	two := catAndMouse(a, b, c)
	fmt.Println(one, two)
}
