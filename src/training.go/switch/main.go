package main
import "fmt"

func main() {
weekday := 7
fmt.Printf("The day of the week is %d\n", weekday)

switch weekday {
case 1:
    fmt.Println("Beginning of the week, let's go to work")
case 3:
	fmt.Println("Midweek, let's keep going")
case 6,7:
	fmt.Println("Weekend, time to relax")
default:
	fmt.Println("It's just another day")
}

hour := 10
fmt.Printf("The hour is %d\n", hour)
switch {
case hour < 12:
	fmt.Println("It's morning")
case hour > 12 && hour < 18:
	fmt.Println("It's afternoon")
default:
	fmt.Println("It's evening")
}
}