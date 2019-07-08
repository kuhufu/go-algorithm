package search

import (
	"testing"
	"fmt"
)

var bst = NewBST()

func init()  {
	bst.Put("1",1)
	bst.Put("2",1)
	bst.Put("5",5)
	bst.Put("4",1)
	bst.Put("3",1)
}

func TestBST_PreOrderKeys(t *testing.T) {

	fmt.Println(bst.PreOrderKeys())
	fmt.Print(bst.size)
}

func TestBST_Get(t *testing.T) {
	if exist, val:= bst.Get("5"); !exist || val != 5 {
		t.Error("error")
	}
}

func TestBST_Remove(t *testing.T) {
	if exist, val := bst.Remove("2"); !exist || val != 1{
		t.Error("error")
	}

	if exist, _:= bst.Get("2"); exist {
		t.Error("error")
	}

	bst.Put("2", 2)
	if exist, val := bst.Get("2"); !exist || val != 2 {
		t.Error("error")
	}
}
