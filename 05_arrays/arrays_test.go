package arrays

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSingleArray(t *testing.T) {
	t.Run("add", func(t *testing.T) {
		inst := NewSingleArr[int]()

		inst.Put(10)
		inst.Put(11)
		inst.Put(12)

		require.Equal(t, []int{10, 11, 12}, inst.GetArray())
	})

	t.Run("add by index", func(t *testing.T) {
		inst := NewSingleArr[int]()

		inst.Put(10)
		inst.Put(11)
		inst.Put(12)
		inst.PutByIndex(14, 2)

		require.Equal(t, []int{10, 11, 14, 12}, inst.GetArray())
	})

	t.Run("remove", func(t *testing.T) {
		inst := NewSingleArr[int]()

		inst.Put(10)
		inst.Put(11)
		inst.Put(12)
		inst.PutByIndex(14, 2)

		result := inst.Remove(2)

		require.Equal(t, []int{10, 11, 12}, inst.GetArray())
		require.Equal(t, 14, result)
	})

	t.Run("get", func(t *testing.T) {
		inst := NewSingleArr[int]()

		inst.Put(10)
		inst.Put(11)
		inst.Put(12)
		inst.PutByIndex(14, 2)

		result := inst.Get(2)

		require.Equal(t, 14, result)
	})
}

func TestVectorArray(t *testing.T) {
	t.Run("add", func(t *testing.T) {
		inst := NewVectorArr(10)

		inst.Put(10)
		inst.Put(11)
		inst.Put(12)

		require.Equal(t, []int{10, 11, 12}, inst.GetArray())
	})

	t.Run("add by index", func(t *testing.T) {
		inst := NewVectorArr(3)

		inst.Put(10)
		inst.Put(11)
		inst.Put(12)
		inst.PutByIndex(14, 2)

		require.Equal(t, []int{10, 11, 14, 12}, inst.GetArray())
	})

	t.Run("remove", func(t *testing.T) {
		inst := NewVectorArr(3)

		inst.Put(10)
		inst.Put(11)
		inst.Put(12)
		inst.PutByIndex(14, 2)

		result := inst.Remove(2)

		require.Equal(t, []int{10, 11, 12}, inst.GetArray())
		require.Equal(t, 14, result)
	})

	t.Run("get", func(t *testing.T) {
		inst := NewVectorArr(3)

		inst.Put(10)
		inst.Put(11)
		inst.Put(12)
		inst.PutByIndex(14, 2)

		result := inst.Get(2)

		require.Equal(t, 14, result)
	})
}

func TestFactorArray(t *testing.T) {
	t.Run("add", func(t *testing.T) {
		inst := NewFactorArr()

		inst.Put(10)
		inst.Put(11)
		inst.Put(12)

		require.Equal(t, []int{10, 11, 12}, inst.GetArray())
	})

	t.Run("add by index", func(t *testing.T) {
		inst := NewFactorArr()

		inst.Put(10)
		inst.Put(11)
		inst.Put(12)
		inst.PutByIndex(14, 2)

		require.Equal(t, []int{10, 11, 14, 12}, inst.GetArray())
	})

	t.Run("remove", func(t *testing.T) {
		inst := NewFactorArr()

		inst.Put(10)
		inst.Put(11)
		inst.Put(12)
		inst.PutByIndex(14, 2)

		result := inst.Remove(2)

		require.Equal(t, []int{10, 11, 12}, inst.GetArray())
		require.Equal(t, 14, result)
	})

	t.Run("get", func(t *testing.T) {
		inst := NewFactorArr()

		inst.Put(10)
		inst.Put(11)
		inst.Put(12)
		inst.PutByIndex(14, 2)

		result := inst.Get(2)

		require.Equal(t, 14, result)
	})
}
