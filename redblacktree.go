package main

const (
	red   = iota
	black = iota
)

//RedBlackTree type
type RedBlackTree struct {
	head *redBlackNode
}

type redBlackNode struct {
	value     int64
	colour    int
	leftNode  *redBlackNode
	rightNode *redBlackNode
	parent    *redBlackNode
}

func initRedBlackTree() *RedBlackTree {
	return &RedBlackTree{}
}

//Add func
func (rbtree *RedBlackTree) Add(value int64) {
	if rbtree.head == nil {
		rbtree.head = &redBlackNode{
			colour: black,
			value:  value,
		}
	} else {
		rbtree.addToSubtree(rbtree.head, value)
		rbtree.head.colour = black
	}
}

func (rbtree *RedBlackTree) addToSubtree(rbnode *redBlackNode, value int64) {
	if rbnode.value == value {
		panic("Key already exists in red-black tree")
	}
	if value > rbnode.value {
		if rbnode.rightNode != nil {
			rbtree.addToSubtree(rbnode.rightNode, value)
		} else {
			rbnode.rightNode = &redBlackNode{
				value:  value,
				colour: red,
				parent: rbnode,
			}
			rbtree.rebalance(rbnode.rightNode)
		}
	} else {
		if rbnode.leftNode != nil {
			rbtree.addToSubtree(rbnode.leftNode, value)
		} else {
			rbnode.leftNode = &redBlackNode{
				value:  value,
				colour: red,
				parent: rbnode,
			}
			rbtree.rebalance(rbnode.leftNode)
		}
	}
}

func (rbtree *RedBlackTree) rebalance(node *redBlackNode) {
	if node.parent == nil {
		return
	}
	grandparent := node.getGrandparent()
	if grandparent == nil {
		return
	}
	uncle := node.getUncle()
	uncleIsRed := uncle != nil && uncle.colour == red
	uncleIsBlack := uncle == nil || uncle.colour == black
	if uncleIsRed {
		node.parent.colour = black
		uncle.colour = black
		grandparent.colour = red
		rbtree.rebalance(grandparent)
	} else if uncleIsBlack {
		if grandparent.leftNode != nil && grandparent.leftNode.colour == red &&
			grandparent.leftNode.leftNode != nil && grandparent.leftNode.leftNode.colour == red {
			rbtree.leftLeftCase(grandparent)
		} else if grandparent.leftNode != nil && grandparent.leftNode.colour == red &&
			grandparent.leftNode.rightNode != nil && grandparent.leftNode.rightNode.colour == red {
			rbtree.leftRightCase(grandparent)
		} else if grandparent.rightNode != nil && grandparent.rightNode.colour == red &&
			grandparent.rightNode.rightNode != nil && grandparent.rightNode.rightNode.colour == red {
			rbtree.rightRightCase(grandparent)
		} else if grandparent.rightNode != nil && grandparent.rightNode.colour == red &&
			grandparent.rightNode.leftNode != nil && grandparent.rightNode.leftNode.colour == red {
			rbtree.rightLeftCase(grandparent)
		}
	} else {
		panic("Impossible situation")
	}
}

func (rbtree *RedBlackTree) leftLeftCase(grandparent *redBlackNode) {
	parent := grandparent.leftNode
	rbtree.rotateRight(grandparent)
	//swap colours of parent and grandparent
	parent.colour = black
	grandparent.colour = red
}

func (rbtree *RedBlackTree) leftRightCase(grandparent *redBlackNode) {
	parent := grandparent.leftNode
	rbtree.rotateLeft(parent)
	rbtree.rotateRight(grandparent)
	//swap colours of parent and grandparent
	parent.colour = black
	grandparent.colour = red
}

func (rbtree *RedBlackTree) rightRightCase(grandparent *redBlackNode) {
	parent := grandparent.rightNode
	rbtree.rotateLeft(grandparent)
	//swap colours of parent and grandparent
	parent.colour = black
	grandparent.colour = red
}

func (rbtree *RedBlackTree) rightLeftCase(grandparent *redBlackNode) {
	parent := grandparent.rightNode
	rbtree.rotateRight(parent)
	rbtree.rotateLeft(grandparent)
	//swap colours of parent and grandparent
	parent.colour = black
	grandparent.colour = red
}

