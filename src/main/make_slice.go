package main

import "fmt"

func main() {
	var slice1 []int = make([]int, 10)
	for i := 0; i < len(slice1); i++ {
		slice1[i] = 5 * i
	}

	fmt.Printf("The len of slice is %d\n", len(slice1))
	fmt.Printf("The cap of slice is %d\n", cap(slice1))

	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
}
