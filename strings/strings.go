package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "Hello World"
	fmt.Println(s)
	fmt.Println(len(s))    // string length
	fmt.Println(s[0])      // ascii value
	fmt.Printf("%c", s[0]) // H
	fmt.Println()
	fmt.Println(s[:6]) // portion of string
	fmt.Println(s[6:]) // portion of string

	// concatenate string
	s = s + " Again"
	fmt.Println(s)

	// strings literal
	fmt.Println("Hello \t World ")
	fmt.Println("Hello \n World ")
	fmt.Println("Hello \b World ")

	// strings are slices of bytes
	abc := "abc"
	b := []byte(abc)
	fmt.Println(abc, b)
	abc2 := string(b)
	fmt.Println(abc2)

	// strings library
	fmt.Println(strings.ToUpper("Hello World"))
	fmt.Println(strings.ToLower("Hello World"))
	fmt.Println(strings.HasPrefix("Hello World", "Hello"))           // True
	fmt.Println(strings.HasSuffix("Hello World", "Hello"))           // False
	fmt.Println(strings.HasSuffix("Hello World", "World"))           // True
	fmt.Println(strings.Replace("Hello World", "World", "Hello", 1)) // Hello Hello

	// String Builder: Tool that we can use to build strings
	fmt.Println("String Builder")
	var sb strings.Builder
	fmt.Println("This is a String Builder: ", sb.String())
	fmt.Println("This is the capacity: ", sb.Cap()) // whole capacity, space allocated by string builder
	fmt.Println("This is the length: ", sb.Len())   // len is the length occupied

	sb.WriteString("Hello")
	fmt.Println("This is a String Builder: ", sb.String())
	fmt.Println("This is the capacity: ", sb.Cap()) // 0 2 4 8
	fmt.Println("This is the length: ", sb.Len())

	sb.Grow(100)                                    // we can grow the capacity
	fmt.Println("This is the capacity: ", sb.Cap()) // 0 2 4 8
	fmt.Println("This is the length: ", sb.Len())
	// with the grow we get 116 because it duplicates the length and added the grow term
	fmt.Println(sb.String())

	// We can reset the string builder
	sb.Reset()
	fmt.Println("This is the capacity: ", sb.Cap()) // 0 2 4 8
	fmt.Println("This is the length: ", sb.Len())
	fmt.Println(sb.String())

	// We can write the string by passing in slices of bytes
	sb.Write([]byte{65, 65, 65})
	fmt.Println(sb.String())

	// Conversions of strings and other data type
	fmt.Println("Conversion of strings and other data type")
	x := 123
	y := strconv.Itoa(x)
	fmt.Println(y)
	fmt.Printf("%T", y)
	fmt.Println()

	z, _ := strconv.Atoi(y) // The second term is the error, but i dont care here
	fmt.Println(z)
	fmt.Printf("%T", z)
	fmt.Println()
}
