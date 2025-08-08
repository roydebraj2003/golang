package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// Example: Capture output
	cmd := exec.Command("ls", "-l")
	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Output of ls -l:\n%s\n", string(out))

	// Example: Run a process without waiting for output immediately
	cmd1 := exec.Command("sleep", "3")

	fmt.Println("Starting sleep...")
	err1 := cmd1.Start()
	if err1 != nil {
		panic(err1)
	}

	// Wait for the process to finish
	err2 := cmd1.Wait()
	if err2 != nil {
		panic(err2)
	}
	fmt.Println("Completed sleep.")
}
