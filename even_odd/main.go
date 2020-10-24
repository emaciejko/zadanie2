package main

import (
	"fmt"
	"strconv"
)

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, n := range numbers {
		s := strconv.Itoa(numbers[i])
		fmt.Println(s + " is " + evenOdd(n))
	}
}

func evenOdd(n int) string {
	if n%2 == 0 {
		return "even number"
	} else {
		return "odd number"
	}
}
