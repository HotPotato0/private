package main

import (
   "bufio"
   "fmt"
   "os"
)

var in = bufio.NewScanner(os.Stdin)

func nextInt() int {
   in.Scan()
   r := 0
   for _, c := range in.Bytes() {
      r *= 10
      r += int(c - '0')
   }
   return r
}

func main() {
  in.Split(bufio.ScanWords)
  var N = nextInt()
  iqTest := make([]int, N)

  for i := range iqTest {
    fmt.Scanf("%d", &iqTest[i])
  }

  first := iqTest[0]
  var d1, d2 int

  if N == 1 {
    fmt.Println("A")
  } else if N == 2{
    if iqTest[0] == iqTest[1]  {
      fmt.Println(iqTest[0])
    } else {
    fmt.Println("A")
  }
  } else if isValid(iqTest) {
    d1 = iqTest[1] - iqTest[0]
    d2 = iqTest[2] - iqTest[1]
    if d1 == 0 && d2 == 0{
      fmt.Println(iqTest[0])
    } else {
    a := d2 / d1
    result := first
    for ix := 0 ; ix < N ; ix++{
      result += d1
      d1 = d1*a
    }
    fmt.Println(result)
  }
  } else {
    fmt.Println("B")
  }
}


func isValid(input []int) bool {
  var d1 int
  d1 = input[1] - input[0]
  if d1 == 0 {
    for ii := 0 ; ii < len(input) - 1 ; ii++{
        if input[ii+1] - input[ii] != 0 {
          return false
          }
        }
        return true
  } else {
  var a int
  a = (input[2] - input[1])/(input[1] - input[0])
  for ix := 0 ; ix < len(input)-2 ; ix++{
    if (input[ix+2] - input[ix+1])/(input[ix+1] - input[ix]) != a {
      return false
    }
    if (input[ix+2] - input[ix+1])%(input[ix+1] - input[ix]) != 0 {
      return false
    }
  }
  }
  return true

}
