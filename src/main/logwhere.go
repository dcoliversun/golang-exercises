package main

import (
	"fmt"
	"log"
)

func main() {
	var where = log.Print
	a := 1
	where()
	b := a + 1
	where()
	fmt.Printf("b = %d\n", b)
	where()
	i := add(a, b)
	fmt.Printf("res = %d\n", i)
}

func add(a int, b int) int {
	return a + b
}
