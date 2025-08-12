package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
	// if use Combination
	// initCombTable()
}

const MAX_COMB_CACHE = 10000000
const MOD = 1000000007

func main() {
	defer flush()

	nm := nis(2)
	n, m := nm[0], nm[1]

	type Segment struct {
		l, r int
	}
	seg := make([]Segment, m)
	for i := 0; i < m; i++ {
		lr := nis(2)
		seg[i] = Segment{l: lr[0], r: lr[1]}
	}
	sort.Slice(seg, func(i, j int) bool {
		if seg[i].l == seg[j].l {
			return seg[i].r < seg[j].r
		}
		return seg[i].l < seg[j].l
	})

	q := ni()
	query := make([]*Segment, q)
	for i := range query {
		lr := nis(2)
		query[i] = &Segment{l: lr[0], r: lr[1]}
	}

	// 線分Mを辺、別れた領域をノードとする木を作る
	tree := make([][]int, m+1)
	type Frame struct {
		r      int
		region int
	}
	stack := NewStack[Frame](2*n + 1)
	stack.Push(Frame{r: 2*n + 1, region: 0}) // 番兵
	region := 1
	for i := 0; i < m; i++ {
		_, r := seg[i].l, seg[i].r
		for stack.Top().r < r {
			stack.Pop()
		}
		par := stack.Top().region
		tree[par] = append(tree[par], region)
		tree[region] = append(tree[region], par)
		stack.Push(Frame{r: r, region: region})
		region++
	}

	// 円周上の奇数の点がどの領域に属するかを求める
	pointRegion := make([]int, 2*n+2)
	st := NewStack[Frame](2*n + 2)
	st.Push(Frame{r: 2*n + 2, region: 0})
	cur := 0
	for i := 2; i <= 2*n+1; i++ {
		for cur < m && seg[cur].l == i {
			st.Push(Frame{r: seg[cur].r, region: cur + 1})
			cur++
		}
		for st.Top().r <= i {
			st.Pop()
		}
		if i%2 == 1 {
			pointRegion[i] = st.Top().region
		}
	}

	// LCA 用の前処理
	const LOG = 21
	up := make([][]int, m+1)
	for i := range up {
		up[i] = make([]int, LOG)
	}
	depth := make([]int, m+1)
	var dfs func(v, p int)
	dfs = func(v, p int) {
		up[v][0] = p
		for i := 1; i < LOG; i++ {
			up[v][i] = up[up[v][i-1]][i-1]
		}
		for _, u := range tree[v] {
			if u != p {
				depth[u] = depth[v] + 1
				dfs(u, v)
			}
		}
	}
	dfs(0, 0)

	lca := func(u, v int) int {
		if depth[u] < depth[v] {
			u, v = v, u
		}
		for i := LOG - 1; i >= 0; i-- {
			if depth[u]-(1<<i) >= depth[v] {
				u = up[u][i]
			}
		}
		if u == v {
			return u
		}
		for i := LOG - 1; i >= 0; i-- {
			if up[u][i] != up[v][i] {
				u = up[u][i]
				v = up[v][i]
			}
		}
		return up[u][0]
	}

	for i := 0; i < q; i++ {
		c, d := query[i].l, query[i].r
		u, v := pointRegion[c], pointRegion[d]
		ancestor := lca(u, v)
		out(depth[u] + depth[v] - 2*depth[ancestor])
	}
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

// ======================
// data structure
// ======================
// BIT is a Binary Indexed Tree (Fenwick Tree) implementation
type BIT struct {
	n   int
	bit []int
}

func NewBIT(n int) *BIT {
	return &BIT{n: n + 2, bit: make([]int, n+3)}
}

func (b *BIT) Add(i, x int) {
	i++
	for i < len(b.bit) {
		b.bit[i] += x
		i += i & -i
	}
}

func (b *BIT) Sum(i int) int {
	i++
	res := 0
	for i > 0 {
		res += b.bit[i]
		i -= i & -i
	}
	return res
}

// Stack is a simple stack implementation
type Stack[T any] struct {
	data []T
}

func NewStack[T any](size int) *Stack[T] {
	return &Stack[T]{data: make([]T, size)}
}

func (s *Stack[T]) Push(v T) {
	s.data = append(s.data, v)
}

func (s *Stack[T]) Pop() T {
	last := len(s.data) - 1
	v := s.data[last]
	s.data = s.data[:last]
	return v
}

func (s *Stack[T]) Empty() bool {
	return len(s.data) == 0
}

func (s *Stack[T]) Len() int {
	return len(s.data)
}

func (s *Stack[T]) Top() T {
	return s.data[len(s.data)-1]
}

// Queue is a simple queue implementation
type Queue[T any] struct {
	data []T
	head int
	tail int
}

func NewQueue[T any](size int) *Queue[T] {
	return &Queue[T]{data: make([]T, size), head: 0, tail: 0}
}
func (q *Queue[T]) Enqueue(v T) {
	q.data = append(q.data, v)
	q.tail++
}
func (q *Queue[T]) Dequeue() T {
	if q.head == q.tail {
		panic("queue is empty")
	}
	v := q.data[q.head]
	q.head++
	if q.head == len(q.data)/2 {
		q.data = q.data[q.head:]
		q.tail -= q.head
		q.head = 0
	}
	return v
}
func (q *Queue[T]) Empty() bool {
	return q.head == q.tail
}
func (q *Queue[T]) Len() int {
	return q.tail - q.head
}
func (q *Queue[T]) Top() T {
	if q.head == q.tail {
		panic("queue is empty")
	}
	return q.data[q.head]
}

// =====================
// Math utils
// ======================
// isPrime checks if n is prime
func isPrime(n int) bool {
	if n < 2 { return false }
	if n == 2 { return true }
	cur := 3
	max := int(math.Floor(float64(math.Sqrt(float64(n)))))
	for cur <= max {
		if m := n % cur; m == 0 {
			return false
		}
		cur++
	}
	return true
}

// gcd calculates the greatest common divisor of a and b
func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}
	for 1 <= a && 1 <= b {
		mod := a % b
		if mod == 0 {
			return b
		}
		a, b = b, mod
	}
	if 1 <= a { return a}
	return b
}

// gcds calculates the greatest common divisor of a slice of integers
func gcds(a ...int) int {
	if len(a) < 2 {
		panic("gcds: at least 2 arguments required")
	}
	g := a[0]
	for i := 1; i < len(a); i++ {
		g = gcd(g, a[i])
	}
	return g
}
