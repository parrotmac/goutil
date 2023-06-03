package goutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnwrapOr(t *testing.T) {

	// Test nil pointer
	t.Run("Nil pointer", func(t *testing.T) {
		var i *int
		assert.Equal(t, UnwrapOr(i, 1), 1)
	})

	// Test non-nil pointer
	t.Run("Non-nil pointer", func(t *testing.T) {
		i := new(int)
		*i = 2
		assert.Equal(t, UnwrapOr(i, 1), 2)
	})
}

func TestOutside(t *testing.T) {

	// Test val < min
	t.Run("val < min", func(t *testing.T) {
		assert.True(t, Outside(1, 2, 3))
	})

	// Test val > max
	t.Run("val > max", func(t *testing.T) {
		assert.True(t, Outside(3, 1, 2))
	})

	// Test val == min
	t.Run("val == min", func(t *testing.T) {
		assert.False(t, Outside(1, 1, 2))
	})

	// Test val == max
	t.Run("val == max", func(t *testing.T) {
		assert.False(t, Outside(2, 1, 2))
	})

	// Floats

	// Test val < min
	t.Run("val < min", func(t *testing.T) {
		assert.True(t, Outside(1.0, 2.0, 3.0))
	})

	// Test val > max
	t.Run("val > max", func(t *testing.T) {
		assert.True(t, Outside(3.0, 1.0, 2.0))
	})

	// Test val == min
	t.Run("val == min", func(t *testing.T) {
		assert.False(t, Outside(1.0, 1.0, 2.0))
	})
}

func TestBounded(t *testing.T) {

	t.Run("val < min", func(t *testing.T) {
		assert.Equal(t, Bounded(1, 2, 3), 2)
	})

	t.Run("val > max", func(t *testing.T) {
		assert.Equal(t, Bounded(3, 1, 2), 2)
	})

	t.Run("val == min", func(t *testing.T) {
		assert.Equal(t, Bounded(1, 1, 2), 1)
	})

	t.Run("val == max", func(t *testing.T) {
		assert.Equal(t, Bounded(2, 1, 2), 2)
	})

	// Floats
	t.Run("val < min", func(t *testing.T) {
		assert.Equal(t, Bounded(1.0, 2.0, 3.0), 2.0)
	})

	t.Run("val > max", func(t *testing.T) {
		assert.Equal(t, Bounded(3.0, 1.0, 2.0), 2.0)
	})

	t.Run("val == min", func(t *testing.T) {
		assert.Equal(t, Bounded(1.0, 1.0, 2.0), 1.0)
	})

	t.Run("val == max", func(t *testing.T) {
		assert.Equal(t, Bounded(2.0, 1.0, 2.0), 2.0)
	})
}

func TestFlatMap(t *testing.T) {

	t.Run("Empty array", func(t *testing.T) {
		assert.Equal(t, FlatMap([]int{}, func(i int) []int { return []int{i} }), []int{})
	})

	t.Run("Non-empty array", func(t *testing.T) {
		assert.Equal(t, FlatMap([]int{1, 2, 3}, func(i int) []int { return []int{i} }), []int{1, 2, 3})
	})

	t.Run("Non-empty array with duplicates", func(t *testing.T) {
		assert.Equal(t, FlatMap([]int{1, 2, 3}, func(i int) []int { return []int{i, i} }), []int{1, 1, 2, 2, 3, 3})
	})
}
