package sort

import (
	"math"
)

type Sort struct{}

func (s *Sort) swap(arr []int, a int, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}

func (s *Sort) moveMaxToRoot(arr []int, root int, size int) {
	for j := root + 1; j < size; j++ {
		if arr[root] < arr[j] {
			s.swap(arr, root, j)
		}
	}
}

func (s *Sort) moveMaxToRootHeap(arr []int, root int, size int) {
	L := 2*root + 1
	R := 2*root + 2
	X := root

	if L < size && arr[L] > arr[X] {
		X = L
	}
	if R < size && arr[R] > arr[X] {
		X = R
	}
	if X == root {
		return
	}

	s.swap(arr, X, root)
	s.moveMaxToRootHeap(arr, X, size)
}

func (s *Sort) Bubble(input []int) (arr []int) {
	arr = append(arr, input...)

	N := len(arr)
	for i := 0; i < N; i++ {
		for j := 0; j < N-i-1; j++ {
			if arr[j] > arr[j+1] {
				s.swap(arr, j, j+1)
			}
		}

	}

	return arr
}

func (s *Sort) Insertion(input []int) (arr []int) {
	arr = append(arr, input...)

	N := len(input)

	for i := 1; i < N; i++ {
		j := i

		for j > 0 && arr[j-1] > arr[j] {
			s.swap(arr, j-1, j)

			j--
		}
	}

	return arr
}

func (s *Sort) Shell(input []int) (arr []int) {
	arr = append(arr, input...)

	N := len(arr)
	gap := int(math.Floor(float64(N) / 2))

	for gap > 0 {
		for i := gap; i < N; i++ {
			temp := arr[i]
			j := i

			for j >= gap && arr[j-gap] > temp {
				arr[j] = arr[j-gap]
				j -= gap
			}

			arr[j] = temp
		}

		gap = gap / 2
	}

	return arr
}

func (s *Sort) Selection(input []int) (arr []int) {
	arr = append(arr, input...)

	N := len(arr)

	s.moveMaxToRoot(arr, 0, N)

	for k := N - 1; k > 0; k-- {
		s.swap(arr, 0, k)
		s.moveMaxToRoot(arr, 0, k)
	}

	return arr
}

func (s *Sort) Heap(input []int) (arr []int) {
	arr = append(arr, input...)

	N := len(arr)

	for root := N/2 - 1; root >= 0; root-- {
		s.moveMaxToRootHeap(arr, root, N)
	}

	for k := N - 1; k > 0; k-- {
		s.swap(arr, 0, k)
		s.moveMaxToRootHeap(arr, 0, k)
	}

	return arr
}

func (s *Sort) BucketSort(input []int, chunk int) (result []int) {
	array := append([]int{}, input...)
	result = make([]int, 0)

	var max, min int
	for _, n := range array {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}

	nBuckets := int(max-min)/chunk + 1
	buckets := make([][]int, nBuckets)
	for i := 0; i < nBuckets; i++ {
		buckets[i] = make([]int, 0)
	}

	for _, n := range array {
		idx := int(n-min) / chunk
		buckets[idx] = append(buckets[idx], n)
	}

	for _, bucket := range buckets {
		if len(bucket) > 0 {
			bucket = s.Insertion(bucket)
			result = append(result, bucket...)
		}
	}

	return result
}

func (s *Sort) CountingSort(input []int, max int) (arr []int) {
	arr = append(arr, input...)

	n := len(input)

	output := make([]int, n)
	count := make([]int, max)

	for i := 0; i < max; i++ {
		count[i] = 0
	}

	for i := 0; i < n; i++ {
		count[arr[i]]++
	}

	for i := 1; i < max; i++ {
		count[i] += count[i-1]
	}

	for i := n - 1; i >= 0; i-- {
		output[count[arr[i]]-1] = arr[i]
		count[arr[i]]--
	}

	for i := 0; i < n; i++ {
		arr[i] = output[i]
	}

	return arr
}
