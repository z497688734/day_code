package main

import "fmt"

func getTheNum(num int) int {
	var dp [5][10000]int
	moneys := [...]int{1, 5, 10 ,20 , 50, 100}// 面额数组
	for i := 0; i < 5; i++ { // 用前i种面额拼凑1rmb的方法数均为1
		dp[i][0] = 1
	}
	for i := 0; i <= num; i++ { // 用1rmb面额拼凑0金额的方法数都为1
		dp[0][i] = 1
	}

	for i := 1; i < 5; i++ { // 每一种面额组合递推
		for j := 1; j <= num; j++ {
			if j - moneys[i] >= 0 { // 当当前金额大于这次循环的面额值，则组合数等于当前i种面额拼凑j金额的组合数+前i+1种面额拼凑j - moneys[i]金额的组合数
				dp[i][j] = dp[i-1][j] + dp[i][j - moneys[i]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[4][num] // 返回最后一项
}

func main() {
	a:= getTheNum(5)
	fmt.Println(a)
}