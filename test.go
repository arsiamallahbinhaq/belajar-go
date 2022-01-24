package main

import (
	"fmt"
	"strconv"
)

var binaryNumbers = 0
var maxGap = 0
var currentGap int = 0

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func solution(N int) {
	var binaryNumbers = strconv.FormatInt(int64(N), 2)
	fmt.Println(binaryNumbers)
	for _, digit := range binaryNumbers {
		if digit == '0' {
			currentGap++
		} else {
			maxGap = Max(maxGap, currentGap)
			currentGap = 0
		}
		return maxGap
	}
}

func main() {
	solution(10)
}
