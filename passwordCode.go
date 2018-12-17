package main

import (
	"fmt"
	"strconv"
)

func main() {

  var DP[5001] int
	var input int = 25114
	//	fmt.Scanln(&input)

	s := strconv.Itoa(input)
	length := len(s)
  for i := 1 ; i <= length ; i++{
    DP[i] =
    last2 := s[length-2 : length]
    int2, err := strconv.Atoi(last2)
    if int2 <= 26 && int2 >= 11 && int2 != 20 { // 둘다 가능
      DP[i] = DP[i-1]+DP[i-2]
    }else {
      if int2 % 10 == 0{
        return 0
      } else {
        DP[i] = DP[i-2]
      }
    }  // 하나만 가능
  }


	last2 := s[length-2 : length]
	int2, err := strconv.Atoi(last2)

	if int2 > 30 {
		if err == nil {
			fmt.Println(int2)
		}
	}

	fmt.Println(last2)

}
