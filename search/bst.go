package search

type Node struct {
	key string
	val interface{}
	h int
	l *Node
	r *Node
}

type BST struct {
	root *Node
	size int
}

func NewBST() *BST {
	return &BST{}
}

func (t *BST)Put(key string, val interface{})  {
	t.root = t.put(t.root, &Node{key:key, val:val})
}

func (t *BST)put(node, newNode *Node) *Node {
	if node == nil {
		t.size++
		return newNode
	}

	switch  {
	case newNode.key < node.key:
		node.l = t.put(node.l, newNode)
	case newNode.key > node.key:
		node.r = t.put(node.r, newNode)
	default:
		node.val = newNode.val
	}

	return node
}

func (t *BST) PreOrderKeys() []string{
	result := make([]string,0, t.size)
	t.preOrderKeys(t.root, result)
	return t.preOrderKeys(t.root, result)
}

func (t *BST)preOrderKeys(node *Node, result []string) []string {
	if node == nil {
		return result
	}
	result = t.preOrderKeys(node.l, result)
	result = append(result, node.key)
	result = t.preOrderKeys(node.r, result)
	return result
}

func (t *BST)Remove(key string) (bool, interface{}) {
	var res *Node
	t.root = t.remove(t.root, key, &res)
	if res != nil {
		return true, res.val
	}
	return false, nil
}

func (t *BST) remove(node *Node, key string, res **Node) *Node {
	if node == nil {
		return nil
	}

	switch  {
	case key < node.key:
		node.l =  t.remove(node.l, key, res)
	case key > node.key:
		node.r =  t.remove(node.r, key, res)
	default:
		*res = node
		newNode := node.l
		if newNode == nil {
			return node.r
		}
		lr := newNode.r

		newNode.r = node.r
		t.put(newNode, lr)
		return newNode
	}
	return node
}

func (t *BST)Get(key string) (bool, interface{}) {
	if exist, node := t.get(t.root, key); exist {
		return true, node.val
	}
	return false, nil
}

func (t *BST) get(node *Node, key string) (bool, *Node) {
	if node == nil {
		return false, nil
	}

	switch  {
	case key < node.key:
		return t.get(node.l, key)
	case key > node.key:
		return t.get(node.r, key)
	default:
		return true, node
	}
}