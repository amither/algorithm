package bruteforce

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubset(t *testing.T) {
	type a struct {
		input  []byte
		expect [][]byte
	}

	A := []a{
		{[]byte{'a', 'b'}, [][]byte{{}, {'b'}, {'a'}, {'a', 'b'}}},
	}

	for _, v := range A {
		result := subset(v.input)
		resultBinary, _ := subsetBinary(v.input)
		sort.Sort(ByValue(v.expect))
		sort.Sort(ByValue(result))
		sort.Sort(ByValue(resultBinary))
		assert.Equal(t, v.expect, result)
		assert.Equal(t, v.expect, resultBinary)
	}
}
