package main

import "fmt"

// var name = "Ferdous"

func main() {
	// MAIN TYPES
	//string
	//bool
	//int
	//int int8 int16 int32
	//uint uint8 uint16 uint32
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// Using var

	var age int32 = 37
	var isCool = true
	isCool = false

	name := "Ferdous"
	size := 1.3

	name, email := "Ferdous", "ferdousul.haque@gmail.com"

	fmt.Println(name, email, age, isCool)
	fmt.Printf("%T\n", size)
}
