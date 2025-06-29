package main

import (
	"fmt"
	"strings"
)

func TolowerStr(name string) (string, int) {
     return strings.ToLower(name), len(name)
}

func main() {
    LowerName, len :=TolowerStr("ALICE")
	fmt.Printf("The name in lowercase is %s and its length is %d\n", LowerName, len)
	_, len = TolowerStr("BOB")
	fmt.Printf("The length of the name BOB is %d\n", len)
	LowerName, _ = TolowerStr("CHARLIE")
	fmt.Printf("The name in lowercase is %s\n", LowerName)
}