package search

type AVLNode struct {
	key string
	val interface{}
	h int
	l *Node
	r *Node
}

type AVL struct {
	root *Node
	size int
}

func NewAVL() *AVL {
	return &AVL{}
}

func (t *AVL)Put(key string, val interface{})  {
	t.root = t.put(t.root, &Node{key:key, val:val, h:1})
}

func (t *AVL)put(node, newNode *Node) *Node {
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

	return checkAndRotate(node)
}

func (t *AVL) PreOrderKeys() []string{
	result := make([]string,0, t.size)
	t.preOrderKeys(t.root, result)
	return t.preOrderKeys(t.root, result)
}

func (t *AVL)preOrderKeys(node *Node, result []string) []string {
	if node == nil {
		return result
	}
	result = t.preOrderKeys(node.l, result)
	result = append(result, node.key)
	result = t.preOrderKeys(node.r, result)
	return result
}

func (t *AVL)Remove(key string) (bool, interface{}) {
	var res *Node
	t.root = t.remove(t.root, key, &res)
	if res != nil {
		return true, res.val
	}
	return false, nil
}

func (t *AVL) remove(node *Node, key string, res **Node) *Node {
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
			return   node.r
		}
		lr := newNode.r

		newNode.r = node.r
		t.put(newNode, lr)
		return newNode
	}
	return checkAndRotate(node)
}

func checkAndRotate(node *Node) *Node {
	hl := height(node.l)
	hr := height(node.r)

	d := hl - hr
	if d > 1 {
		node = RotateToRight(node)
	} else if d < -1 {
		node = RotateToLeft(node)
	}

	hl = height(node.l)
	hr = height(node.r)
	node.h = 1 + max(hl, hr)
	return node
}

func (t *AVL)Get(key string) (bool, interface{}) {
	if exist, node := t.get(t.root, key); exist {
		return true, node.val
	}
	return false, nil
}

func (t *AVL) get(node *Node, key string) (bool, *Node) {
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

func RotateToLeft(node *Node) *Node {
	rNode := node.r
	if height(rNode.l) > height(rNode.r) {
		rNode.l.h = rNode.h
		rNode.h--
		node.r = rotateToRight(rNode)
	}

	res := rotateToLeft(node)
	node.h--
	res.h = node.h
	return res
}

func rotateToLeft(node *Node) *Node {
	right := node.r
	node.r = right.l
	right.l = node
	return right
}

func RotateToRight(node *Node) *Node {
	lNode := node.l
	if height(lNode.r) > height(lNode.l) {
		lNode.r.h = lNode.h
		lNode.h--
		node.l = rotateToLeft(lNode)
	}

	res := rotateToRight(node)
	node.h--
	res.h = node.h
	return res
}

func rotateToRight(node *Node) *Node {
	left := node.l
	node.l = left.r
	left.r = node
	return  left
}