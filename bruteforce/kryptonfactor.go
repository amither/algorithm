package bruteforce

func factordfs(n int, l int) []int {
	s := make([]int, 0)
	factordfsImpl(n, l, &s)
	return s
}

func factordfsImpl(n int, l int, s *[]int) int {
	if len(*s) == n {
		return 0 //找到解就直接退出
	}

	for i := 0; i < l; i++ {
		*s = append(*s, i)

		//检查是否包含重复子串
		has := false
		for j := 1; j*2 <= len(*s); j++ {
			for k := 0; k < j; k++ {
				if (*s)[len(*s)-k-1] != (*s)[len(*s)-j-k-1] {
					break
				}
				has = true
			}
			if has == true {
				break
			}
		}

		if has {
			//有重复子串，回溯
			*s = (*s)[0 : len(*s)-1]
		} else {
			if 0 == factordfsImpl(n, l, s) {
				return 0
			}
		}
	}

	return 1
}
