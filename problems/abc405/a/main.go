package main

import "fmt"

// 1 <= r <= 4229
// 1 <= x <= 2
// r x

func main() {
	var r, x int
	fmt.Scan(&r, &x)

	res := ""
	if x == 1 {
		if 1600 <= r && r <= 2999 {
			res = "Yes"
		} else {
			res = "No"
		}
	} else {
		if 1200 <= r && r <= 2399 {
			res = "Yes"
		} else {
			res = "No"
		}
	}
	
	fmt.Println(res)
}
