package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noktpArsi NoKTP = "350404040404040"
	var marriedStatus Married = true
	fmt.Println(noktpArsi)
	fmt.Println(marriedStatus)

}
