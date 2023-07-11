package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println("Hello" + "World")
	fmt.Println(string("Hello World"[0]))
	var s string = "Hello World"

	s = strings.Replace(s, "H", "x", 1)
	fmt.Println(s)
	fmt.Println(strings.Contains(s, "World"))

	fmt.Println("Test\n" + "Test")
	fmt.Println(`Test
Test`)

	fmt.Println("\"")
	fmt.Println(`"`)
}
