package student

func ForEach(f func(int), slice []int) {
	for _, v := range slice {
		f(v)
	}
}
