package main

import "fmt"

func printInfoNoparam() {
	fmt.Printf("Name =%s, ages =%d, email = %s\n","Bob", 30, "bob@bob.com")
}

func printInfoparam(name string, ages int, email string) {
	fmt.Printf("Name = %s, ages = %d, email = %s\n",name, ages, email)
}

func average(x, y float64) float64{
	return (x + y) / 2.0
}

func sumNamedReturn(x,y,z int) (sum int) {
	sum = x + y + z
	return sum
}

func main() {
	printInfoNoparam()
	printInfoparam("Alice", 25, "alice@alice.com")
	result := average(10.0, 20.0)
	fmt.Printf("The average is %f\n", result)
	result2 := sumNamedReturn(1, 2, 3)
	fmt.Printf("The sum is %d\n", result2)

}