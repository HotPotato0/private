//첫째 줄에 나무의 수 N과 상근이가 집으로 가져가려고 하는 나무의 길이 M이 주어진다. (1 ≤ N ≤ 1,000,000, 1 ≤ M ≤ 2,000,000,000)
//둘째 줄에는 나무의 높이가 주어진다. 나무의 높이의 합은 항상 M을 넘기 때문에, 상근이는 집에 필요한 나무를 항상 가져갈 수 있다. 높이는 1,000,000,000보다 작거나 같은 양의 정수 또는 0이다.


package main

import (
   "bufio"
   "fmt"
   "os"
)

type tree struct{
  total, get int }

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
  var wantTree = nextInt()
  t := make([]tree, N)
  max := 0

  for i := range t {
    t[i].total = nextInt()
    if max < t[i].total {
      max = t[i].total
    }
  }

  cutTree(t, wantTree, max)
}


func cutTree(Tree []tree, wantTree int, max int){
  var cutLength, getTree int
  min := 0

  for min < max {
    getTree = 0
    cutLength = (max + min + 1) / 2

    for i := range Tree {
      if Tree[i].total > cutLength {
        Tree[i].get = Tree[i].total - cutLength
        getTree += Tree[i].get
      }
    }
    if getTree < wantTree {
      max = cutLength - 1
//       fmt.Println(getTree, wantTree)
    } else {
      min = cutLength
    }
  }
  fmt.Println(min)
}
