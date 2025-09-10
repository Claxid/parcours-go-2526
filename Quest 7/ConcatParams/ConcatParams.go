package main

import (
	"fmt"
)

func main() {
	test := []string{"Hello", "how", "are", "you?"}
	fmt.Println(ConcatParams(test))
}

func ConcatParams(test []string) string {
	result := ""
	for i := 0; i < len(test); i++ {
		result += test[i]
		if i != len(test)-1 {
			result += "\n"
		}

	}
	return result
}
