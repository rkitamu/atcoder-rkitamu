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
	t := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		if t[i] {
			continue
		}
		if isPrime(i) {
			for j := i * 2; j <= n; j += i {
				t[j] = true
			}
		}
	}

	t[0] = true
	t[1] = true
	for i := 0; i <= n; i++ {
		if !t[i] {
			out(i)
		}
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

// nis2d reads n * m integers from stdin with offset support.
func nis2d(h, w, offset int) [][]int {
	a := make([][]int, h+offset)
	// offset
	for i := 0; i < offset; i++ {
		a[i] = make([]int, w+offset)
	}
	// content
	for i := offset; i < h+offset; i++ {
		a[i] = nis(w, offset)
	}
	return a
}

func nl() int64 {
	sc.Scan()
	i, e := strconv.ParseInt(sc.Text(), 10, 64)
	if e != nil {
		panic(e)
	}
	return i
}

var bufBytes []byte
var bufIdx int

func nr() byte {
	for {
		if bufIdx < len(bufBytes) {
			r := bufBytes[bufIdx]
			bufIdx++
			return r
		}

		if !sc.Scan() {
			panic("failed to scan next token")
		}
		bufBytes = []byte(sc.Text())
		bufIdx = 0
	}
}

// nr reads a single string from stdin.
func ns() string {
	if !sc.Scan() {
		panic("failed to scan next token")
	}
	return sc.Text()
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

// nr reads n bytes from stdin.
func nrs(n int) []byte {
	a := make([]byte, n)
	for i := 0; i < n; i++ {
		a[i] = nr()
	}
	return a
}

// nrs2d reads n * m bytes from stdin.
func nrs2d(n, m, offset int) [][]byte {
	a := make([][]byte, n+offset)
	for i := offset; i < n+offset; i++ {
		tmp := nrs(m)
		prepended := make([]byte, offset+len(tmp))
		copy(prepended[offset:], tmp)
		a[i] = prepended
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

// 相対誤差が10^-6: formatFloat(f, 7)
func formatFloat(f float64, precision int) string {
	if math.IsNaN(f) || math.IsInf(f, 0) {
		return fmt.Sprintf("%v", f)
	}

	magnitude := math.Log10(math.Abs(f))
	effectivePrecision := precision - int(magnitude) - 1

	if effectivePrecision < 0 {
		effectivePrecision = 0
	}
	if effectivePrecision > 15 { // float64の限界
		effectivePrecision = 15
	}

	return fmt.Sprintf("%.*f", effectivePrecision, f)
}

// out writes the output to stdout.
func out(v ...interface{}) {
	_, e := fmt.Fprintln(wtr, v...)
	if e != nil {
		panic(e)
	}
}

// out writes the output to stdout without a newLine.
func outNoLn(v ...interface{}) {
	_, e := fmt.Fprint(wtr, v...)
	if e != nil {
		panic(e)
	}
}

func out1dNumber[T Number](v []T) {
	for i, val := range v {
		if i > 0 {
			fmt.Fprint(wtr, " ")
		}
		fmt.Fprint(wtr, val)
	}
	fmt.Fprintln(wtr)
}

func out2dNumber[T Number](v [][]T) {
	for _, i := range v {
		out1dNumber(i)
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

// ======================
// data structure
// ======================
// BIT is a Binary Indexed Tree (Fenwick Tree) implementation
type BIT struct {
	n   int
	bit []int
}

// NewBIT initializes a new Binary Indexed Tree with size n
func NewBIT(n int) *BIT {
	return &BIT{n: n + 2, bit: make([]int, n+3)}
}

// Add x to index i (O(log n))
func (b *BIT) Add(i, x int) {
	i++
	for i < len(b.bit) {
		b.bit[i] += x
		i += i & -i
	}
}

// Sum returns the [0, i] (inclusive) sum (O(log n))
func (b *BIT) Sum(i int) int {
	i++
	res := 0
	for i > 0 {
		res += b.bit[i]
		i -= i & -i
	}
	return res
}

// RangeSum returns the sum of the range [l, r] (inclusive) (O(log n))
func (b *BIT) RangeSum(l, r int) int {
	if l > r {
		return 0
	}
	if l == 0 {
		return b.Sum(r)
	}
	return b.Sum(r) - b.Sum(l-1)
}

// SegmentTree (WIP(Implemented: push, add, get, sum))
type SegmentTree struct {
	n    int
	data []int
	lazy []int
}

func NewSegmentTree(n int) *SegmentTree {
	size := 1
	for size < n {
		size <<= 1
	}
	return &SegmentTree{
		n:    size,
		data: make([]int, 2*size),
		lazy: make([]int, 2*size),
	}
}

func (st *SegmentTree) push(k, l, r int) {
	if st.lazy[k] != 0 {
		st.data[k] += (r - l) * st.lazy[k]
		if r-l > 1 {
			st.lazy[2*k+1] += st.lazy[k]
			st.lazy[2*k+2] += st.lazy[k]
		}
		st.lazy[k] = 0
	}
}

// Add x to range [a, b)
func (st *SegmentTree) Add(a, b, x int) {
	var f func(k, l, r int)
	f = func(k, l, r int) {
		st.push(k, l, r)
		if r <= a || b <= l {
			return
		}
		if a <= l && r <= b {
			st.lazy[k] += x
			st.push(k, l, r)
		} else {
			mid := (l + r) / 2
			f(2*k+1, l, mid)
			f(2*k+2, mid, r)
			st.data[k] = st.data[2*k+1] + st.data[2*k+2]
		}
	}
	f(0, 0, st.n)
}

// Get value at index i
func (st *SegmentTree) Get(i int) int {
	k := 0
	l, r := 0, st.n
	for r-l > 1 {
		st.push(k, l, r)
		mid := (l + r) / 2
		if i < mid {
			k = 2*k + 1
			r = mid
		} else {
			k = 2*k + 2
			l = mid
		}
	}
	st.push(k, l, r)
	return st.data[k]
}

// Get sum of range [a, b)
func (st *SegmentTree) Sum(a, b int) int {
	var f func(k, l, r int) int
	f = func(k, l, r int) int {
		st.push(k, l, r)
		if r <= a || b <= l {
			return 0
		}
		if a <= l && r <= b {
			return st.data[k]
		}
		mid := (l + r) / 2
		return f(2*k+1, l, mid) + f(2*k+2, mid, r)
	}
	return f(0, 0, st.n)
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

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{data: make([]T, 0), head: 0, tail: 0}
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
//
//	import "container/heap"
//	h := &ItemHeap{}
//	heap.Init(h)
//	heap.Push(h, &Item{value: tc.tcase[i]})
//	heap.Pop(h).(*Item)
type Item struct {
	value int
}
type ItemHeap []*Item

func (h ItemHeap) Len() int { return len(h) }

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

type Vector struct {
	X, Y int
}

func NewVector(x, y int) *Vector {
	return &Vector{X: x, Y: y}
}

func NewVectorFromPointsSlice(start, end []int) *Vector {
	if len(start) < 2 || len(end) < 2 {
		panic("require at least 2 elements (x, y)")
	}
	return &Vector{
		X: end[0] - start[0],
		Y: end[1] - start[1],
	}
}

func (v Vector) Add(other Vector) *Vector {
	return &Vector{
		X: v.X + other.X,
		Y: v.Y + other.Y,
	}
}

func (v Vector) Magnitude() float64 {
	return math.Sqrt(float64(v.X*v.X + v.Y*v.Y))
}

func (v Vector) Dot(other Vector) float64 {
	return float64(v.X*other.X + v.Y*other.Y)
}

func (v Vector) Cross(other Vector) int {
	return v.X*other.Y - v.Y*other.X
}

func (v Vector) CrossMagnitude(other Vector) float64 {
	return math.Abs(float64(v.Cross(other)))
}

// Edge represents a connection from one node to another with an optional weight.
type Edge struct {
	To     int
	Weight int
}

// Node represents a node in the graph with optional value and its outgoing edges.
type Node struct {
	ID    int
	Value int
	Edges []Edge
}

// Graph represents a generic directed or undirected graph.
type Graph struct {
	Nodes map[int]*Node
}

// NewGraph initializes an empty graph.
func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[int]*Node),
	}
}

// AddNode adds a new node with the given ID and optional value.
func (g *Graph) AddNode(id int, value int) {
	g.Nodes[id] = &Node{
		ID:    id,
		Value: value,
		Edges: []Edge{},
	}
}

// AddEdge adds a directed edge from u to v with given weight.
func (g *Graph) AddEdge(u, v, weight int) {
	if _, ok := g.Nodes[u]; !ok {
		g.AddNode(u, 0)
	}
	if _, ok := g.Nodes[v]; !ok {
		g.AddNode(v, 0)
	}
	g.Nodes[u].Edges = append(g.Nodes[u].Edges, Edge{To: v, Weight: weight})
}

func (g *Graph) AddUndirectedEdge(u, v, weight int) {
	g.AddEdge(u, v, weight)
	g.AddEdge(v, u, weight)
}

// GetNode returns the node with the given ID.
func (g *Graph) GetNode(id int) *Node {
	return g.Nodes[id]
}

// =====================
// Math utils
// ======================
// isPrime checks if n is prime
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
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
func gcd[T Integer](a, b T) T {
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
	if 1 <= a {
		return a
	}
	return b
}

// gcds calculates the greatest common divisor of a slice of integers
func gcds[T Integer](a ...T) T {
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
func lcm[T Integer](a, b T) T {
	if a < b {
		a, b = b, a
	}
	return a / gcd(a, b) * b
}

func lcms[T Integer](a ...T) T {
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

func factorial(n int) int64 {
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

func pow[T Integer](base, exp T) T {
	if exp == 0 {
		return 1
	}
	if exp == 1 {
		return base
	}

	result := T(1)
	for i := T(0); i < exp; i++ {
		result *= base
	}
	return result
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

type Matrix[T Number] struct {
	Rows int
	Cols int
	Data [][]T
}

// NewMatrix creates a new matrix of given size initialized to zero
func NewMatrix[T Number](rows, cols int) *Matrix[T] {
	data := make([][]T, rows)
	for i := range data {
		data[i] = make([]T, cols)
	}
	return &Matrix[T]{Rows: rows, Cols: cols, Data: data}
}

// IdentityMatrix creates an identity matrix of size n
func IdentityMatrix[T Number](n int) *Matrix[T] {
	one := T(1)
	m := NewMatrix[T](n, n)
	for i := 0; i < n; i++ {
		m.Data[i][i] = one
	}
	return m
}

func (m *Matrix[T]) Add(b *Matrix[T]) *Matrix[T] {
	res := NewMatrix[T](m.Rows, m.Cols)
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			res.Data[i][j] = m.Data[i][j] + b.Data[i][j]
		}
	}
	return res
}

func (m *Matrix[T]) Sub(b *Matrix[T]) *Matrix[T] {
	res := NewMatrix[T](m.Rows, m.Cols)
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			res.Data[i][j] = m.Data[i][j] - b.Data[i][j]
		}
	}
	return res
}

func (m *Matrix[T]) Mul(b *Matrix[T]) *Matrix[T] {
	if m.Cols != b.Rows {
		panic("matrix dimensions do not match for multiplication")
	}
	res := NewMatrix[T](m.Rows, b.Cols)
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < b.Cols; j++ {
			var sum T
			for k := 0; k < m.Cols; k++ {
				sum += m.Data[i][k] * b.Data[k][j]
			}
			res.Data[i][j] = sum
		}
	}
	return res
}

func (m *Matrix[T]) Pow(n int64) *Matrix[T] {
	if m.Rows != m.Cols {
		panic("matrix must be square for exponentiation")
	}
	res := IdentityMatrix[T](m.Rows)
	base := m.Copy()
	for n > 0 {
		if n%2 == 1 {
			res = res.Mul(base)
		}
		base = base.Mul(base)
		n /= 2
	}
	return res
}

// MulMod performs matrix multiplication with modular arithmetic
func (m *Matrix[T]) MulMod(b *Matrix[T]) *Matrix[T] {
	if m.Cols != b.Rows {
		panic("matrix dimensions do not match for multiplication")
	}
	res := NewMatrix[T](m.Rows, b.Cols)
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < b.Cols; j++ {
			sum := 0
			for k := 0; k < m.Cols; k++ {
				product := mulMod(int(m.Data[i][k]), int(b.Data[k][j]))
				sum = addMod(sum, product)
			}
			res.Data[i][j] = T(sum)
		}
	}
	return res
}

// PowMod performs matrix exponentiation with modular arithmetic
func (m *Matrix[T]) PowMod(n int64) *Matrix[T] {
	if m.Rows != m.Cols {
		panic("matrix must be square for exponentiation")
	}
	res := IdentityMatrix[T](m.Rows)
	base := m.Copy()

	for n > 0 {
		if n%2 == 1 {
			res = res.MulMod(base)
		}
		base = base.MulMod(base)
		n /= 2
	}
	return res
}

func (m *Matrix[T]) Copy() *Matrix[T] {
	cpy := NewMatrix[T](m.Rows, m.Cols)
	for i := range m.Data {
		copy(cpy.Data[i], m.Data[i])
	}
	return cpy
}

func (m *Matrix[T]) Print() {
	for _, row := range m.Data {
		fmt.Println(row)
	}
}

// toBase converts an integer n to a base representation as a byte slice.
func toBase(n, base int) []byte {
	if n == 0 {
		return []byte{'0'}
	}

	digits := make([]byte, 0, 32) // 十分な長さを確保
	for n > 0 {
		r := n % base
		var b byte
		if r < 10 {
			b = '0' + byte(r)
		} else {
			b = 'a' + byte(r-10)
		}
		digits = append(digits, b)
		n /= base
	}

	// 反転（下位→上位）
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}
	return digits
}

// =====================
// Programming utils
// ======================
// constraints.Ordered
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

// Pair is like C++'s std::pair
type Pair[T, U Ordered] struct {
	First  T
	Second U
}

func NewPair[T, U Ordered](first T, second U) Pair[T, U] {
	return Pair[T, U]{First: first, Second: second}
}

func (p Pair[T, U]) Equals(other Pair[T, U]) bool {
	return p.First == other.First && p.Second == other.Second
}

func (p Pair[T, U]) Cmp(other Pair[T, U]) int {
	if p.First < other.First {
		return -1
	}
	if p.First > other.First {
		return 1
	}
	if p.Second < other.Second {
		return -1
	}
	if p.Second > other.Second {
		return 1
	}
	return 0
}

func (p Pair[T, U]) Lt(other Pair[T, U]) bool  { return p.Cmp(other) < 0 }
func (p Pair[T, U]) Lte(other Pair[T, U]) bool { return p.Cmp(other) <= 0 }
func (p Pair[T, U]) Gt(other Pair[T, U]) bool  { return p.Cmp(other) > 0 }
func (p Pair[T, U]) Gte(other Pair[T, U]) bool { return p.Cmp(other) >= 0 }

func (p Pair[T, U]) Max(other Pair[T, U]) Pair[T, U] {
	if p.Lt(other) {
		return other
	}
	return p
}

func (p Pair[T, U]) Min(other Pair[T, U]) Pair[T, U] {
	if p.Lt(other) {
		return p
	}
	return other
}

// Swap swaps the values of two variables.
func Swap[T any](a, b *T) {
	*a, *b = *b, *a
}

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type PracticalInteger interface {
	~int | ~int32 | ~int64 | ~uint | ~uint32 | ~uint64
}

// MOD
// 基本的なmod演算
func mod(a int, m ...int) int {
	modVal := MOD
	if len(m) > 0 {
		modVal = m[0]
	}
	return ((a % modVal) + modVal) % modVal // 負の数にも対応
}

// mod加算
func addMod[T PracticalInteger](a, b T, m ...int) T {
	modVal := T(MOD)
	if len(m) > 0 {
		modVal = T(m[0])
	}
	return (a + b) % modVal
}

// mod減算
func subMod(a, b int, m ...int) int {
	modVal := MOD
	if len(m) > 0 {
		modVal = m[0]
	}
	return ((a-b)%modVal + modVal) % modVal
}

// mod乗算
func mulMod(a, b int, m ...int) int {
	modVal := MOD
	if len(m) > 0 {
		modVal = m[0]
	}
	return (a * b) % modVal
}

// mod累乗 (a^n mod m)
func powMod(a, n int, m ...int) int {
	modVal := MOD
	if len(m) > 0 {
		modVal = m[0]
	}
	if n == 0 {
		return 1
	}
	result := 1
	a %= modVal
	for n > 0 {
		if n&1 == 1 {
			result = (result * a) % modVal
		}
		a = (a * a) % modVal
		n >>= 1
	}
	return result
}

// modの逆元 (a^(-1) mod m) - mが素数の場合
func invMod(a int, m ...int) int {
	modVal := MOD
	if len(m) > 0 {
		modVal = m[0]
	}
	return powMod(a, modVal-2, modVal) // フェルマーの小定理を利用
}

// mod除算 (a/b mod m) - mが素数の場合
func divMod(a, b int, m ...int) int {
	modVal := MOD
	if len(m) > 0 {
		modVal = m[0]
	}
	return mulMod(a, invMod(b, modVal), modVal)
}

// ビット数を数える
func popcount(x int) int {
	count := 0
	for x > 0 {
		count += x & 1
		x >>= 1
	}
	return count
}

// --------
// parindromes
// --------
func generatePalindrome(min, max int) *[]string {
	palindromes := make([]string, 0)

	// 範囲内の全ての回文を生成
	for length := len(strconv.Itoa(min)); length <= len(strconv.Itoa(max)); length++ {
		palindromes = append(palindromes, *generatePalindromesByLength(length)...)
	}

	// 範囲内の回文のみをフィルタリング
	// 遅そう
	result := make([]string, 0)
	for _, p := range palindromes {
		num, _ := strconv.Atoi(p)
		if num >= min && num <= max {
			result = append(result, p)
		}
	}

	return &result
}

func generatePalindromesByLength(length int) *[]string {
	palindromes := make([]string, 0, 10000) // 予測できるなら capacity を指定

	if length == 1 {
		for i := byte('0'); i <= '9'; i++ {
			palindromes = append(palindromes, string([]byte{i}))
		}
		return &palindromes
	}

	halfLength := (length + 1) / 2
	start := int(math.Pow10(halfLength - 1))
	end := int(math.Pow10(halfLength)) - 1

	for i := start; i <= end; i++ {
		left := strconv.Itoa(i)
		leftBytes := []byte(left)

		var palindrome []byte
		if length%2 == 0 {
			palindrome = make([]byte, 0, 2*len(leftBytes))
			palindrome = append(palindrome, leftBytes...)
			for j := len(leftBytes) - 1; j >= 0; j-- {
				palindrome = append(palindrome, leftBytes[j])
			}
		} else {
			palindrome = make([]byte, 0, 2*len(leftBytes)-1)
			palindrome = append(palindrome, leftBytes...)
			for j := len(leftBytes) - 2; j >= 0; j-- {
				palindrome = append(palindrome, leftBytes[j])
			}
		}
		palindromes = append(palindromes, string(palindrome))
	}
	return &palindromes
}

// 回文判定
func isPalindrome(bs []byte) bool {
	for i, j := 0, len(bs)-1; i < j; i, j = i+1, j-1 {
		if bs[i] != bs[j] {
			return false
		}
	}
	return true
}

// -------------
// 文字列操作
// -------------
// ReverseString reverses a string
func reverseString(input *[]byte) *[]byte {
	src := *input
	n := len(src)
	dst := make([]byte, n)
	for i := 0; i < n; i++ {
		dst[i] = src[n-1-i]
	}
	return &dst
}
