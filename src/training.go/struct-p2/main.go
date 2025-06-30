package main

import (
	"fmt"

   "training.go/player" // Importing the player package
)


func main() {
   var p player.Player
   p.Name = "Bob"
   p.Age = 30
   
   fmt.Printf("Player Name: %s, Age: %d\n", p.Name, p.Age)

}