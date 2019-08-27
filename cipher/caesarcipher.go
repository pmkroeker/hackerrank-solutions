package main

import (
	"bytes"
	"fmt"
)

func main() {
	var length, delta int
	var input string
	fmt.Scanf("%d\n", &length)
	fmt.Scanf("%s\n", &input)
	fmt.Scanf("%d\n", &delta)

	var ret bytes.Buffer
	for _, ch := range input {
		ret.WriteRune(cipher(ch, delta))
	}
	fmt.Println(ret.String())
}

func rotate(r rune, base, delta int) rune {
	tmp := int(r) - base
	tmp = (tmp + delta) % 26
	return rune(tmp + base)
}

func cipher(r rune, delta int) rune {
	if r >= 'A' && r <= 'Z' {
		return rotate(r, 'A', delta)
	} else if r >= 'a' && r <= 'z' {
		return rotate(r, 'a', delta)
	}
	return r
}
