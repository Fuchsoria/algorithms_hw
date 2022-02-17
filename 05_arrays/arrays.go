package arrays

type array[T any] struct {
	size int
	arr  []T
}

func (a *array[T]) isEmpty() bool {
	return a.size == 0
}

func (a *array[T]) Size() int {
	return a.size
}

func (a *array[T]) Len() int {
	return len(a.arr)
}

func (a *array[T]) Get(index int) T {
	return a.arr[index]
}

func (a *array[T]) GetArray() []T {
	return a.arr[:a.size]
}

func (a *array[T]) AddByIndex(item T, index int) {
	copyArr := make([]T, len(a.arr))
	copy(copyArr, a.arr)

	newArr := make([]T, 0)
	newArr = append(newArr, copyArr[:index]...)
	newArr = append(newArr, item)
	newArr = append(newArr, copyArr[index:]...)

	a.arr = newArr
	a.size++
}

func (a *array[T]) Remove(index int) T {
	value := a.arr[index]

	a.arr = append(a.arr[:index], a.arr[index+1:]...)

	a.size--

	return value
}

type singleArray[T any] struct {
	array[T]
}

func (s *singleArray[T]) Resize() {
	newArr := make([]T, s.Size()+1)

	copy(newArr, s.arr)

	s.arr = newArr
	s.size++
}

func (s *singleArray[T]) Add(item T) {
	s.Resize()
	s.arr[s.Size()-1] = item
}

func NewSingleArr[T any]() *singleArray[T] {
	return &singleArray[T]{array[T]{arr: make([]T, 0), size: 0}}
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

func (s *vectorArray[T]) GetArray() []T {
	return s.arr[:s.size]
}

func (s *vectorArray[T]) Add(item T) {
	if s.size == s.Len() {
		s.resize()
	}

	s.arr[s.size] = item
	s.size++
}

func (s *vectorArray[T]) AddByIndex(item T, index int) {
	copyArr := make([]T, len(s.arr))
	copy(copyArr, s.arr)

	newArr := make([]T, 0)
	newArr = append(newArr, copyArr[:index]...)
	newArr = append(newArr, item)
	newArr = append(newArr, copyArr[index:]...)

	s.arr = newArr
	s.size = len(s.arr)
}

func (s *vectorArray[T]) Remove(index int) T {
	value := s.arr[index]

	s.arr = append(s.arr[:index], s.arr[index+1:]...)

	s.size--

	return value
}

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

func (s *factorArray[T]) isEmpty() bool {
	return s.size == 0
}

func (s *factorArray[T]) resize() {
	if s.isEmpty() {
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

func (s *factorArray[T]) GetArray() []T {
	return s.arr[:s.size]
}

func (s *factorArray[T]) Add(item T) {
	if s.size == s.Len() {
		s.resize()
	}

	s.arr[s.size] = item
	s.size++
}

func (s *factorArray[T]) AddByIndex(item T, index int) {
	copyArr := make([]T, len(s.arr))
	copy(copyArr, s.arr)

	newArr := make([]T, 0)
	newArr = append(newArr, copyArr[:index]...)
	newArr = append(newArr, item)
	newArr = append(newArr, copyArr[index:]...)

	s.arr = newArr
	s.size++
}

func (s *factorArray[T]) Remove(index int) T {
	value := s.arr[index]

	s.arr = append(s.arr[:index], s.arr[index+1:]...)

	s.size--

	return value
}

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
