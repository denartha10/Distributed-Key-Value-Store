package hashing

import (
	"bytes"
)

type AVLTree struct {
	root *AVLNode
}

func (t *AVLTree) Insert(n *AVLNode) {
	t.root = t.root.insert(n)
}

func (t *AVLTree) Search(key []byte) *AVLNode {
	return t.root.search(key)
}

func (t *AVLTree) Delete(key []byte) {
	t.root = t.root.delete(key)
}

func (t *AVLTree) Do(do func(*AVLNode)) {
	if t.root != nil {
		t.root.do(do)
	}
}

type AVLNode struct {
	Key    []byte
	Val    interface{}
	left   *AVLNode
	right  *AVLNode
	height int // the height of nodes
}

func NewAVLNode(key []byte, v interface{}) *AVLNode {
	return &AVLNode{
		Key:    key,
		Val:    v,
		left:   nil,
		right:  nil,
		height: 0,
	}
}

// i.e., max(n.right.height, n.left.height) + 1
func (n *AVLNode) updateHeight() {
	left, right := -1, -1
	if n.left != nil {
		left = n.left.height
	}
	if n.right != nil {
		right = n.right.height
	}
	if left > right {
		n.height = left + 1
	} else {
		n.height = right + 1
	}
}

func (n *AVLNode) rightRotation() *AVLNode {
	n1 := n.left
	n.left = n1.right
	n1.right = n
	n.updateHeight()
	n1.updateHeight()
	return n1
}

func (n *AVLNode) leftRotation() *AVLNode {
	n1 := n.right
	n.right = n1.left
	n1.left = n
	n.updateHeight()
	n1.updateHeight()
	return n1
}

func (n *AVLNode) balanceFactor() int {
	// calculate balance factor
	left, right := -1, -1
	if n.left != nil {
		left = n.left.height
	}
	if n.right != nil {
		right = n.right.height
	}
	return left - right
}

// return the new root
func (n *AVLNode) rebalance() *AVLNode {
	bf := n.balanceFactor()
	if bf > 1 {
		// LR
		if n.left != nil && n.left.balanceFactor() < 0 {
			n.left = n.left.leftRotation()
		}
		// LL
		n = n.rightRotation()
	} else if bf < -1 {
		// RL
		if n.right != nil && n.right.balanceFactor() > 0 {
			n.right = n.right.rightRotation()
		}
		// RR
		n = n.leftRotation()
	}
	return n
}

func (n *AVLNode) insert(m *AVLNode) *AVLNode {
	if n == nil {
		return m
	}

	c := bytes.Compare(m.Key, n.Key)

	if c == -1 {
		n.left = n.left.insert(m)
	} else if c == 1 {
		n.right = n.right.insert(m)
	} else if c == 0 {
		n.Val = m.Val
	}

	n.updateHeight()
	return n.rebalance()
}

func (n *AVLNode) search(key []byte) *AVLNode {

	for current := n; current != nil; {
		c := bytes.Compare(key, current.Key)
		if c == -1 {
			current = current.left
		} else if c == 1 {
			current = current.right
		} else {
			return current
		}
	}
	// if not found
	return nil
}

func (n *AVLNode) delete(key []byte) *AVLNode {
	if n == nil {
		return nil
	}

	c := bytes.Compare(key, n.Key)

	if c == -1 {
		n.left = n.left.delete(key)
	} else if c == 1 {
		n.right = n.right.delete(key)
	} else {
		if n.right != nil {
			successor := n.right
			for successor.left != nil {
				successor = successor.left
			}
			n.Key = successor.Key
			n.Val = successor.Val
			n.right = n.right.delete(successor.Key)
		} else if n.left != nil {
			// node with only left child
			n = n.left
		} else {
			// node has no children
			n = nil
			return n
		}
	}
	if n != nil {
		n.updateHeight()
		n = n.rebalance()
	}
	return n
}

func (n *AVLNode) do(do func(*AVLNode)) {
	if n != nil {
		n.left.do(do)
		do(n)
		n.right.do(do)
	}
}
