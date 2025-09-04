package main

import "github.com/01/z01"

func main() {
	for x := 'z'; x <= 'a'; x-- {
		z01.PrintRune(x)
	}
	z01.PrintRune('\n')
}
