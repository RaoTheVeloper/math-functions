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

func LCMPrimeFactorization(a, b int) int {
	if a < 0 {
		a = -a
	}

	if b < 0 {
		b = -b
	}

	if a == 0 || b == 0 {
		return 0
	}

	aFactors := PrimeFactorization(a)
	bFactors := PrimeFactorization(b)

	factors := merge(aFactors, bFactors)

	result := 1
	for _, factor := range factors {
		result *= factor
	}

	return result
}

type number interface {
	int | float64 | float32 | uint | uint8 | uint16 | uint32 | uint64 | int8 | int16 | int32 | int64
}

func merge[T number](a, b []T) []T {
	result := []T{}
	if len(a) == 0 {
		return b
	}
	if len(b) == 0 {
		return a
	}
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] == b[j] {
			result = append(result, a[i])
			i++
			j++
		} else if a[i] < b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}
	for i < len(a) {
		result = append(result, a[i])
		i++
	}
	for j < len(b) {
		result = append(result, b[j])
		j++
	}
	return result
}
