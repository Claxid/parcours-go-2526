package main

import "fmt"

func main() {
	PrintComb()
}

func PrintComb() {

	for X := 0; X <= 7; X++ {
		for Y := X + 1; Y <= 8; Y++ {
			for Z := Y + 1; Z <= 9; Z++ {
				fmt.Println(X, Y, Z)
				fmt.Println('\n')
			}
		}
	}
}
