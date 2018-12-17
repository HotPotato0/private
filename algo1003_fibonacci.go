//N이 주어졌을 때, fibonacci(N)을 호출했을 때, 0과 1이 각각 몇 번 출력되는지 구하는 프로그램을 작성하시오.
package main

import "fmt"

func main() {
	var N, input int
	fmt.Scanln(&N)

	for ix := 1; ix <= N; ix++ {
		fmt.Scanln(&input)
		dp0 := make([]int, input+2)
		dp1 := make([]int, input+2)

		dp0[0] = 1
		dp0[1] = 0
		dp1[0] = 0
		dp1[1] = 1
		if input > 1 {
			for i := 2; i <= input; i++ {
				dp0[i] = dp0[i-1] + dp0[i-2]
				dp1[i] = dp1[i-1] + dp1[i-2]
			}
		}
		fmt.Printf("%d %d\n", dp0[input], dp1[input])
	}
}
