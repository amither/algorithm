package bruteforce

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactor(t *testing.T) {
	type a struct {
		n      int
		l      int
		expect []int
	}

	A := []a{
		{2, 2, []int{0, 1}},
	}

	for _, v := range A {
		result := factordfs(v.n, v.l)

		assert.Equal(t, v.expect, result)
	}
}
