package primes

import "testing"

func BenchmarkPrimes(b *testing.B) {
	for _, bm := range Cases {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = CountPrimes(IsPrime, bm.inputN)
			}
		})
	}
}

func BenchmarkPrimesSqrt(b *testing.B) {
	for _, bm := range Cases {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = CountPrimes(IsPrimeSqrt, bm.inputN)
			}
		})
	}
}
