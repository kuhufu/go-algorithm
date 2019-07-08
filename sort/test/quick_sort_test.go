package test

import (
	"fmt"
	"kuhufu/algorithm/sort"
	"kuhufu/algorithm/util"
	"testing"
)

func TestQuickSort(t *testing.T) {
	a := util.RandIntArray(TestNum)
	sort.QuickSort(a)
	if !util.IsSorted(a) {
		t.Fatal("error, array is not sorted")
	}

	fmt.Println(a)
}

func TestQuick3Way(t *testing.T) {
	a := util.RandIntArray(TestNum)
	sort.Quick3Way(a)
	if !util.IsSorted(a) {
		t.Fatal("error, array is not sorted")
	}

	fmt.Println(a)
}

func BenchmarkQuickSort(b *testing.B) {
	a := util.RandIntArray(BenchNum)
	b.ReportAllocs()
	b.ResetTimer()
	for i :=0 ; i < b.N; i++ {
		sort.QuickSort(a)
	}
}

func BenchmarkQuick3Way(b *testing.B) {
	a := util.RandIntArray(BenchNum)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.Quick3Way(a)
	}
}
