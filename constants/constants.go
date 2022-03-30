package main

import (
	"fmt"
	"math"
	"math/big"
	"time"
)

func main() {
	fmt.Println("Hello World!")

	const five = 5 // constants cant be reassigned
	// five = 10 this fails
	fmt.Println(five)

	const pi = 3.14
	fmt.Println(math.Pi)

	// we can declare multiple constants
	const (
		a = 1
		b // 1
		c = 3
		d // 3
	)

	fmt.Println(a, b, c, d)

	const (
		zero int = iota // auto increment
		one
		two
	)
	fmt.Println(zero, one, two)

	const (
		_ = 1 << (10 * iota) // binary shift
		kb
		mb
		gb
		tb
		pb
		eb
		zb
		yb
	)
	fmt.Println(kb, mb, gb)
	// fmt.Println(yb) // overflow
	fmt.Println(yb / zb)

	// common built-in constants
	fmt.Println(math.Pi)
	fmt.Println(time.February)
	fmt.Println(time.Second)
	fmt.Println(time.UTC)
	fmt.Println(big.MaxExp)
	fmt.Println(big.MinExp)
}
