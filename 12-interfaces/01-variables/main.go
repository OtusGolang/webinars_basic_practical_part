package main

import (
	"fmt"
)

func main() {

	var a int = 5
	var b float64 = 3.14
	var c string = "Hello, Go!"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// interfaces

	var i any

	i = a
	fmt.Printf("i is of type %T and value %v\n", i, i)

	i = b
	fmt.Printf("i is of type %T and value %v\n", i, i)

	i = c
	fmt.Printf("i is of type %T and value %v\n", i, i)

}
