package main

import (
	"fmt"
	"strings"

 // Importing the input package
	"training.go/helloworld/input"

)

func main() {
	input.Keyboard() // Call the Keyboard function from the input package
	input.Mouse()    // Call the Mouse function from the input package
	// Print a greeting message to the console
	fmt.Println(strings.ToUpper("Hello, World!"))
	
	// Print a message indicating the end of the program
	fmt.Println("End of program")
}
