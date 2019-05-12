package bruteforce

import (
	"errors"
	"sort"
)

//求集合的子集，可以用类似求全排列的方法，每个排列只有一个从小到大的一种集合。
//也可以用二进制的方法

//subset 求集合s的所有子集
//s中没有重复的元素
func subset(s []byte) [][]byte {
	//先将s排序
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})

	prefix := make([]byte, 0, len(s))
	result := make([][]byte, 0, 0)
	subsetImpl(s, prefix, &result)
	return result
}

func subsetImpl(s []byte, prefix []byte, result *[][]byte) {
	//copy prefix
	r := make([]byte, len(prefix), len(prefix))
	copy(r, prefix)
	*result = append(*result, r)
	if len(prefix) == len(s) {
		return
	}

	//选择一个元素加入prefix，这个元素比prefix现存的元素都大
	n := sort.Search(len(s), func(i int) bool {
		if len(prefix) < 1 {
			return true
		}
		return s[i] > prefix[len(prefix)-1]
	})

	for i := n; i < len(s); i++ {
		prefix = append(prefix, s[i])
		subsetImpl(s, prefix, result)
		//注意要把改变了的prefix还原
		prefix = prefix[0 : len(prefix)-1]
	}
}

//subsetBinary 二进制求解集合子集
func subsetBinary(s []byte) ([][]byte, error) {
	if len(s) > 32 {
		return nil, errors.New("len of input larger than 32")
	}
	end := 1 << uint(len(s))
	result := make([][]byte, 0)
	for i := 0; i < end; i++ {
		t := make([]byte, 0)
		for j := uint(0); j < uint(len(s)); j++ {
			if (i & (1 << j)) != 0 {
				t = append(t, s[j])
			}
		}
		result = append(result, t)
	}
	return result, nil
}
