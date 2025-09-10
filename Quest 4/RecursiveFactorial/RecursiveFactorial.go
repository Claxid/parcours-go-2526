package main

import "fmt"

func main() {
	RecursiveFactorial(8)

}

func RecursiveFactorial(Nb int) int {
	if Nb < 0 {
		return 0
	}
	if Nb > 20 {
		return 0
	}
	if Nb == 0 {
		return 1
	}
	fmt.Println(Nb)
	return Nb * RecursiveFactorial(Nb-1)

}
