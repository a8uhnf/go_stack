package main

import "testing"

var a []int

func init() {
	for i := 0; i < 10000; i++ {
		a = append(a, i)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		MergeSort(a)
	}
}
func BenchmarkMergeSortMulti(b *testing.B) {
	for n := 0; n < b.N; n++ {
		MergeSortMulti(a)
	}
}

func BenchmarkMergeSortMultiEffi(b *testing.B) {
	for n := 0; n < b.N; n++ {
		MergeSortMultiEffi(a)
	}
}
