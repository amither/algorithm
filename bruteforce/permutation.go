package bruteforce

//permutation 求集合s的全排列
func permutation(s []byte) [][]byte {
	cur := make([]byte, 0, len(s))
	result := make([][]byte, 0)
	vis := make([]bool, len(s), len(s))
	permutationImpl(s, cur, vis, &result)
	return result
}

func permutationImpl(s []byte, cur []byte, vis []bool, result *[][]byte) {
	if len(cur) == len(s) {
		//copy cur
		t := make([]byte, len(cur), len(cur))
		copy(t, cur)
		*result = append(*result, t)
		return
	}

	//找下一个要加入到cur中的元素
	//遍历vis，第一个为非true的下标
	for i, v := range vis {
		if v == false {
			cur = append(cur, s[i])
			vis[i] = true
			permutationImpl(s, cur, vis, result)
			cur = cur[0 : len(cur)-1]
			vis[i] = false
		}
	}
}
