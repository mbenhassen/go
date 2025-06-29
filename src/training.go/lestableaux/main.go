package main
import "fmt"




func main() {
  var names [3]string
  fmt.Printf("names=%v len=%v\n", names,len(names))
  fmt.Println("======================================")
  names[0] = "Bob"
  names[2] = "Alice"
  fmt.Printf("names=%v len=%v\n", names,len(names))
  fmt.Println("======================================")
  fmt.Printf("names[2]=%v\n", names[2])
  fmt.Println("======================================")
  odds := [5]int{1, 3, 5, 7, 9}
  fmt.Printf("odds=%v len=%v\n", odds, len(odds))
  fmt.Println("======================================")
  



}