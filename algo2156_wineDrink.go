package main

import (
   "fmt"
)

func main() {
  var N int
  fmt.Scan(&N)
  wine := make([]int, N)

  for i := range wine {
    fmt.Scan(&wine[i])
  }

  dp := make([]int, N)
  dp[0] = wine[0]
  if N == 1{

  } else if N == 2 {
    dp[1] = wine[1] + wine[0]
  } else {
    dp[1] = dp[0] + wine[1]
    dp[2] = isBigger(dp[1], dp[0]+wine[2], wine[1]+wine[2])
    for i := 3 ; i < N ; i++ {
      dp[i] = isBigger(dp[i-1], dp[i-2]+wine[i], dp[i-3]+wine[i-1]+wine[i])
    }
  }
  fmt.Println( dp[N-1] )
}


func isBigger(n int, ns ...int) int {
	x := n
	for _, v := range ns {
		if v > x {
			x = v
		}
	}
	return x
}

// func drinkWine(wine []int) {
//   dp := make([]int, len(wine))
//   dp[0] = wine[0]
//   if len(wine) == 1{
//
//   } else if len(wine) == 2 {
//     dp[1] = wine[1] + wine[0]
//   } else {
//     dp[1] = dp[0] + wine[1]
//     dp[2] = isBigger(dp[1], dp[0]+wine[2], wine[1]+wine[2])
//     for i := 3 ; i < len(wine) ; i++ {
//       dp[i] = isBigger(dp[i-1], dp[i-2]+wine[i], dp[i-3]+wine[i-1]+wine[i])
//     }
//   }
//   fmt.Println( dp[len(wine)-1] )
// }
