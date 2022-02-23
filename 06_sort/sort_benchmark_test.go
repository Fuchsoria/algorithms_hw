package sort

import "testing"

func BenchmarkBubble(b *testing.B) {
	sort := Sort{}

	b.Run("100", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.Bubble(case100)
		}
	})

	b.Run("1000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.Bubble(case1000)
		}
	})

	b.Run("10000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.Bubble(case10000)
		}
	})
}

func BenchmarkInsertion(b *testing.B) {
	sort := Sort{}

	b.Run("100", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.Insertion(case100)
		}
	})

	b.Run("1000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.Insertion(case1000)
		}
	})

	b.Run("10000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.Insertion(case10000)
		}
	})
}

func BenchmarkShell(b *testing.B) {
	sort := Sort{}

	b.Run("100", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.Shell(case100)
		}
	})

	b.Run("1000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.Shell(case1000)
		}
	})

	b.Run("10000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.Shell(case10000)
		}
	})
}
