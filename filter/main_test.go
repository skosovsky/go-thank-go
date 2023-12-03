package main

import "testing"

func BenchmarkDiv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isEvenDiv(10000000)
	}
}

func BenchmarkBit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isEvenBit(10000000)
	}
}
