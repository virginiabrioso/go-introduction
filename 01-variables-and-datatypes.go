package main

import "fmt"

// this is a single line comment
/*
	This
	is a
	multiline
	comment
*/

// Variable at the package level
// string
var s string = "Hello World"
var s1, s2, s3 string = "Hello", "World", "All"

// signed integer +/-
var a int = -2
var a1 int8 = 10 // signed 8-bit integers (-128 to 127)
var a3 int16     // signed 16-bit integers
var a4 int32     // signed 32-bit integers
var a5 int64     // signed 64-bit integers

// unsigned integer +; uint == UP or above or equal to 0
var b uint = 2
var b1 uint8 = 255                   // unsigned 8-bit integers (0 to 255) | 2^8=256
var b2 uint16 = 65535                // unsigned 8-bit integers (0 to 65535) | 2^16=65536
var b3 uint32 = 4294967295           // unsigned 8-bit integers (0 to 4294967295) | 2^32=4294967296
var b4 uint64 = 18446744073709551615 // unsigned 8-bit integers (0 to 18446744073709551615) | 2^64=18446744073709551616

// float, use for decimals usually
var c float32
var c1 float64

// complex
var d complex64
var d1 complex128

// constant
const e string = "This will never change"

func main() {

	// Printing package level variables
	fmt.Println(s)
	fmt.Println(s1, s2, s3)
	fmt.Println(a, b)
	fmt.Println(a1, b1)

	// Variable at function level
	var s4 string = "Everyone"
	fmt.Println(s1, s2, s3, s4)

	// short assignment operator, work only at function level
	s5 := ":)"
	fmt.Println(s1, s2, s3, s4, s5)

	fmt.Println(e)

}
