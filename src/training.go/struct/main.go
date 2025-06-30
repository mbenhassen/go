package main


type Adress struct {
   street,city,state string
}

type Person struct {
	Name string
	Age  int
	Addr Adress

}

func main() {
	var p Person
	p.Name = "Bob"
	p.Addr.city = "Lyon"
	
}
