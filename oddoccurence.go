package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	var elmt int
	for i := 0; i < len(A); i++ {
		elmt ^= A[i]
	}
	return elmt
}
