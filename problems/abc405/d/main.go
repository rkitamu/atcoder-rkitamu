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

// H height, i index
// W width, j index
// S(i, j) ., # or E
// d(i, j) shortest distance from (i, j) to E

// 2 <= H <= 1000
// 2 <= W <= 1000

func main() {
	defer flush()
	h := ni()
	w := ni()
	s := nrs2d(h, w)

	type Pos struct {i, j int}
	posQueue := make([]*Pos, 0)

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if s[i][j] == 'E' {
				posQueue = append(posQueue, &Pos{i, j})
			}
		}
	}

	for len(posQueue) > 0 {
		p := posQueue[0]
		posQueue = posQueue[1:]
		for _, d := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			i, j := p.i+d[0], p.j+d[1]
			if i < 0 || i >= h || j < 0 || j >= w {
				continue
			}
			if s[i][j] != '.' {
				continue
			}
			arrow := ' '
			if d[0] == -1 {
				arrow = 'v'
			} else if d[0] == 1 {
				arrow = '^'
			} else if d[1] == -1 {
				arrow = '>'
			} else if d[1] == 1 {
				arrow = '<'
			}
			s[i][j] = arrow
			posQueue = append(posQueue, &Pos{i, j})
		}
	}

	outr2d(s)
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

// ni reads n integers from stdin.
func nis(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = ni()
	}
	return a
}

var bufRunes []rune
var bufIdx int

func nr() rune {
	for {
		if bufIdx < len(bufRunes) {
			r := bufRunes[bufIdx]
			bufIdx++
			return r
		}

		if !sc.Scan() {
			panic("failed to scan next token")
		}
		bufRunes = []rune(sc.Text())
		bufIdx = 0
	}
}

/* なんかtest.sh実行時だけエラーでる
// nr reads a single rune from stdin.
func nr() rune {
	for {
		r, _, err := rdr.ReadRune()
		if err != nil {
			panic(err)
		}
		if r != '\n' && r != '\r' {
			return r
		}
	}
}*/

// nr reads n runes from stdin.
func nrs(n int) []rune {
	a := make([]rune, n)
	for i := 0; i < n; i++ {
		a[i] = nr()
	}
	return a
}

// nrs2d reads n * m runes from stdin.
func nrs2d(n, m int) [][]rune {
	a := make([][]rune, n)
	for i := 0; i < n; i++ {
		a[i] = nrs(m)
	}
	return a
}

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

// outr2d writes a 2D slice of runes to stdout.
func outr2d(a [][]rune) {
	for _, r := range a {
		_, e := fmt.Fprintln(wtr, string(r))
		if e != nil {
			panic(e)
		}
	}
}
