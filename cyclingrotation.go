package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int, K int) []int {
	// write your code in Go 1.4
	length := len(A)
	if length > 0 {
		if K > length {
			K = K % length
		}
		A = append(A[length-K:length], A[0:length-K]...)
	}
	return A
}
