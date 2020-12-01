package main

import "testing"

func BenchmarkfindProductNaive(b *testing.B) {
	codes := getInput("input")
	findProductNaive(codes)

}
