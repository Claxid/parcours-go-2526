package main

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0 // les factoriels négatifs n'existent pas
	}

	result := 1
	for i := 1; i <= nb; i++ {
		// vérifier un overflow simple
		if result > (1<<63-1)/i { // limite pour int64
			return 0
		}
		result *= i
	}
	return result
}
