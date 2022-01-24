package main

import "fmt"

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func main() {
	sayHelloTo("Budi", "Hamon")
	pertama := "Arsi"
	kedua := "Amallah"
	sayHelloTo(pertama, kedua)

}
