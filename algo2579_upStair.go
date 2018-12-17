package main

import "fmt"


func main() {
  var N int
  fmt.Scan(&N)
  stair := make([]int, 301)

  for i := 0 ; i < N ; i++ {
    fmt.Scanln(&stair[i])
  }
  dp := make([]int, 301)
  dp[0] = stair[0]
  dp[1] = getMax(stair[1], stair[0]+stair[1])
  dp[2] = getMax(stair[0]+stair[2], stair[1]+stair[2])
  for i := 3 ; i < N ; i++ {
    dp[i] = getMax(dp[i-2]+stair[i], dp[i-3]+stair[i]+stair[i-1])
  }
  fmt.Println(dp[N-1])

}

func getMax(a, b int) int {
  if a >= b {
    return a
  } else {
    return b
  }
}
