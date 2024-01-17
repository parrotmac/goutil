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

func TestPtr(t *testing.T) {

	// Test int
	t.Run("int", func(t *testing.T) {
		assert.Equal(t, *Ptr(1), 1)
	})

	// Test float64
	t.Run("float64", func(t *testing.T) {
		assert.Equal(t, *Ptr(1.0), 1.0)
	})

	// Test struct
	t.Run("struct", func(t *testing.T) {
		type S struct {
			A int
		}
		assert.Equal(t, *Ptr(S{1}), S{1})
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

func TestMap(t *testing.T) {

	t.Run("int", func(t *testing.T) {
		assert.Equal(t, Map([]int{1, 2, 3}, func(v int) int { return v * 2 }), []int{2, 4, 6})
	})

	t.Run("float64", func(t *testing.T) {
		assert.Equal(t, Map([]float64{1.0, 2.0, 3.0}, func(v float64) float64 { return v * 2.0 }), []float64{2.0, 4.0, 6.0})
	})

	t.Run("struct", func(t *testing.T) {
		type S struct {
			A int
		}
		assert.Equal(t, Map([]S{{1}, {2}, {3}}, func(v S) int { return v.A * 2 }), []int{2, 4, 6})
	})
}
