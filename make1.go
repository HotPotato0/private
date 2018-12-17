package main

import "fmt"

func main() {
	var input int = 0

	fmt.Scanln(&input)
	var dp [1000001]int
	dp[0] = 0
	dp[1] = 0
	if input > 1 {
		for i := 2; i <= input; i++ {
			dp[i] = dp[i-1] + 1
			if i%2 == 0 {
				dp[i] = min(dp[i], dp[i/2]+1)
			}
			if i%3 == 0 {
				dp[i] = min(dp[i], dp[i/3]+1)
			}
		}
	}
	fmt.Println(dp[input])
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
