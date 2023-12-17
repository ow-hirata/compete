package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

func Solve(r *Scanner, w *Writer) {
	S := r.String()
	ans := math.MaxInt32
	for i := 0; i < len(S)-2; i++ {
		ChooseMin(&ans, Abs(Atoi(S[i:i+3])-753))
	}
	w.Println(ans)
}

func main() {
	r, w := NewScanner(os.Stdin, maxBufferSize), NewWriter(os.Stdout)
	defer w.Flush()
	Solve(r, w)
}

type signed interface{ ~int | ~int32 | ~int64 }

type unsigned interface{ ~uint | ~uint32 | ~uint64 }

type integer interface{ signed | unsigned }

type float interface{ ~float32 | ~float64 }

type actual interface{ integer | float }

type ordered interface{ integer | float | ~string }

func Unwrap[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}

func Atoi(s string) int { return Unwrap(strconv.Atoi(s)) }

func ChooseMin[T ordered](a *T, b T) T {
	if *a > b {
		*a = b
	}
	return *a
}

func Abs[T actual](x T) T {
	if x < T(0) {
		return -x
	}
	return x
}

const maxBufferSize = 1 * 1024 * 1024

type Scanner struct{ sc *bufio.Scanner }

func NewScanner(r io.Reader, size int) *Scanner {
	sc := bufio.NewScanner(r)
	sc.Buffer(make([]byte, size), size)
	sc.Split(bufio.ScanWords)
	return &Scanner{sc}
}

func (s *Scanner) scan() bool { return s.sc.Scan() }

func (s *Scanner) text() string { return s.sc.Text() }

func (s *Scanner) String() string { s.scan(); return s.text() }

type Writer struct{ bf *bufio.Writer }

func NewWriter(w io.Writer) *Writer { return &Writer{bufio.NewWriter(w)} }

func (w *Writer) Println(a ...interface{}) { fmt.Fprintln(w.bf, a...) }

func (w *Writer) Flush() { w.bf.Flush() }

type UnionFind struct {
	parent []int
	size   []int
}
