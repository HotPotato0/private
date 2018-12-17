package main

import (
   "fmt"
)


func main() {
  var N int
  fmt.Scanln(&N)
  iqTest := make([]int, N)

  for i := 0 ; i < N ; i++ {
    fmt.Scanf("%d", &iqTest[i])
  }
  var d1, d2 int

  if N == 1 {
    fmt.Println("A")
  } else if N == 2{
    if iqTest[0] ==  iqTest[1]  {
      fmt.Println(iqTest[0])
    } else {
      fmt.Println("A")
    }
  } else if N >= 3 {
    d1 = iqTest[1] - iqTest[0]
    d2 = iqTest[2] - iqTest[1]
    if d1 != 0 {
        a := d2/d1
        b := iqTest[1] - a*iqTest[0]
        isValid(iqTest, a, b)
    } else {
      isValid(iqTest, 1, 0)
    }
  }
}


func isValid(input []int, a, b int) {
  for ii := 0 ; ii < len(input)-1; ii++{
    if input[ii+1] != input[ii]*a + b {
      fmt.Println("B")
      return
    }
  }
  fmt.Println(input[len(input)-1]*a+b)
}
