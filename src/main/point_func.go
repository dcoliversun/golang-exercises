package main

import "fmt"

func main() {
	var a = 1
	var b = &a
	method1(b)
	method2(a)
}

func method1(in *int) {
	c := in
	fmt.Println(c)
	fmt.Println(*c)
	fmt.Println(&c)
}

func method2(in int) {
	d := &in
	fmt.Println(d)
	fmt.Println(*d)
	fmt.Println(&d)
}