package main

import "fmt"

func main() {
	fmt.Println("Hello world!")

	// var name type = expression
	var integer int = 1

	// zero value (for int 0, for string empty string, for boolean False)
	var integer_2 int

	// multiple values
	var integer_3, integer_4 int

	// multiple values assign
	var integer_5, string = 1, "string"

	// short declaration name := expression
	boolean := true

	fmt.Println(integer)
	fmt.Println(integer_2)
	fmt.Println(integer_3, integer_4)
	fmt.Println(integer_5, string)
	fmt.Println(boolean)

	// *** POINTERS ***
	x := 1
	p := &x // address of variable x

	var x_2 int = 2
	var p_2 = &x_2 // address of variable x_2

	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(p_2)
	fmt.Println(*p_2)
}
