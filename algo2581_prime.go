//소수 구하는 문제
//자연수 M과 N이 주어질 때 M이상 N이하의 자연수 중 소수인 것을 모두 골라 이들 소수의 합과 최솟값을 찾는 프로그램을 작성하시오.
package main

import "fmt"

func main() {
	var M, N int
	fmt.Scanln(&M)
	fmt.Scanln(&N)

	var Sum, Minimum int = 0, 0

	for i := M; i <= N; i++ {
		if isPrime(i) == true {
			if Minimum == 0 {
				Minimum = i
			}
			Sum += i
		}
	}
	if Minimum == 0 {
		fmt.Println("-1")
	} else {
		fmt.Println(Sum)
		fmt.Println(Minimum)
	}
}

func isPrime(number int) bool {
	if number <= 1 {
		return false
	}
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}
