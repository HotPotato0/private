package main

import "fmt"

func main() {
  var input int
  var result int = 0

  for i := 5 ; i > 0 ; i--{
    fmt.Scanln(&input)
    result += getAverage(input)
  }
  fmt.Println(result/5)
}

func getAverage(score int)(int){
  if score < 40{
    score = 40
  }
  return score
}
