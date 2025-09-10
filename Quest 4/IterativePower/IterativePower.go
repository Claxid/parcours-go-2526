package main

import "fmt"

func main() {
	fmt.Println(IterativePower(5, 5))

}

func IterativePower(Nb int, power int) int {
	if power < 0 {
		return 0
	}
	if Nb == 0 {
		return 1
	}
	result := 1
	for i := 0; i < power; i++ {
		result *= Nb
	}

	return result

}
