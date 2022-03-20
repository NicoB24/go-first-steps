// scope of a variables is where we can access and use it
// there are 3 levels of scope: block, package, universe. block -> package -> universe
// block: between curly braces. We can have variables with the same names in differents blocks
// package: in the same folder
// universe: in different folder

package main // Its recommended to create a folder with the same name as the package

import (
	"fmt"
	"nico/go-programs/integers"
)

var five = 5 // scope: package, we can access to this variable in differents functions and files

func main() {

	{
		var diff = 2
		fmt.Println(diff)
	}

	var integer = 1 // scope: block
	fmt.Println("Hello world!")
	fmt.Println(integer)
	// fmt.Println(diff) if i try to access this variable it will fail because it was defined in another scope
	fmt.Println("Five in Main")
	fmt.Println(five)
	fmt.Println("Three from other file (integers.go)")
	fmt.Println(three) // I have to execute the two files using go run *.go
	fmt.Println("Three from other file (integers/integers_new.go)")
	fmt.Println(integers.Four)
}

func nonMain() {
	fmt.Println("Five in nonMain")
	fmt.Println(five)
}
