package main

import "testing"

var codes []int

func init() {
	codes = getInput("input")

}

func BenchmarkfindProductNaive(b *testing.B) {
	for n := 0; n < b.N; n++ {
		findProductNaive(codes)
	}
}

func BenchmarkfindProductRecurse(b *testing.B) {
	for n := 0; n < b.N; n++ {
		findProductRecurse(codes)
	}
}
