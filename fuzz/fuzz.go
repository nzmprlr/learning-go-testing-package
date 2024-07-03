package fuzz

func Echo(x int) int {
	if x == 5 {
		return 5
	}

	if x == 10 {
		return 0
	}
	return x
}
