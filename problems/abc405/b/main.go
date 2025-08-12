package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
}

// N, M
// A = (A1, A2, .., An)
// M
// 1 <= M <= N <= 100
// 1 <= Ai <= N

func main() {
	defer flush()
	n := ni()
	m := ni()
	a := nis(n)

	amap := make(map[int]int, m)
	all := false
	res := 0

	for i := 0; i < n; i++ {
		_, ok := amap[a[i]]
		if !ok {
			amap[a[i]] = i
		}

		if len(amap) == m {
			all = true
			break
		}
	}

	if all {
		for _, v := range amap {
			if res < v {
				res = v
			}
		}
		res = n - res
	}

	out(res)
}

// io
var sc = bufio.NewScanner(os.Stdin)
var wtr = bufio.NewWriter(os.Stdout)

func ni() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func nis(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = ni()
	}
	return a
}

func flush() {
	e := wtr.Flush()
	if e != nil {
		panic(e)
	}
}

func out(v ...interface{}) {
	_, e := fmt.Fprintln(wtr, v...)
	if e != nil {
		panic(e)
	}
}
