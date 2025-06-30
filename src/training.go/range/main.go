package main

import "fmt"

func main() {
  names := []string{"Bob", "Alice", "Bobette", "John"}
  for i,n := range names {
	fmt.Printf("Username=%s index=%d\n", n, i)
  }
  fmt.Println("======================================")
  for _,c:= range "golang"{
	fmt.Printf("Character=%c\n", c)
  }
}
