package main

import (
	"strconv"
	"rkitamu/contest/lib/io"
)

func main() {
	defer stdio.Flush()
	a := stdio.Ns()
	b := stdio.Ns()
	c := stdio.Ns()
	d := stdio.Ns()
	pad := func(s []rune) []rune {
		if len(s) == 1 {
			return append([]rune{'0'}, s...)
		}
		return s
	}
	a = pad(a)
	b = pad(b)
	c = pad(c)
	d = pad(d)

	ab, _ := strconv.Atoi(string(a) + string(b))
	cd, _ := strconv.Atoi(string(c) + string(d))

	ans := "No"
	if ab > cd {
		ans = "Yes"
	}
	stdio.Out(ans)
}
