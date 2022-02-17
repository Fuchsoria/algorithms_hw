package arrays

type singleArray[T any] struct {
	arr []T
}

func (s *singleArray[T]) Size() int {
	return len(s.arr)
}

func (s *singleArray[T]) Resize() {
	newArr := make([]T, s.Size()+1)

	copy(newArr, s.arr)

	s.arr = newArr
}

func (s *singleArray[T]) Get(index int) T {
	return s.arr[index]
}

func (s *singleArray[T]) Add(item T) {
	s.Resize()
	s.arr[s.Size()-1] = item
}

func (s *singleArray[T]) AddByIndex(item T, index int) {
	firstPart := s.arr[:index]
	secondPart := s.arr[index+1:]

	newArr := append(firstPart, item)
	s.arr = append(newArr, secondPart...)
}

func (s *singleArray[T]) Remove(index int) T {
	value := s.arr[index]

	s.arr = append(s.arr[:index], s.arr[index+1:]...)

	return value
}

func NewSingleArr[T any]() *singleArray[T] {
	return &singleArray[T]{make([]T, 0)}
}

type vectorArray[T any] struct {
	vector int
	size   int
	arr    []T
}

func (s *vectorArray[T]) Len() int {
	return len(s.arr)
}

func (s *vectorArray[T]) resize() {
	newArr := make([]T, s.size+s.vector)

	copy(newArr, s.arr)

	s.arr = newArr
}

func (s *vectorArray[T]) Get(index int) T {
	return s.arr[index]
}

func (s *vectorArray[T]) Add(item T) {
	if s.size == s.Len() {
		s.resize()
	}

	s.arr[s.size] = item
	s.size++
}

func (s *vectorArray[T]) AddByIndex(item T, index int) {
	// need to implement
}

// func (s *vectorArray[T]) Remove(index int) T {
// 	// need to implement
// }

func NewVectorArr[T int](vector int) *vectorArray[T] {
	return &vectorArray[T]{vector, 0, make([]T, 0)}
}

type factorArray[T int] struct {
	factor int
	size   int
	arr    []T
}

func (s *factorArray[T]) Len() int {
	return len(s.arr)
}

func (s *factorArray[T]) resize() {
	if s.size == 0 {
		s.arr = make([]T, 1)

		return
	}

	newArr := make([]T, s.size*s.factor)

	copy(newArr, s.arr)

	s.arr = newArr
}

func (s *factorArray[T]) Get(index int) T {
	return s.arr[index]
}

func (s *factorArray[T]) Add(item T) {
	if s.size == s.Len() {
		s.resize()
	}

	s.arr[s.size] = item
	s.size++
}

func (s *factorArray[T]) AddByIndex(item T, index int) {
	// need to implement
}

// func (s *factorArray[T]) Remove(index int) T {
// 	// need to implement
// }

func NewFactorArr[T int]() *factorArray[T] {
	return &factorArray[T]{2, 0, make([]T, 0)}
}

type matrixArray[T int] struct {
	vector int
	size   int
	arr    *singleArray[*vectorArray[T]]
}

func (s *matrixArray[T]) Get(index int) T {
	return s.arr.Get(index / s.vector).Get(index % s.vector)
}

func (s *matrixArray[T]) Add(item T) {
	if s.size == s.arr.Size()*s.vector {
		s.arr.Add(NewVectorArr[T](s.vector))
	}

	s.arr.Get(s.size / s.vector).Add(item)
	s.size++
}

func NewMatrixArr[T int]() *matrixArray[T] {
	return &matrixArray[T]{100, 0, NewSingleArr[*vectorArray[T]]()}
}
