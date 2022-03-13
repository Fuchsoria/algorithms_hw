package sort

import (
	"fmt"
	"testing"
)

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

func BenchmarkSelection(b *testing.B) {
	sort := Sort{}

	b.Run("100", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.Selection(case100)
		}
	})

	b.Run("1000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.Selection(case1000)
		}
	})

	b.Run("10000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.Selection(case10000)
		}
	})
}

func BenchmarkHeap(b *testing.B) {
	sort := Sort{}

	b.Run("100", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.Heap(case100)
		}
	})

	b.Run("1000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.Heap(case1000)
		}
	})

	b.Run("10000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.Heap(case10000)
		}
	})
}

func BenchmarkMerge(b *testing.B) {
	sort := Sort{}

	b.Run("100", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.Merge(case100)
		}
	})

	b.Run("1000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.Merge(case1000)
		}
	})

	b.Run("10000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.Merge(case10000)
		}
	})

	b.Run("1kk", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			saveFile(sort.Merge(dataArr1kk), fmt.Sprintf("merge-1kk-%d", i))
		}
	})
}

func BenchmarkQuick(b *testing.B) {
	sort := Sort{}

	b.Run("100", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.Quick(case100)
		}
	})

	b.Run("1000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.Quick(case1000)
		}
	})

	b.Run("10000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.Quick(case10000)
		}
	})

	b.Run("1kk", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			saveFile(sort.Quick(dataArr1kk), fmt.Sprintf("quick-1kk-%d", i))
		}
	})
}

func BenchmarkBucket(b *testing.B) {
	sort := Sort{}

	b.Run("100", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.BucketSort(case100, 10)
		}
	})

	b.Run("1000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.BucketSort(case1000, 10)
		}
	})

	b.Run("10000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.BucketSort(case10000, 10)
		}
	})

	b.Run("1kk", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			saveFile(sort.BucketSort(dataArr1kk, 100), fmt.Sprintf("bucket-1kk-%d", i))
		}
	})
}

func BenchmarkCounting(b *testing.B) {
	sort := Sort{}

	b.Run("100", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.CountingSort(case100, 65535)
		}
	})

	b.Run("1000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.CountingSort(case1000, 65535)
		}
	})

	b.Run("10000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.CountingSort(case10000, 65535)
		}
	})

	b.Run("1kk", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			saveFile(sort.CountingSort(dataArr1kk, 65535), fmt.Sprintf("counting-1kk-%d", i))
		}
	})
}
