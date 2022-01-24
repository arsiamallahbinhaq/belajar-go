package main

import "fmt"

func main() {
	name := "Arsi"

	switch name {
	case "Arsi":
		fmt.Println("Hello Arsi")
	case "Joko":
		fmt.Println("Hello Joko")
	default:
		fmt.Println("Kenalan Donk")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Terlalu panjang")
	case false:
		fmt.Println("Sesuai")
	}

}
