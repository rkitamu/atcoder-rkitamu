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
	// if use Combination
	initCombTable()
}

const MAX_COMB_CACHE = 10000000
const MOD = 998244353

func main() {
	defer flush()
		in := nis(4)
	a := in[0]
	b := in[1]
	c := in[2]
	d := in[3]
	ans := 0
	for i := 0; i <= b; i++ {
		ai := combMod(a-1+i, i)
		bdc := combMod((b-i+d)+c, c)
		ans += ai * bdc % MOD
	}
	out(ans % MOD)
}

// =====================
// Portions of this file are based on code from:
// https://github.com/gosagawa/atcoder
// Copyright (c) gosagawa
// Licensed under the MIT License
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

// =====================
// calculation utils
// =====================
var fact, invFact []int

// combination calculates nCk
func initCombTable() {
	fact = make([]int, MAX_COMB_CACHE+1)
	invFact = make([]int, MAX_COMB_CACHE+1)
	fact[0] = 1
	for i := 1; i <= MAX_COMB_CACHE; i++ {
		fact[i] = fact[i-1] * i % MOD
	}
	invFact[MAX_COMB_CACHE] = powMod(fact[MAX_COMB_CACHE], MOD-2)
	for i := MAX_COMB_CACHE - 1; i >= 0; i-- {
		invFact[i] = invFact[i+1] * (i + 1) % MOD
	}
}

// comb calculates nCk
func combMod(n, k int) int {
	if n < 0 || k < 0 || k > n {
		return 0
	}
	return (fact[n] * invFact[k] % MOD * invFact[n-k] % MOD) % MOD
}

func powMod(x, e int) int {
	res := 1
	for e > 0 {
		if e%2 == 1 {
			res = res * x % MOD
		}
		x = x * x % MOD
		e /= 2
	}
	return res
}
