package goutil

type Numeric interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

func UnwrapOr[V any](i *V, fallback V) V {
	if i != nil {
		return *i
	}
	return fallback
}

func Ptr[V any](v V) *V {
	return &v
}

func Outside[V Numeric](val, min, max V) bool {
	return val < min || val > max
}

func Bounded[V Numeric](val, min, max V) V {
	if val < min {
		return min
	}
	if val > max {
		return max
	}
	return val
}

func Map[V any, R any](arr []V, f func(V) R) []R {
	result := make([]R, len(arr))
	for i, v := range arr {
		result[i] = f(v)
	}
	return result
}
