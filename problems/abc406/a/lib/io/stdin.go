package stdio

import (
	"bufio"
	"os"
	"strconv"
	"math"
)

// =====================
// input utils
// =====================
// ni reads a single integer from stdin.
func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
}

var sc = bufio.NewScanner(os.Stdin)
var rdr = bufio.NewReader(os.Stdin)

func Ni() int {
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
		a[i] = Ni()
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

// nr reads a single rune from stdin.
func Ns() []rune {
	if !sc.Scan() {
		panic("failed to scan next token")
	}
	return []rune(sc.Text())
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
