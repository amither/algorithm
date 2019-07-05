package probability

import (
	"fmt"
	"math/rand"
	"time"
)

// gen_random_without_repeating 生成n个在[0,m)范围内的整数
func gen_random_without_repeating(m int, n int) {
	rand.Seed(time.Now().UnixNano())
	c := make([]int, m, m)
	for i, _ := range c {
		c[i] = i
	}

	r := make([]int, 0, n)

	for i := 0; i < n; i++ {
		k := rand.Intn(m)
		r = append(r, c[k])
		c[k] = c[m-1]
		m--
	}

	for _, k := range r {
		fmt.Printf("%d ", k)
	}
	fmt.Println()
}
