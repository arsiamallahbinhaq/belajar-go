package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(N int) int {
	// write your code in Go 1.4
	var result int
	var result_temp int
	var calc bool

	for N > 0 {
		if N%2 == 1 {
			if !calc {
				calc = true
			} else {
				if result_temp > result {
					result = result_temp
				}
				result_temp = 0
			}
		} else {
			if calc {
				result_temp++
			}
		}
		N = N / 2
	}

	return result
}
