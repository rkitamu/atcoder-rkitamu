package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
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

	set := make(map[int]struct{})

	for i := 0; i < n; i++ {
		set[a[i]] = struct{}{}
	}

	ans := make([]string, 0)
	for i := 0; i <= 100; i++ {
		_, ok := set[i]
		if ok {
			ans = append(ans, fmt.Sprintf("%d", i))
		}
	}
	out(len(ans))
	out(strings.Join(ans, " "))
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

// nr reads a single rune from stdin.
func ns() []rune {
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

// ======================
// type utils
// ======================
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

// =====================
// calculation utils
// =====================
var fact, invFact []int
var factorialInitialized = false
// initFactorialTable initializes the factorial cache table
func initFactorialTable() {
	if factorialInitialized {
		return
	}
	factorialInitialized = true

	fact = make([]int, FACTORIAL_CACHE_SIZE+1)
	invFact = make([]int, FACTORIAL_CACHE_SIZE+1)
	fact[0] = 1
	for i := 1; i <= FACTORIAL_CACHE_SIZE; i++ {
		fact[i] = fact[i-1] * i % MOD
	}
	invFact[FACTORIAL_CACHE_SIZE] = powMod(fact[FACTORIAL_CACHE_SIZE], MOD-2)
	for i := FACTORIAL_CACHE_SIZE - 1; i >= 0; i-- {
		invFact[i] = invFact[i+1] * (i + 1) % MOD
	}
}

// combination calculates nCk
func combMod(n, k int) int {
	initFactorialTable()
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

// Priority Queue
// usage:
// 	import "container/heap"
// 	h := &ItemHeap{}
// 	heap.Init(h)
// 	heap.Push(h, &Item{value: tc.tcase[i]})
// 	heap.Pop(h).(*Item)
type Item struct {
	value int
}
type ItemHeap []*Item
func (h ItemHeap) Len() int            { return len(h) }
// min-heap implementation
func (h ItemHeap) Less(i, j int) bool  { return h[i].value < h[j].value }
func (h ItemHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *ItemHeap) Push(x interface{}) { *h = append(*h, x.(*Item)) }
func (h *ItemHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
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

// lcm calculates the least common multiple of a and b
func lcm(a, b int) int {
	if a < b {
		a, b = b, a
	}
	return a / gcd(a, b) * b
}

func lcms(a ...int) int {
	if len(a) < 2 {
		panic("lcms: at least 2 arguments required")
	}
	l := a[0]
	for i := 1; i < len(a); i++ {
		l = lcm(l, a[i])
	}
	return l
}

var factorialCache = make([]int64, 0)
func factorial(n  int) int64 {
	if n < 0 {
		panic("factorial: n must be non-negative")
	}
	if len(factorialCache) > n {
		return factorialCache[n]
	}
	for i := len(factorialCache); i <= n; i++ {
		if i == 0 {
			factorialCache = append(factorialCache, 1)
		} else {
			factorialCache = append(factorialCache, factorialCache[i-1]*int64(i))
		}
	}
	return factorialCache[n]
}

func comb(n, k int) int64 {
	if n < 0 || k < 0 || k > n {
		return 0
	}
	return factorial(n) / (factorial(k) * factorial(n-k))
}

func min[T Number](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func max[T Number](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func abs[T Number](a T) T {
	if a < 0 {
		return -a
	}
	return a
}

func fibonacci(n int) int {
	if n < 0 {
		panic("fibonacci: n must be non-negative")
	}
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}
