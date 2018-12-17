package main

import "fmt"

func main() {
   var N, endTime int
   var inputStartTime, inputEndTime [100000]int
   var time []int // timetable i: 시작시간 / v: 종료시간(없을 경우 -1)
   var isUsedTime []bool

   // 입력
   fmt.Scanln(&N)
   for i := 0; i < N; i++ {
      fmt.Scanf("%d %d\n", &inputStartTime[i], &inputEndTime[i])

      if endTime < inputEndTime[i] {
         endTime = inputEndTime[i]
      }
   }

   // INIT time, isUsedTime
   // time - i: 시작시간 / v: 종료시간(없을 경우 -1)
   time = make([]int, endTime)
   isUsedTime = make([]bool, endTime)
   // INIT
   for i := range time {
      time[i] = -1
   }
   for i := 0; i < N; i++ {
      v := inputStartTime[i]
      if time[v] < 0 || time[v] > inputEndTime[i] {
         time[v] = inputEndTime[i]
      }
   }

   // count 방 개수
   result := 0
   for i := endTime - 1; i >= 0; i-- {
      if isAvail(i, time[i], isUsedTime) {
         result++
      }
   }

   fmt.Println(result)
}

func isAvail(start int, end int, isUsedTime []bool) bool {
   // 해당 시간에 사용신청한 사람 없음
   if end < start {
      return false
   }
   // 원하는 회의 시간에 이미 사용되고 있는지 검사
   for i := start; i < end; i++ {
      if isUsedTime[i] {
         return false
      }
   }
   // 사용가능하므로 true로 표시
   for i := start; i < end; i++ {
      isUsedTime[i] = true
   }
   return true
}
