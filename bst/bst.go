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

func (tree *BinarySearchTree) delete(v int) bool {
	if tree.root != nil {
		r := delete(tree.root, v)
		if r != nil {
			tree.root = r
			return true
		}
	}
	return false

}

func delete(root *Node, v int) *Node {
	dummy := &Node{
		value: 0,
	}
	dummy.nodes[1] = root

	current := dummy
	index := 1

	var r *Node
	var parent *Node

	for current.nodes[index] != nil {
		parent = current
		current = current.nodes[index]
		if current.value <= v {
			index = 1
		} else {
			index = 0
		}

		if current.value == v {
			r = current
		}
	}

	if r != nil {
		r.value = current.value

		dirp := 0
		if parent.nodes[1] == current {
			dirp = 1
		}

		dirc := 0
		if current.nodes[0] == nil {
			dirc = 1
		}
		parent.nodes[dirp] = current.nodes[dirc]
		return dummy.nodes[1]
	}
	return nil
}
