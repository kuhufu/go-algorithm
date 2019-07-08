package search

func height(node *Node) int {
	if node == nil {
		return 0
	}
	return 1 + max(height(node.l), height(node.r))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
