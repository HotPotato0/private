//한 개의 회의실이 있는데 이를 사용하고자 하는 N개의 회의들에 대하여 회의실 사용표를 만들려고 한다.
//각 회의 I에 대해 시작시간과 끝나는 시간이 주어져 있고,
//각 회의가 겹치지 않게 하면서 회의실을 사용할 수 있는 최대수의 회의를 찾아라.
//단, 회의는 한번 시작하면 중간에 중단될 수 없으며 한 회의가 끝나는 것과 동시에 다음 회의가 시작될 수 있다.
//회의의 시작시간과 끝나는 시간이 같을 수도 있다. 이 경우에는 시작하자마자 끝나는 것으로 생각하면 된다.

package main

import (
	"fmt"
	"math/rand"
)

//Struct 정의
type MeetingRoomTime struct {
	Start_time int
	End_time   int
}

type MeetingRoomTimes []MeetingRoomTime

func partition(mrt MeetingRoomTimes) int {
	pivot := mrt[0].End_time
	start, end := 0, len(mrt)-1
	for {
		for start < len(mrt) && mrt[start].End_time <= pivot {
			start++
		}
		for mrt[end].End_time > pivot {
			end--
		}
		if start > end {
			mrt[0].End_time, mrt[end].End_time = mrt[end].End_time, mrt[0].End_time
			return end
		} else {
			mrt[start].End_time, mrt[end].End_time = mrt[end].End_time, mrt[start].End_time
		}
	}
}

func quick(mrt MeetingRoomTimes) {
	if len(mrt) > 1 {
		p := partition(mrt)
		quick(mrt[0:p])
		quick(mrt[p+1:])
	}
}

func assignMeetingRoom(mrt MeetingRoomTimes) int {
	var result int = 1
	var end_time = mrt[0].End_time
	for i := 0; i < len(mrt); i++ {
		if mrt[i].Start_time >= end_time {
			result++
			end_time = mrt[i].End_time
		}
	}
	return result
}

func main() {
	var N int
	fmt.Scanln(&N)
	slice := make(MeetingRoomTimes, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d %d\n", &slice[i].Start_time, &slice[i].End_time)
	}
	quick(slice)

	fmt.Println(assignMeetingRoom(slice))
}
