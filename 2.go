package main

import (
	"fmt"
)

func fibonacci(limit int) []int {
	numbrs := []int{0, 1}
	for i := 0; i < (limit / 2); i++ {
		if numbrs[i+1] <= limit {
			numbrs = append(numbrs, numbrs[i]+numbrs[i+1])
		} else {
			return numbrs
		}
	}
	return numbrs
}
func main() {
	sum := 0
	numbrs := fibonacci(4000000)
	for _, element := range numbrs {
		if element%2 == 0 {
			sum += element
		}
	}
	fmt.Println(sum)
}
