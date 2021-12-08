package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{2, 3, 1, 5, 2}
	sort.Ints(a)
	fmt.Println(sort.IntsAreSorted(a))
	fmt.Println(a)
}
