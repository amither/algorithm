package dp

import "fmt"

//表示杨辉三角变形版（每个节点的值被任意的修改过）
var yanghui = [][]int{
	{5},
	{7, 8},
	{2, 3, 4},
	{4, 9, 6, 1},
	{2, 7, 9, 4, 5},
}

// 查找从最顶层到最底层的路径节点和的最小值
func yh() {
	//保存状态, 每一层到当前节点为止的最短长度
	state := make([][]int, len(yanghui), len(yanghui))

	state[0] = append(state[0], yanghui[0][0])
	for i := 1; i < len(yanghui); i++ {
		state[i] = make([]int, len(yanghui[i]), len(yanghui[i]))
		for j := 0; j < len(yanghui[i]); j++ {
			//找这个节点的两个上层节点
			k := j / 2
			v := state[i-1][k] + yanghui[i][j]

			if j != 0 && k+1 < len(yanghui[i-1]) {
				q := state[i-1][k+1] + yanghui[i][j]
				if q < v {
					v = q
				}
			}
			state[i][j] = v
		}
	}

	fmt.Println(state)
}
