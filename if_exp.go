package main

import "fmt"

func main() {
	name := "Ayah"

	if name == "Betu" {
		fmt.Println("Hello Betu")
	} else {
		fmt.Println("Hello", name)

	}

	if length := len(name); length > 5 {
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Cukup")
	}

}
