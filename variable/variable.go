package main

import "fmt"


func main() {
	fmt.Println("<------Variable---->")

	var name string = "Wiliam Sagala"
	fmt.Println(name)

	fmt.Println("<----dapat ubah nilai nya---->")
	name = "Wiliam Fery Sagala"
	fmt.Println(name)

	fmt.Println("<----Penggunaan keyword := ----->")
	fullName := name
	fmt.Println("fullName := ", fullName)

	
}

