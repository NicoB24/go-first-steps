package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello World!")

	var x int = 1 // int32 or int64 depending on the machine
	var x32 int32 = 1
	var x64 int64 = 1
	fmt.Println(x)
	fmt.Println(x32)
	fmt.Println(x64)

	fmt.Printf("%T,%T,%T \n", x, x32, x64) // print types

	var unsigned int = 1 // number without sign
	fmt.Println(unsigned)

	y := 2
	fmt.Println(x + y)  // 3
	fmt.Println(x % y)  // 1
	fmt.Println(x == y) // false
	fmt.Println(x < y)  // true
	fmt.Println()

	// floating point numbers
	fmt.Println("Floating point numbers")
	pi := 3.141
	fmt.Println(pi)
	fmt.Printf("%T", pi)
	fmt.Println()

	var pi32 float32 = 3.141
	fmt.Printf("%T", pi32)
	fmt.Println()

	pi = float64(3.141)
	fmt.Println(pi)

	ninjas := 1e100
	fmt.Println(ninjas)

	// conversion between int and float
	a := 1
	b := 3.1416
	fmt.Println(a)
	fmt.Println(b)

	a = int(b) // converted b from float to int
	fmt.Println(a)

	// math library
	fmt.Println("Math library")
	fmt.Println(math.Ceil(3.0001))     // 4
	fmt.Println(math.Floor(3.0000001)) // 3
	fmt.Println(math.Min(1, 0))        // 0
	fmt.Println(math.Max(1, 0))        // 1
	fmt.Println(math.Abs(-1))          // 1
	fmt.Println(math.Sqrt(100))        // 10
	fmt.Println(math.Pow(2, 3))        // 8

	// complex numbers
	fmt.Println("Complex numbers")
	fmt.Println(complex(1, 2))
}
