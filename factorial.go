package main

import "fmt"

func main() {
	var value int
	var result int = 1
	fmt.Scanln(&value)

	//  fmt.Println(getFact(value))
	tmp := value
	for tmp > 0 {
		result *= tmp
		tmp--
	}
	fmt.Println(result)

}

/*
func getFact var(a *int){
  if a > 1
    return a * getFact(a-1)
  else
   return 1;
}
*/
