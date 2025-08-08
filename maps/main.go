package main

import (
	"fmt"
)

func main() {
	employees := make(map[string]int, 5)

	fmt.Println(employees)

	if employees == nil {
		fmt.Println("Employee map is nil")
	} else {
		fmt.Println("Employee map is not nil")
	}
	employees["John Doe"] = 23
	employees["Mary Doe"] = 19

	val, exist := employees["John Do"]
	if exist {
		fmt.Printf("Employee exits [%d]: %v", val, exist)
	}

}
