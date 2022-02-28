package sort

import (
	"math"
)

type Sort struct{}

func (s *Sort) swap(arr []int, a int, b int) {
	arr[a], arr[b] = arr[b], arr[a]
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

// need to implement
func (s *Sort) Selection(input []int) (arr []int) {
	return arr
}

// need to implement
func (s *Sort) Heap(input []int) (arr []int) {
	return arr
}