func (rbtree *RedBlackTree) rotateLeft(pointOfRotating *redBlackNode) {
	rightChild := pointOfRotating.rightNode
	parent := pointOfRotating.parent
	if parent == nil {
		rbtree.head = rightChild
		rbtree.head.parent = nil
	} else {
		rightChild.parent = parent
		if parent.leftNode == pointOfRotating {
			parent.leftNode = rightChild
		} else if parent.rightNode == pointOfRotating {
			parent.rightNode = rightChild
		} else {
			panic("Impossible situation")
		}
	}
	pointOfRotating.rightNode = rightChild.leftNode
	if rightChild.leftNode != nil {
		rightChild.leftNode.parent = pointOfRotating
	}
	rightChild.leftNode = pointOfRotating
	pointOfRotating.parent = rightChild
}

func (rbtree *RedBlackTree) rotateRight(pointOfRotating *redBlackNode) {
	leftChild := pointOfRotating.leftNode
	parent := pointOfRotating.parent
	if parent == nil {
		rbtree.head = leftChild
		rbtree.head.parent = nil
	} else {
		leftChild.parent = parent
		if parent.leftNode == pointOfRotating {
			parent.leftNode = leftChild
		} else if parent.rightNode == pointOfRotating {
			parent.rightNode = leftChild
		} else {
			panic("Impossible situation")
		}
	}
	pointOfRotating.leftNode = leftChild.rightNode
	if leftChild.rightNode != nil {
		leftChild.rightNode.parent = pointOfRotating
	}
	leftChild.rightNode = pointOfRotating
	pointOfRotating.parent = leftChild
}

//Delete func
func (rbtree *RedBlackTree) Delete(value int64) bool {
	if rbtree.head == nil {
		return false
	}
	foundNode := rbtree.head.find(value)
	if foundNode == nil {
		return false
	}
	rbtree.delete(foundNode)
	return true
}

func (rbtree *RedBlackTree) delete(rbnode *redBlackNode) {
	if rbnode.leftNode == nil && rbnode.rightNode == nil {
		if rbnode.parent != nil {
			if rbnode.parent.leftNode == rbnode {
				rbnode.parent.leftNode = nil
			} else if rbnode.parent.rightNode == nil {
				rbnode.parent.rightNode = nil
			} else {
				panic("Impossible situation")
			}
		} else {
			rbtree.head = nil
		}
	} else if rbnode.leftNode != nil && rbnode.rightNode == nil {
		if rbnode.parent != nil {
			if rbnode.parent.leftNode == rbnode {
				rbnode.parent.leftNode = rbnode.leftNode
			} else if rbnode.parent.rightNode == rbnode {
				rbnode.parent.rightNode = rbnode.leftNode
			} else {
				panic("Impossible situation")
			}
			rbnode.leftNode.parent = rbnode.parent
		} else {
			rbtree.head = rbnode.leftNode
			rbnode.leftNode.parent = nil
		}
	} else if rbnode.leftNode == nil && rbnode.rightNode != nil {
		if rbnode.parent != nil {
			if rbnode.parent.leftNode == rbnode {
				rbnode.parent.leftNode = rbnode.rightNode
			} else if rbnode.parent.rightNode == rbnode {
				rbnode.parent.rightNode = rbnode.rightNode
			} else {
				panic("Impossible situation")
			}
			rbnode.rightNode.parent = rbnode.parent
		} else {
			rbtree.head = rbnode.rightNode
			rbnode.rightNode.parent = nil
		}
	} else {
		leafNode := rbnode.leftNode
		for leafNode.leftNode != nil {
			leafNode = leafNode.leftNode
		}
		rbnode.value = leafNode.value
		rbtree.delete(leafNode)
	}
}

//Find func
func (rbtree *RedBlackTree) Find(value int64) bool {
	if rbtree.head == nil {
		return false
	}
	return rbtree.head.find(value) != nil
}

func (rbnode *redBlackNode) find(value int64) *redBlackNode {
	if rbnode == nil {
		return nil
	}
	if rbnode.value == value {
		return rbnode
	}
	if rbnode.value > value {
		return rbnode.leftNode.find(value)
	}
	return rbnode.rightNode.find(value)
}

func (rbnode *redBlackNode) getGrandparent() *redBlackNode {
	if rbnode.parent != nil && rbnode.parent.parent != nil {
		return rbnode.parent.parent
	}
	return nil
}

func (rbnode *redBlackNode) getUncle() *redBlackNode {
	grandparent := rbnode.getGrandparent()
	if grandparent != nil {
		if rbnode.parent == grandparent.rightNode {
			return grandparent.leftNode
		}
		return grandparent.rightNode
	}
	return nil
}
