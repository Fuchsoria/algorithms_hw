package primes

import (
	"math"
)

type PrimeFunc = func(n int64) bool

func CountPrimes(f PrimeFunc, n int64) int {
	count := 0

	for i := int64(2); i <= n; i++ {
		if f(i) {
			count++
		}
	}

	return count
}

func IsPrime(n int64) bool {
	for i := int64(2); i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func IsPrimeSqrt(n int64) bool {
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	for i := int64(3); float64(i) <= math.Sqrt(float64(n)); i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}
