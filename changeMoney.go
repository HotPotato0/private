package main

import "fmt"

func main() {
	var price int
	fmt.Scanln(&price)
	var coinNum int = 0
	var change int = 1000 - price
	coinNum += change / 500
	change = change % 500
	coinNum += change / 100
	change = change % 100
	coinNum += change / 50
	change = change % 50
	coinNum += change / 10
	change = change % 10
	coinNum += change / 5
	change = change % 5
	coinNum += change / 1
	change = change % 1

	fmt.Println(coinNum)
}

/*
func giveChange(price int) int {
	var coinNum int = 0
	var change int = 1000 - price
	coinNum += change / 500
	change = change % 500
	coinNum += change / 100
	change = change % 100
	coinNum += change / 50
	change = change % 50
	coinNum += change / 10
	change = change % 10
	coinNum += change / 5
	change = change % 5
	coinNum += change / 1
	change = change % 1

	return coinNum
}
*/
