package student

func Countif(f func(string) bool, a []string) int {
	count := 0
	for _, value := range a {
		if f(value) {
			count++
		}
	}
	return count

}
