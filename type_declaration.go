package main

import "fmt"

func main() {
	// ** TYPE **

	// farenheit & celsius
	type farenheit int
	type celsius int

	var f farenheit = 32
	var c celsius = 0

	// i cant assign c = f, insted c = celsius(f)
	c = celsius((f - 32) * 5 / 9)

	fmt.Println(f, c)
}
