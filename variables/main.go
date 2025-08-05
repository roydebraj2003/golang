package main

import "fmt"

func main() {

	//unsigned int - uint8 -> uint64
	var unum uint64 = 4
	//signed int - int8 -> int64
	var snum int64 = 64
	//string
	var str = "this is a string"
	//byte
	var bite byte = 5
	//rune
	var run rune = 10

	var fnum float64 = 3.14

	msg := fmt.Sprintf("This is a formatter string with integer %d and float %f\n", unum, fnum)
	fmt.Printf(msg)

	fmt.Printf("This is a number type %T, value %d\n", unum, unum)
	fmt.Printf("This is a number type %T, value %d\n", snum, snum)
	fmt.Printf("This is a variable type %T, value %s\n", str, str)
	fmt.Printf("This is a variable type %T, value %d\n", bite, bite)
	fmt.Printf("This is a variable type %T, value %d\n", run, run)

}
