package main

import "fmt"

func UpdateVal(val string){
	val = "Value"
}

func UpdatePtr(ptr *string){
	*ptr = "pointer"
}

func main() {
	i:= 1
	var p *int = &i // Pointer to i
	fmt.Printf("Value of i: %d\n", i)
	fmt.Printf("Valeur de p: %v\n", p)
	fmt.Printf("Value pointed by p: %d\n", *p) // Dereferencing the pointer
	fmt.Println("===========================")

	s := "Hello, World!"
	sPtr := &s
	fmt.Printf("Value of s: %s\n", s)
	fmt.Printf("Value of sPtr: %v\n", sPtr)
	s2 := *sPtr // Dereferencing the pointer to get the value
	fmt.Printf("Value of s2: %s\n", s2)
	fmt.Printf("Value pointed by sPtr: %s\n", *sPtr) // Dereferencing the pointer	
	fmt.Println("===========================")
    
	*sPtr = "Alice"
	fmt.Printf("Updated value of s: %s\n", s)
	fmt.Printf("Updated value of s2: %s\n", s2)
	fmt.Printf("Value of sPtr: %v\n", &sPtr)
	fmt.Printf("Updated value pointed by sPtr: %s\n", *sPtr)
	fmt.Println("===========================")

	UpdateVal(s)
	fmt.Println("func UpdateVal(s) called")
	fmt.Printf("Updated value of s: %s\n", s)
	fmt.Printf("Updated value pointed by sPtr: %s\n", *sPtr)
	fmt.Println("===========================")

	UpdatePtr(&s)
	fmt.Println("func UpdatePtr(s) called")
	fmt.Printf("Updated value of s: %s\n", s)
	fmt.Printf("Updated value pointed by sPtr: %s\n", *sPtr)
    fmt.Println("===========================")

}