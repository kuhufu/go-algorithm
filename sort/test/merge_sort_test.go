package test

import (
	"fmt"
	"kuhufu/algorithm/sort"
	"kuhufu/algorithm/util"
	"testing"
)

func TestMergeSort(t *testing.T) {
	var data = util.RandIntArray(TestNum)
	sort.MergeSort(data)
	if !util.IsSorted(data) {
		t.Error("un sort")
		t.Log(data)
	}
}

func TestMergeSortFullCopy(t *testing.T) {
	var data = util.RandIntArrayUseNewSeed(TestNum)
	fmt.Println(data)
	sort.MergeSortFullCopy(data)
	if !util.IsSorted(data) {
		t.Error("un sort")
		t.Log(data)
	}
	fmt.Println(data)
}

func TestMergeSortHalfCopy(t *testing.T) {
	var data = util.RandIntArrayUseNewSeed(TestNum)
	sort.MergeSortHalfCopy(data)
	if !util.IsSorted(data) {
		t.Error("un sort")
		t.Log(data)
	}
	fmt.Println(data)
}
func TestMergeSortWithP(t *testing.T) {
	var data = util.RandIntArray(TestNum)
	sort.MergeSortWithP(data)
	if !util.IsSorted(data) {
		t.Error("un sort")
		t.Log(data)
	}
	fmt.Println(data)
}

func TestMergeSortTest(t *testing.T) {
	var data = util.RandIntArray(TestNum)
	sort.MergeSortFinal(data)
	if !util.IsSorted(data) {
		t.Error("un sort")
		t.Log(data)
	}
	fmt.Println(data)
}

func BenchmarkMergeSortHalfCopy(b *testing.B) {
	data := util.RandIntArray(BenchNum)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sort.MergeSortHalfCopy(data)
	}
}

func BenchmarkMergeSortFullCopy(b *testing.B) {
	data := util.RandIntArray(BenchNum)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sort.MergeSortFullCopy(data)
	}
}

func BenchmarkMergeSortWithP(b *testing.B) {
	data := util.RandIntArray(BenchNum)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.MergeSortWithP(data)
	}
}

func BenchmarkMergeSortTest(b *testing.B) {
	data := util.RandIntArray(BenchNum)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.MergeSortFinal(data)
	}
}