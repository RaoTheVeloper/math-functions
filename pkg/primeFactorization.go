package pkg

func PrimeFactorization(n int) []int {
	factors := []int{}
	if n < 0 {
		n = -n
	}

	if n == 0 || n == 1 {
		return factors
	}

	for n%2 == 0 {
		factors = append(factors, 2)
		n /= 2
	}

	for i := 3; i*i <= n; i += 2 {
		for n%i == 0 {
			factors = append(factors, i)
			n /= i
		}
	}

	if n > 2 {
		factors = append(factors, n)
	}

	return factors
}
