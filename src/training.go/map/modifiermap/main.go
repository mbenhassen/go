package main

import "fmt"

func main() {

	m := make(map[string]int)
	fmt.Printf("Map content %v (len=%v)\n", m, len(m))

	m["hello"] = 5
	m["goodbye"] = 10
	fmt.Printf("Map content %v (len=%v)\n", m, len(m))

	fmt.Printf("key=hello, value=%v \n", m["hello"])
	i := m["goodbye"]
	fmt.Println(i)
    
	j, ok := m["helo"]
	fmt.Printf("value j=%v et ok=%v\n",j,ok)
	
	m["beatles"] = 2
	if _,ok := m["beatles"] ; ok {
		fmt.Println("Beatles key exist ! ")
	}

	fmt.Printf("Map content avant delete %v (len=%v)\n", m, len(m))
	delete(m,"goodbye")
	fmt.Printf("Map content apres delete %v (len=%v)\n", m, len(m))

	m2 := m
	fmt.Printf("Content m2 %v et len(m2)=%v\n",m2,len(m2))
	m2["help"] = 44
	fmt.Printf("Map content m %v (len=%v)\n", m, len(m))
	fmt.Printf("Content m2 %v et len(m2)=%v\n",m2,len(m2))



}