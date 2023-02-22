package pkg

// calculate least common multiple of a and b
func LCM(a, b int) int {
	if a < 0 {
		a = -a
	}

	if b < 0 {
		b = -b
	}

	gcd := GCD(a, b)
	if gcd == 0 {
		return 0
	}

	return a * b / gcd
}
