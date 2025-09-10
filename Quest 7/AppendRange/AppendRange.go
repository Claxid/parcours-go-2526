package main

import "fmt"

func main() {
	fmt.Println(AppendRange(5, 10))
	fmt.Println(AppendRange(10, 5))
}

func AppendRange(min int, max int) []int {
	if min >= max {
		return []int{}
	}
	var result []int
	for i := min; i < max; i++ {
		result = append(result, i)
	}
	return result
}
