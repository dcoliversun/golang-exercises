package main

import (
	"log"
)

func main() {
	test("abc")
}

func test(s string) (res string, err error) {
	defer func() {
		log.Printf("test(%s) = %s, %v", s, res, err)
	}()
	return "result", nil
}