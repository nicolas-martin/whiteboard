package main

import "testing"

func BenchmarkBigInt10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		createBig()
	}
}

func BenchmarkBigInt50(b *testing.B) {
	for n := 0; n < b.N; n++ {
		createBig()
	}
}

func BenchmarkCustomBigInt10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		createCustomBig()
	}
}

func BenchmarkCustomBigInt50(b *testing.B) {
	for n := 0; n < b.N; n++ {
		createCustomBig()
	}
}
