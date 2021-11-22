package main

import "fmt"

func main() {
	typecheck(1, "abc", 3.1415926, 'a');
}

func typecheck(values ... interface{}) {
	for _, value := range values {
		switch value.(type) {
		case int:
			fmt.Println("This is int")
		case string:
			fmt.Println("This is string")
		case float64:
			fmt.Println("This is float64")
		default:
			fmt.Println("Other")
		}
	}
}