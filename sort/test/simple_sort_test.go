package test

import (
	"kuhufu/algorithm/sort"
	"kuhufu/algorithm/util"
	"testing"
)

const (
	TestNum  = 10
	BenchNum = 16
)

func TestBubSort(t *testing.T) {
	var data = util.RandIntArray(TestNum)
	sort.BubSort(data)
	if !util.IsSorted(data) {
		t.Error("un sort")
	}
}

func TestSelectSort(t *testing.T) {
	var data = util.RandIntArray(TestNum)
	sort.SelectSort(data)
	if !util.IsSorted(data) {
		t.Error("un sort")
	}
}
func TestSelectSort2(t *testing.T) {
	var data = util.RandIntArray(TestNum)
	sort.SelectSort2(data)
	if !util.IsSorted(data) {
		t.Error("un sort")
	}
}

func TestSelectSort3(t *testing.T) {
	var data = util.RandIntArray(TestNum)
	sort.SelectSort3(data)
	if !util.IsSorted(data) {
		t.Error("un sort")
	}
}

func TestInsertSort(t *testing.T) {
	var data = util.RandIntArray(TestNum)
	sort.InsertSort(data)
	if !util.IsSorted(data) {
		t.Error("un sort")
	}
}

func BenchmarkBubSort(b *testing.B) {
	data := util.RandIntArray(BenchNum)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.BubSort(data)
	}
}

func BenchmarkInsertSort(b *testing.B) {
	data := util.RandIntArray(BenchNum)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.InsertSort(data)
	}
}

func BenchmarkSelectSort(b *testing.B) {
	data := util.RandIntArray(BenchNum)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.SelectSort(data)
	}
}

func BenchmarkSelectSort2(b *testing.B) {
	data := util.RandIntArray(BenchNum)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.SelectSort2(data)
	}
}

func BenchmarkSelectSort3(b *testing.B) {
	data := util.RandIntArray(BenchNum)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.SelectSort3(data)
	}
}
