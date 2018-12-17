package main

import "fmt"

func main() {
	var x, y, w, h int

	fmt.Scanf("%d %d %d %d", &x, &y, &w, &h)
	fmt.Println(getMin(getMin(w-x, x), getMin(h-y, y)))
}

func getMin(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
