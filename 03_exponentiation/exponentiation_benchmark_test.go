package exponentiation

import "testing"

func BenchmarkCalculate(b *testing.B) {
	for _, bm := range Cases {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = Calculate(bm.inputA, bm.inputN)
			}
		})
	}
}

func BenchmarkCalculateMath(b *testing.B) {
	for _, bm := range Cases {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = CalculateMath(bm.inputA, bm.inputN)
			}
		})
	}
}
