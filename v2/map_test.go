package v2

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeys(t *testing.T) {
	is := assert.New(t)

	results := Keys(map[string]int{"one": 1, "two": 2})
	sort.Strings(results)

	is.Equal(results, []string{"one", "two"})

	fields := Keys(map[int]int{1: 1, 2: 2})

	sort.Ints(fields)

	is.Equal(fields, []int{1, 2})
}

func TestValues(t *testing.T) {
	is := assert.New(t)

	results := Values(map[string]int{"one": 1, "two": 2})
	sort.Ints(results)

	is.Equal(results, []int{1, 2})

	values := Values(map[int]string{1: "one", 2: "two"})

	is.Equal(values, []string{"one", "two"})
}
