package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	// "container/heap"
)

func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
}

const FACTORIAL_CACHE_SIZE = 10000000
const MOD = 1000000007

func main() {
	defer flush()
	n := ni()
	a := nis(n)
	flag := true
	for i := 0; i < n-1; i++ {
		if a[i] >= a[i+1] {
			flag = false
		}
	}
	if flag {
		out("Yes")
	} else {
		out("No")
	}
}

// =====================
// utils
// =====================
// io
var sc = bufio.NewScanner(os.Stdin)
var rdr = bufio.NewReader(os.Stdin)
var wtr = bufio.NewWriter(os.Stdout)

// =====================
// input utils
// =====================
// ni reads a single integer from stdin.
func ni() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

// nis reads n integers from stdin with optional offset.
func nis(n int, offset ...int) []int {
	off := 0
	if len(offset) > 0 {
		off = offset[0]
	}

	a := make([]int, n+off)
	for i := off; i < n+off; i++ {
		a[i] = ni()
	}
	return a
}

var bufBytes []byte
var bufIdx int

// =====================
// output utils
// =====================
// flush flushes the buffered writer.
func flush() {
	e := wtr.Flush()
	if e != nil {
		panic(e)
	}
}


// out writes the output to stdout.
func out(v ...interface{}) {
	_, e := fmt.Fprintln(wtr, v...)
	if e != nil {
		panic(e)
	}
}
