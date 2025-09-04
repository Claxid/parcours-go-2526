package main

import (
	"fmt"
)

func main() {
	IsNegative(-42)
	IsNegative(42)
	IsNegative(0)
}

func IsNegative(nb int) {
	if nb < 0 {
		fmt.Println("F")
	} else {
		fmt.Println("T")
	}
}
