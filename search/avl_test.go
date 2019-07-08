package search

import (
	"fmt"
	"testing"
)

var tree = NewAVL()

func init()  {
	tree.Put("2",1)
	tree.Put("1",1)
	tree.Put("3",1)
	tree.Put("4",1)
	tree.Put("5",5)
	tree.Put("6",5)
}

func TestAVL_PreOrderKeys(t *testing.T) {
	fmt.Println(tree.PreOrderKeys())
	fmt.Print(tree.size)
}

func TestAVL_Get(t *testing.T) {
	if exist, val:= tree.Get("5"); !exist || val != 5 {
		t.Error("error")
	}

	if exist, val:= tree.Get("6"); !exist || val != 5 {
		t.Error("error")
	}


}

func TestAVL_Remove(t *testing.T) {
	if exist, val := tree.Remove("5"); !exist || val != 5{
		t.Error("error")
	}
	if exist, val := tree.Remove("6"); !exist || val != 5{
		t.Error("error")
	}

	fmt.Println(tree.root.h)
}

