package search

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	skipList := New()
	skipList.Put("1", 1)
	skipList.Put("2", 2)
	skipList.Put("3", 3)
	skipList.Put("4", 4)
	skipList.Put("5", 6)

	skipList.Put("2", 8)
	//skipList.Put("2", 5)
	all := skipList.Values()
	fmt.Println(all)
	fmt.Println(skipList.ReverseValues())
	fmt.Println(skipList.Keys())
	fmt.Println(skipList.Get("2"))
	fmt.Println(skipList.Remove("2").value)
	fmt.Println(skipList.Values())

	fmt.Println("range", skipList.Range("3", "6"))
}

func TestSkipList_Get(t *testing.T) {
	var a interface{}
	a = 1
	var b interface{}
	b = 1

	fmt.Println(a != b)
	//fmt.Println(a >= b)
}
