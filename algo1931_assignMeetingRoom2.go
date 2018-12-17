//한 개의 회의실이 있는데 이를 사용하고자 하는 N개의 회의들에 대하여 회의실 사용표를 만들려고 한다.
//각 회의 I에 대해 시작시간과 끝나는 시간이 주어져 있고,
//각 회의가 겹치지 않게 하면서 회의실을 사용할 수 있는 최대수의 회의를 찾아라.
//단, 회의는 한번 시작하면 중간에 중단될 수 없으며 한 회의가 끝나는 것과 동시에 다음 회의가 시작될 수 있다.
//회의의 시작시간과 끝나는 시간이 같을 수도 있다. 이 경우에는 시작하자마자 끝나는 것으로 생각하면 된다.

package main

import (
	"fmt"
	"sort"
)

//Struct 정의
type MeetingRoomTime struct {
	Start_time int
	End_time   int
}

type MeetingRoomTimes []MeetingRoomTime

func (m MeetingRoomTimes) Len() int {
	return len(m) // 데이터 길이 구함
}

func (m MeetingRoomTimes) Less(i, j int) bool {
	if m[i].End_time == m[j].End_time {
		return m[i].Start_time <= m[j].Start_time
	} else {
		return m[i].End_time <= m[j].End_time // 종료시간순 정렬
  }
}

func (m MeetingRoomTimes) Swap(i, j int) {
	m[i], m[j] = m[j], m[i] //두 데이터 위치 바꿈
}

//
// func assignMeetingRoom(mrt MeetingRoomTimes) MeetingRoomTimes {
// 	MeetingRoomTimeList := make(MeetingRoomTimes, 0)
// 	MeetingRoomTimeList = append(MeetingRoomTimeList, mrt[0])
// 	for i := 1; i < mrt.Len(); i++ {
// 		if mrt[i].Start_time >= MeetingRoomTimeList[MeetingRoomTimeList.Len()-1].End_time {
// 			MeetingRoomTimeList = append(MeetingRoomTimeList, mrt[i])
// 		}
// 	}
// 	return MeetingRoomTimeList
// }

// func assignMeetingRoom(mrt MeetingRoomTimes) int {
// 	var result int = 1
// 	MeetingRoomTimeList := make(MeetingRoomTimes, 0)
// 	MeetingRoomTimeList = append(MeetingRoomTimeList, mrt[0])
// 	for i := 1; i < mrt.Len(); i++ {
// 		if mrt[i].Start_time >= MeetingRoomTimeList[MeetingRoomTimeList.Len()-1].End_time {
// 			result++
// 			//			MeetingRoomTimeList = append(MeetingRoomTimeList, mrt[i])
// 		}
// 	}
// 	//	return MeetingRoomTimeList
// 	return result
// }

func main() {
	var N int
	fmt.Scanln(&N)

	var MAX_CNT, limit int = 0, 0

	slice := make(MeetingRoomTimes, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d %d\n", &slice[i].Start_time, &slice[i].End_time)
	}

	sort.Sort(slice)

	for a := 0; a < N; a++ {
		if slice[a].Start_time >= limit {
			MAX_CNT++
			limit = slice[a].End_time
		}
	}
	fmt.Println(MAX_CNT)

	//result := assignMeetingRoom(slice)
	//fmt.Println(result)
	// fmt.Println(result.Len())
	//
	//	fmt.Println(assignMeetingRoom(slice))
}
