package main

import "fmt"

func main() {

	day:= 21
	if day < 15 {
		fmt.Printf("Day %d is in the first half of the month.\n", day)
	} else if day == 18 {
		fmt.Printf("Day %d is the 18th of the month.\n", day)
	} else {
		fmt.Printf("Day %d is in the second half of the month.\n", day)
	}

	year, month, day := 2009, 12, 10

	fmt.Printf("The date is %d/%d/%d\n", day, month, year)
	if year == 2009 && month == 11 && day == 10 {
		fmt.Println("The date is the fist release of Go.")
	} else if year == 2009 || month == 11 || day == 10 {
		fmt.Println("At least one of the date components is correct.")
	} else {
		fmt.Println("Just another day")
	}
}