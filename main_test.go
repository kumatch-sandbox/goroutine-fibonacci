package main

import (
	"testing"
)

func BenchmarkFibonacciA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacciA(30)
	}
}

func BenchmarkFibonacciB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacciB(30)
	}
}
