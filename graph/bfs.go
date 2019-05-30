package graph

import "fmt"

var ve = [5][5]int{
	{0, 1, 0, 1, 1},
	{1, 0, 1, 0, 0},
	{0, 1, 0, 1, 0},
	{1, 0, 1, 0, 1},
	{1, 0, 0, 1, 0}}

var p [5]int
var d = [5]int{0}

func bfs() {
	s := make([]int, 0, 100)
	s = append(s, 0)
	c := 0
	d[0] = -1
	for {
		if len(s) == 0 {
			break
		}
		c++
		for _, x := range s {
			for i, j := range ve[x] {
				if j == 1 {
					if d[i] != 0 {
						continue
					}

					d[i] = c
					s = append(s, i)
				}
			}
		}
		s = s[1:]
	}
	fmt.Println(d)
}
