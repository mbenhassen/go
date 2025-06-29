package main

import "fmt"

func main() {
	fmt.Println("=======================================")
	fmt.Println("Using a for loop to sum numbers from 1 to 10")
	sum := 0
	for i := 1; i <= 10; i++ {
		fmt.Printf("i=%d\n", i)
		sum += i
}
	fmt.Printf("The sum of numbers from 1 to 9 is %d\n", sum)
	fmt.Println("=======================================")
	eventCnt := 0
	for eventCnt < 3 {
		fmt.Printf("Retriving event ...")
		eventCnt++
		if eventCnt == 3 {
			fmt.Println("Got %d events\n", eventCnt)
		}
	}
	fmt.Println("=======================================")
	fmt.Println("Boucle infinie")
	i:= 0
	for {
		i++
		fmt.Printf("i=%d\n", i)
		if i%2 != 0 {
			fmt.Println("Odd Looping")
			continue
		}

		if i >= 10 {
			fmt.Println("Even Looping, breaking the loop")
			break
		}
	}
}