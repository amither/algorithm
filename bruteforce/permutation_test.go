package bruteforce

import (
	"bytes"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type ByValue [][]byte

func (a ByValue) Len() int {
	return len(a)
}
func (a ByValue) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByValue) Less(i, j int) bool {
	return bytes.Compare(a[i], a[j]) < 0
}

func TestPermutation(t *testing.T) {
	type a struct {
		input  []byte
		expect [][]byte
	}

	A := []a{
		{[]byte{'a', 'b'}, [][]byte{{'a', 'b'}, {'b', 'a'}}},
		{[]byte{'a', 'b', 'c'}, [][]byte{{'a', 'b', 'c'}, {'a', 'c', 'b'},
			{'b', 'c', 'a'}, {'b', 'a', 'c'}, {'c', 'a', 'b'}, {'c', 'b', 'a'}}},
	}

	for _, v := range A {
		result := permutation(v.input)
		sort.Sort(ByValue(v.expect))
		sort.Sort(ByValue(result))
		assert.Equal(t, v.expect, result)
	}
}
