//첫째 줄에 정수 N이 주어진다. 다음 N개의 줄에는 각 로프가 버틸 수 있는 최대 중량이 주어진다. 이 값은 10,000을 넘지 않는다.
package main

import (
   "bufio"
   "fmt"
   "os"
   "sort"
)

type rope struct{
  weight,s int }

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
  d := make([]rope, N)

  for i := range d {
    d[i].weight = nextInt()
  }
  sort.Slice(d, func(i, j int) bool {
    return d[i].weight < d[j].weight
  })
   Max := 0
   for ix, d := range d {
      if Max < (N-ix)*d.weight {
        Max = (N-ix)*d.weight
      }
   }
   fmt.Println(Max)
}
