package v2

import "strings"

// Max validates the input, compares the elements and returns the maximum element in an array/slice.
// It accepts []T, T int|int8|int16|int32|int64|float32|float64
// It returns T
func Max[T int | int8 | int16 | int32 | int64 | float32 | float64](list []T) T {
	if len(list) == 0 {
		panic("arg is an empty array/slice")
	}
	var max T
	for i := 0; i < len(list); i++ {
		item := list[i]
		if i == 0 {
			max = item
			continue
		}
		if item > max {
			max = item
		}
	}
	return max
}

// MaxString validates the input, compares the elements and returns the maximum element in an array/slice.
// It accepts []string
// It returns string
func MaxString(i []string) string {
	if len(i) == 0 {
		panic("arg is an empty array/slice")
	}
	var max string
	for idx := 0; idx < len(i); idx++ {
		item := i[idx]
		if idx == 0 {
			max = item
			continue
		}
		max = compareStringsMax(max, item)
	}
	return max
}

// compareStrings uses the strings.Compare method to compare two strings, and returns the greater one.
func compareStringsMax(max, current string) string {
	r := strings.Compare(strings.ToLower(max), strings.ToLower(current))
	if r > 0 {
		return max
	}
	return current
}
