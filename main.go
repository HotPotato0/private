package main

import "fmt"

func main(){
  var Num int = 4
  var pNum *int
  pNum = &Num

  var pNum1 *int = &Num

  fmt.Printf("pNum : %d, Num : %d", pNum, Num)
  fmt.Println()
  fmt.Printf("pNum : %d, Num : %d", pNum1, Num)
  fmt.Println()
  *pNum = 3
  fmt.Printf("pNum : %d, Num : %d", pNum, Num)


}
