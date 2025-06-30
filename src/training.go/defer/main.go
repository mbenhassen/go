package main

import "fmt"

func start() {
	fmt.Println("Starting the program...")
}

func finish() {
	fmt.Println("Finishing the program...")
}

func main() {
	start()
	defer finish()
	name := []string{"Bob", "Alice", "Bobette", "John"}
	for _, n := range name { 
		fmt.Printf("Bonjour %v\n", n)
		defer fmt.Printf("Au revoir %v\n", n)
	}
}
