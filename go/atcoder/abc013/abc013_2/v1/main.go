package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func Solve(r *Scanner, w *Writer) {
	a, b := minmax(r.Int(), r.Int())
	w.Println(min(b-a, 10-(b-a)))
}

func main() {
	r, w := NewScanner(os.Stdin, MaxBufferSize), NewWriter(os.Stdout)
	defer w.Flush()
	Solve(r, w)
}

// signed is a constraint that permits any signed integer type.
type signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// unsigned is a constraint that permits any unsigned integer type.
type unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// integer is a constraint that permits any integer type.
type integer interface {
	signed | unsigned
}

// float is a constraint that permits any floating-point type.
type float interface {
	~float32 | ~float64
}

// ordered is a constraint that permits any ordered type: any type
type ordered interface {
	integer | float | ~string
}

const MaxBufferSize = 1 * 1024 * 1024

type Scanner struct{ sc *bufio.Scanner }

func NewScanner(r io.Reader, size int) *Scanner {
	sc := bufio.NewScanner(r)
	sc.Buffer(make([]byte, size), size)
	sc.Split(bufio.ScanWords)
	return &Scanner{sc}
}
func (s *Scanner) scan() bool       { return s.sc.Scan() }
func (s *Scanner) text() string     { return s.sc.Text() }
func (s *Scanner) String() string   { s.scan(); return s.text() }
func (s *Scanner) Int() int         { return atoi(s.String()) }
func (s *Scanner) Float64() float64 { return atof(s.String()) }

type Writer struct{ bf *bufio.Writer }

func NewWriter(w io.Writer) *Writer        { return &Writer{bufio.NewWriter(w)} }
func (w *Writer) Print(a ...interface{})   { fmt.Fprint(w.bf, a...) }
func (w *Writer) Println(a ...interface{}) { fmt.Fprintln(w.bf, a...) }
func (w *Writer) Flush()                   { w.bf.Flush() }

// unwrap returns the value of v if err is nil and panics otherwise.
func unwrap[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

// atoi returns an integer converted from s.
func atoi(s string) int { return unwrap(strconv.Atoi(s)) }

// atof returns a float converted from s.
func atof(s string) float64 { return unwrap(strconv.ParseFloat(s, 64)) }

// min returns the minimum value of a and b.
func min[T ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// minmax returns the minimum and maximum values of a and b.
func minmax[T ordered](a, b T) (T, T) {
	if a > b {
		return b, a
	}
	return a, b
}
