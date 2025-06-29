package main

import "fmt"

func main() {

	nums := make([]int, 2, 3) // Create a slice with a capacity of 3
	nums[0] = 10
	nums[1] = -1
	fmt.Printf("nums=%v len=%v cap=%v\n", nums, len(nums), cap(nums))
	fmt.Println("======================================")
	fmt.Println("Adding 20 to the slice")
	nums = append(nums, 20) // Append an element to the slice
	fmt.Printf("nums=%v len=%v cap=%v\n", nums, len(nums), cap(nums))
	fmt.Println("======================================")
	nums = append(nums, -8) // Append an element to the slice
	fmt.Printf("nums=%v len=%v cap=%v\n", nums, len(nums), cap(nums))
	fmt.Println("======================================")

	letters := []string{"g", "o", "l", "a", "n", "g"}
	fmt.Printf("letters=%v len=%v cap=%v\n", letters, len(letters), cap(letters))
	fmt.Println("======================================")
	letters = append(letters, "s") // Append an element to the slice
	fmt.Printf("letters=%v len=%v cap=%v\n", letters, len(letters), cap(letters))
	fmt.Println("======================================")
	fmt.Println("================slice de slice===============")
	sub1 := letters[0:2] // Slice from index 0 to 2
	fmt.Printf("sub1=%v len=%v cap=%v\n", sub1, len(sub1), cap(sub1))
	sub2 := letters[2:] // Slice from index 2 to the end
	fmt.Printf("sub2=%v len=%v cap=%v\n", sub2, len(sub2), cap(sub2))
	fmt.Println("======================================")
	sub2[0] = "UP" // Modify the first element of sub2
	fmt.Printf("sub1=%v len=%v cap=%v\n", sub1, len(sub1), cap(sub1))
	fmt.Printf("sub2=%v len=%v cap=%v\n", sub2, len(sub2), cap(sub2))
	fmt.Printf("letters=%v len=%v cap=%v\n", letters, len(letters), cap(letters))
	fmt.Println("======================================")
    subCopy := make([]string, len(sub1))
	copy(subCopy, sub1) // Copy sub1 to subCopy
	subCopy[0] = "DOWN" // Modify the first element of subCopy
	fmt.Printf("subCopy=%v len=%v cap=%v\n", subCopy, len(subCopy), cap(subCopy))
	fmt.Printf("sub1=%v len=%v cap=%v\n", sub1, len(sub1), cap(sub1))
	fmt.Printf("sub2=%v len=%v cap=%v\n", sub2, len(sub2), cap(sub2))
	fmt.Printf("letters=%v len=%v cap=%v\n", letters, len(letters), cap(letters))

}
	