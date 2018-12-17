package main

import "fmt"

func main() {
	var input int
	fmt.Scanln(&input)

	fmt.Println(getPrinaryNum(input))
}

func getPrinaryNum(n int) int {
	if n == 1 {
		return 1
	}
	var left, right int
	left = 0
	right = 1
	var temp int
	var cntPrinaryNum int = 0
	for i := 1; i <= n-1; i++ {
		temp = right
		right = left
		left = left + temp
		cntPrinaryNum = left + right
	}
	return cntPrinaryNum
}
