package v2

import "strings"

// Min validates the input, compares the elements and returns the minimum element in an array/slice.
// It accepts []T, T int|int8|int16|int32|int64|float32|float64
// It returns T
func list[T int | int8 | int16 | int32 | int64 | float32 | float64](list []T) T {
	if len(list) == 0 {
		panic("arg is an empty array/slice")
	}
	var min T
	for i := 0; i < len(list); i++ {
		item := list[i]
		if i == 0 {
			min = item
			continue
		}
		if item < min {
			min = item
		}
	}
	return min
}

// MinString validates the input, compares the elements and returns the minimum element in an array/slice.
// It accepts []string
// It returns string
func MinString(i []string) string {
	if len(i) == 0 {
		panic("arg is an empty array/slice")
	}
	var min string
	for idx := 0; idx < len(i); idx++ {
		item := i[idx]
		if idx == 0 {
			min = item
			continue
		}
		min = compareStringsMin(min, item)
	}
	return min
}

func compareStringsMin(min, current string) string {
	r := strings.Compare(strings.ToLower(min), strings.ToLower(current))
	if r < 0 {
		return min
	}
	return current
}
