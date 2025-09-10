package main

import "fmt"

func main() {
	fmt.Printf("%#v\n", SplitWhiteSpaces("Hello how are you?"))
}

func SplitWhiteSpaces(s string) []string {
	liste := []string{}
	mot := ""

	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			mot += string(s[i])
		} else if mot != "" {
			liste = append(liste, mot)
			mot = ""
		}
	}

	if mot != "" {
		liste = append(liste, mot)
	}

	return liste
}
