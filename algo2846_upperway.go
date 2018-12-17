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
  upperWay := make([]int, N)

  for i := range upperWay {
    upperWay[i] = nextInt()
  }

  MaxUpperWaySize := 0
  upperWaySize := 0
  for ix := 0 ; ix < N-1 ; ix++ {
    if isHigher(upperWay[ix], upperWay[ix+1]) {
      upperWaySize += upperWay[ix+1] - upperWay[ix]
      if MaxUpperWaySize <= upperWaySize {
        MaxUpperWaySize = upperWaySize
      }
    } else {
      upperWaySize = 0
    }
  }
  fmt.Println(MaxUpperWaySize)
}

func isHigher(a, b int) bool {
  if a < b{
    return true
  } else {
    return false
  }
}
