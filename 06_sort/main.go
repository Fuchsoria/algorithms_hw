package main

import "fmt"

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
	copy(arr, input)

	return arr
}

func (s *Sort) Shell(input []int) (arr []int) {
	copy(arr, input)

	return arr
}

func main() {
	sort := Sort{}

	input := []int{8, 9, 6, 7, 3, 4, 2, 1, 5, 0}
	output := sort.Bubble(input)

	fmt.Println(output)
}
