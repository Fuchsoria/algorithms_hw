package fibonacci

import "testing"

func BenchmarkCalculate(b *testing.B) {
	for _, bm := range Cases {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = Calculate(bm.inputN)
			}
		})
	}
}

func BenchmarkCalculateRecursive(b *testing.B) {
	Cases := Cases[:len(Cases)-1]

	for _, bm := range Cases {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = CalculateRecursive(bm.inputN)
			}
		})
	}
}

func BenchmarkCalculateGold(b *testing.B) {
	Cases := Cases[:len(Cases)-1]

	for _, bm := range Cases {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = CalculateGold(bm.inputN)
			}
		})
	}
}
