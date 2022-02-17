package arrays

import "testing"

func BenchmarkSingleArray(b *testing.B) {
	Cases := Cases[:3]

	for _, bm := range Cases {
		b.Run(bm.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				arr := NewSingleArr[int]()

				for j := 0; j < bm.Input; j++ {
					arr.Put(j)
				}
			}
		})
	}
}

func BenchmarkVectorArray(b *testing.B) {
	Cases := Cases[:3]

	for _, bm := range Cases {
		b.Run(bm.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				arr := NewVectorArr(100)

				for j := 0; j < bm.Input; j++ {
					arr.Put(j)
				}
			}
		})
	}
}

func BenchmarkFactorArray(b *testing.B) {
	Cases := Cases[:3]

	for _, bm := range Cases {
		b.Run(bm.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				arr := NewFactorArr()

				for j := 0; j < bm.Input; j++ {
					arr.Put(j)
				}
			}
		})
	}
}

func BenchmarkMatrixArray(b *testing.B) {
	Cases := Cases[:3]

	for _, bm := range Cases {
		b.Run(bm.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				arr := NewMatrixArr()

				for j := 0; j < bm.Input; j++ {
					arr.Add(j)
				}
			}
		})
	}
}
