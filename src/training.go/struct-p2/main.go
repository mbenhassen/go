package main

import (
	"fmt"
   "player"
)


func main() {
   var p player.Player
   p.Name = "Bob"
   p.Age = 10
   
   fmt.Printf("Player Name: %s, Age: %d\n", p.Name, p.Age)
   fmt.Printf("Player Details: %+v\n", p)
   fmt.Printf("Player Name : %v Age: %v\n", p.Name, p.Age)
   fmt.Println("==========================================================================")
   a:= player.Avatar{"http://avatar.jpg"}
   fmt.Printf("Avatar URL: %v\n", a)

   fmt.Println("==========================================================================")
   p3 := player.Player{ Name: "John", Age: 20, Avatar: player.Avatar{Url: "http://john.jpg"},}
   fmt.Printf("Player 3 %v", p3)

}