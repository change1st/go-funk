package v2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase[T comparable] struct {
	LeftArr  []T
	RightArr []T
	Expect   []T
}

func TestJoin_InnerJoin(t *testing.T) {
	strTestCases := []TestCase[string]{
		{[]string{"foo", "bar"}, []string{"bar", "baz"}, []string{"bar"}},
		{[]string{"foo", "bar", "bar"}, []string{"bar", "baz"}, []string{"bar"}},
		{[]string{"foo", "bar"}, []string{"bar", "bar", "baz"}, []string{"bar"}},
		{[]string{"foo", "bar", "bar"}, []string{"bar", "bar", "baz"}, []string{"bar"}},
	}
	intTestCases := []TestCase[int]{
		{[]int{0, 1, 2, 3, 4}, []int{3, 4, 5, 6, 7}, []int{3, 4}},
	}
	structTestCases := []TestCase[*Foo]{
		{[]*Foo{f, b}, []*Foo{b, c}, []*Foo{b}},
	}

	for idx, tt := range strTestCases {
		t.Run(fmt.Sprintf("test case #%d", idx+1), func(t *testing.T) {
			is := assert.New(t)

			actual := Join(tt.LeftArr, tt.RightArr, InnerJoin[string])
			is.Equal(tt.Expect, actual)
		})
	}
	for idx, tt := range intTestCases {
		t.Run(fmt.Sprintf("test case #%d", idx+1), func(t *testing.T) {
			is := assert.New(t)

			actual := Join(tt.LeftArr, tt.RightArr, InnerJoin[int])
			is.Equal(tt.Expect, actual)
		})
	}
	for idx, tt := range structTestCases {
		t.Run(fmt.Sprintf("test case #%d", idx+1), func(t *testing.T) {
			is := assert.New(t)

			actual := Join(tt.LeftArr, tt.RightArr, InnerJoin[*Foo])
			is.Equal(tt.Expect, actual)
		})
	}
}

func TestJoin_OuterJoin(t *testing.T) {
	strTestCases := []TestCase[string]{
		{[]string{"foo", "bar"}, []string{"bar", "baz"}, []string{"foo", "baz"}},
	}
	intTestCases := []TestCase[int]{
		{[]int{0, 1, 2, 3, 4}, []int{3, 4, 5, 6, 7}, []int{0, 1, 2, 5, 6, 7}},
	}
	fooTestCases := []TestCase[*Foo]{
		{[]*Foo{f, b}, []*Foo{b, c}, []*Foo{f, c}},
	}

	for idx, tt := range strTestCases {
		t.Run(fmt.Sprintf("test case #%d", idx+1), func(t *testing.T) {
			is := assert.New(t)

			actual := Join(tt.LeftArr, tt.RightArr, OuterJoin[string])
			is.Equal(tt.Expect, actual)
		})
	}
	for idx, tt := range intTestCases {
		t.Run(fmt.Sprintf("test case #%d", idx+1), func(t *testing.T) {
			is := assert.New(t)

			actual := Join(tt.LeftArr, tt.RightArr, OuterJoin[int])
			is.Equal(tt.Expect, actual)
		})
	}
	for idx, tt := range fooTestCases {
		t.Run(fmt.Sprintf("test case #%d", idx+1), func(t *testing.T) {
			is := assert.New(t)

			actual := Join(tt.LeftArr, tt.RightArr, OuterJoin[*Foo])
			is.Equal(tt.Expect, actual)
		})
	}
}

func TestJoin_LeftJoin(t *testing.T) {
	strTestCases := []TestCase[string]{
		{[]string{"foo", "bar"}, []string{"bar", "baz"}, []string{"foo"}},
	}
	intTestCases := []TestCase[int]{
		{[]int{0, 1, 2, 3, 4}, []int{3, 4, 5, 6, 7}, []int{0, 1, 2}},
	}
	fooTestCases := []TestCase[(*Foo)]{
		{[]*Foo{f, b}, []*Foo{b, c}, []*Foo{f}},
	}

	for idx, tt := range strTestCases {
		t.Run(fmt.Sprintf("test case #%d", idx+1), func(t *testing.T) {
			is := assert.New(t)

			actual := Join(tt.LeftArr, tt.RightArr, LeftJoin[string])
			is.Equal(tt.Expect, actual)
		})
	}
	for idx, tt := range intTestCases {
		t.Run(fmt.Sprintf("test case #%d", idx+1), func(t *testing.T) {
			is := assert.New(t)

			actual := Join(tt.LeftArr, tt.RightArr, LeftJoin[int])
			is.Equal(tt.Expect, actual)
		})
	}
	for idx, tt := range fooTestCases {
		t.Run(fmt.Sprintf("test case #%d", idx+1), func(t *testing.T) {
			is := assert.New(t)

			actual := Join(tt.LeftArr, tt.RightArr, LeftJoin[*Foo])
			is.Equal(tt.Expect, actual)
		})
	}
}

func TestJoin_RightJoin(t *testing.T) {
	strTestCases := []TestCase[string]{
		{[]string{"foo", "bar"}, []string{"bar", "baz"}, []string{"baz"}},
	}
	intTestCases := []TestCase[int]{
		{[]int{0, 1, 2, 3, 4}, []int{3, 4, 5, 6, 7}, []int{5, 6, 7}},
	}
	structTestCases := []TestCase[*Foo]{
		{[]*Foo{f, b}, []*Foo{b, c}, []*Foo{c}},
	}

	for idx, tt := range strTestCases {
		t.Run(fmt.Sprintf("test case #%d", idx+1), func(t *testing.T) {
			is := assert.New(t)

			actual := Join(tt.LeftArr, tt.RightArr, RightJoin[string])
			is.Equal(tt.Expect, actual)
		})
	}
	for idx, tt := range intTestCases {
		t.Run(fmt.Sprintf("test case #%d", idx+1), func(t *testing.T) {
			is := assert.New(t)

			actual := Join(tt.LeftArr, tt.RightArr, RightJoin[int])
			is.Equal(tt.Expect, actual)
		})
	}
	for idx, tt := range structTestCases {
		t.Run(fmt.Sprintf("test case #%d", idx+1), func(t *testing.T) {
			is := assert.New(t)

			actual := Join(tt.LeftArr, tt.RightArr, RightJoin[*Foo])
			is.Equal(tt.Expect, actual)
		})
	}
}
