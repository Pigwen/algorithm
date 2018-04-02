package bst

type BinarySearchTree struct {
	root *Node
}

type Node struct {
	value int
	nodes [2]*Node
}

func (tree *BinarySearchTree) Insert(v int) bool {
	node := newNode(v)

	if tree.root == nil {
		tree.root = node
		return true
	} else {
		return insert(tree.root, v)
	}
}

func newNode(v int) *Node {
	return &Node{
		value: v,
	}
}

func insert(root *Node, v int) bool {
	current := root
	for {
		index := 0
		if current.value < v {
			index = 1
		}

		if current.value == v {
			return false
		} else if current.nodes[index] == nil {
			current.nodes[index] = newNode(v)
			break
		} else {
			current = current.nodes[index]
		}
	}
	return true
}

func (tree *BinarySearchTree) Contains(v int) bool {
	return contains(tree.root, v)
}

func contains(root *Node, v int) bool {
	current := root
	for {
		if current == nil {
			return false
		} else if current.value == v {
			return true
		} else {
			index := 0
			if current.value < v {
				index = 1
			}
			current = current.nodes[index]
		}
	}
}
