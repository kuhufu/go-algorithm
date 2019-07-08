package search

import (
	"math/rand"
)

type SkipNode struct {
	levels []*SkipNode
	prev   *SkipNode
	Key    string
	value  interface{}
}

type SkipList struct {
	rootNode      *SkipNode
	maxLevelIndex int
	size          int
	lastNode      *SkipNode
}

func New() *SkipList {
	return &SkipList{
		rootNode: &SkipNode{
			levels: make([]*SkipNode, 32),
		},
	}
}

func (list *SkipList) Put(key string, value interface{}) {
	if node := list.get(key); node != nil {
		node.value = value
		return
	}
	nodeLevel := rand.Intn(32)
	newNode := &SkipNode{
		levels: make([]*SkipNode, nodeLevel+1),
		Key:    key,
		value:  value,
	}

	if nodeLevel > list.maxLevelIndex {
		list.maxLevelIndex = nodeLevel
	}

	// curNode为比newNode小的node
	curNode := list.rootNode
loop:
	for i := nodeLevel; i >= 0; {
		if curNode.levels[i] == nil {
			curNode.levels[i] = newNode
			if i == 0 {
				newNode.prev = curNode
				break loop
			}
			i--
			continue
		}
		for {
			nextNode := curNode.levels[i]
			if newNode.Key < nextNode.Key {
				curNode.levels[i] = newNode
				newNode.levels[i] = nextNode
				if i == 0 {
					nextNode.prev = newNode
					newNode.prev = curNode
					break loop
				}
				i--
				continue
			} else if (newNode.Key > nextNode.Key) {
				curNode = nextNode
				break
			} else {
				nextNode.value = value
				break loop
			}
		}
	}
	if newNode.levels[0] == nil {
		list.lastNode = newNode
	}
	list.size++
}

func (list *SkipList) Get(key string) interface{} {
	node := list.get(key)
	if node == nil {
		return nil
	}
	return node.value
}

func (list *SkipList) get(key string) *SkipNode {
	curNode := list.rootNode
	for i := list.maxLevelIndex; i >= 0; i-- {
		nextNode := curNode.levels[i]
		for nextNode != nil {
			if nextNode.Key == key {
				return nextNode
			} else if (key > nextNode.Key) {
				curNode = nextNode
				for i >= 0 && curNode.levels[i] == nil {
					i--
				}
				if i < 0 {
					return nil
				}
				nextNode = curNode.levels[i]
			} else if (key < nextNode.Key) {
				break
			}
		}

	}
	return nil
}

func (list *SkipList) Remove(key string) *SkipNode {
	curNode := list.rootNode
	var result *SkipNode
	for i := list.maxLevelIndex; i >= 0; i-- {
		nextNode := curNode.levels[i]
		for nextNode != nil {
			if nextNode.Key == key {
				result = nextNode
				for ; i >= 0; {
					if curNode.levels[i] == result {
						curNode.levels[i] = result.levels[i]
						result.levels[i] = nil
						i--
						continue
					}
					curNode = curNode.levels[i]
				}
				list.size--
				return nextNode
			} else if (key > nextNode.Key) {
				curNode = nextNode
				for i >= 0 && curNode.levels[i] == nil {
					i--
				}
				if i < 0 {
					return nil
				}
				nextNode = curNode.levels[i]
			} else if (key < nextNode.Key) {
				break
			}
		}

	}
	return result
}

func (list *SkipList) Contains(key string) bool {
	return list.get(key) != nil
}

func (list *SkipList) Keys() []string {
	result := make([]string, 0, list.size)
	curNode := list.rootNode.levels[0]
	for i := 0; i < list.size; i++ {
		result = append(result, curNode.Key)
		curNode = curNode.levels[0]
	}
	return result
}

func (list *SkipList) Values() []interface{} {
	result := make([]interface{}, 0, list.size)

	for curNode := list.rootNode.levels[0]; curNode != nil; curNode = curNode.levels[0] {
		result = append(result, curNode.value)
	}
	return result
}

func (list *SkipList) ReverseValues() []interface{} {
	result := make([]interface{}, 0, list.size)
	curNode := list.lastNode
	for i := 0; i < list.size; i++ {
		result = append(result, curNode.value)
		curNode = curNode.prev
	}
	return result
}

func (list *SkipList) Range(begin, end string) []interface{} {
	result := []interface{}{}
	curNode := list.get(begin)
	for curNode != nil && curNode.Key < end {
		result = append(result, curNode.value)
		curNode = curNode.levels[0]
	}
	return result
}
