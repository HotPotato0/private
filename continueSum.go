package main

import (
	"fmt"
)

func main() {
	var dp [100001]int
	var num [100001]int
	var inputNum int
	fmt.Scanln(&inputNum)

	for i := 0; i < inputNum; i++ {
		fmt.Scanf("%d", &num[i])
	}

	dp[0] = num[0]
	maxSum := dp[0]
	for i := 1; i < inputNum; i++ {
		dp[i] = getBigger(dp[i-1]+num[i], num[i])
		maxSum = getBigger(dp[i], maxSum)
	}
	fmt.Println(maxSum)

}

func getBigger(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
