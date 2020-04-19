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
			rbtree.rebalanceInsertCase1(rbnode.rightNode)
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
			rbtree.rebalanceInsertCase1(rbnode.leftNode)
		}
	}
}

func (rbtree *RedBlackTree) rebalanceInsertCase1(node *redBlackNode) {
	if node.parent == nil {
		node.colour = black
	} else {
		rbtree.rebalanceInsertCase2(node)
	}
}

func (rbtree *RedBlackTree) rebalanceInsertCase2(node *redBlackNode) {
	if !node.parent.isBlack() {
		rbtree.rebalanceInsertCase3(node)
	}
}

func (rbtree *RedBlackTree) rebalanceInsertCase3(node *redBlackNode) {
	uncle := node.getUncle()
	if uncle.isRed() {
		node.parent.colour = black
		uncle.colour = black
		grantparent := node.getGrandparent()
		grantparent.colour = red
		rbtree.rebalanceInsertCase1(grantparent)
	} else {
		rbtree.rebalanceInsertCase4(node)
	}
}

func (rbtree *RedBlackTree) rebalanceInsertCase4(node *redBlackNode) {
	grandparent := node.getGrandparent()
	if node == node.parent.rightNode && node.parent == grandparent.leftNode {
		rbtree.rotateLeft(node.parent)
		node = node.leftNode
	} else if node == node.parent.leftNode && node.parent == grandparent.rightNode {
		rbtree.rotateRight(node.parent)
		node = node.rightNode
	}
	rbtree.rebalanceInsertCase5(node)
}

func (rbtree *RedBlackTree) rebalanceInsertCase5(node *redBlackNode) {
	grandparent := node.getGrandparent()
	node.parent.colour = black
	grandparent.colour = red
	if node == node.parent.leftNode && node.parent == grandparent.leftNode {
		rbtree.rotateRight(grandparent)
	} else {
		rbtree.rotateLeft(grandparent)
	}
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
	if foundNode.leftNode != nil && foundNode.rightNode != nil {
		leafNode := foundNode.rightNode
		for leafNode.leftNode != nil {
			leafNode = leafNode.leftNode
		}
		foundNode.value = leafNode.value
		rbtree.delete(leafNode)
	} else {
		rbtree.delete(foundNode)
	}
	return true
}

func (rbtree *RedBlackTree) delete(rbnode *redBlackNode) {
	child := rbnode.getChild()
	if rbnode.isRed() {
		if rbnode.parent != nil {
			if rbnode.parent.leftNode == rbnode {
				rbnode.parent.leftNode = child
			} else {
				rbnode.parent.rightNode = child
			}
		} else {
			rbtree.head = child
		}
		if child != nil {
			child.parent = rbnode.parent
		}
	} else {
		if child.isRed() {
			child.colour = black
			rbnode.colour = red
			rbtree.delete(rbnode)
		} else {
			rbtree.rebalanceDeleteCase1(rbnode)
			if rbnode.parent != nil {
				if rbnode.parent.leftNode == rbnode {
					rbnode.parent.leftNode = child
				} else {
					rbnode.parent.rightNode = child
				}
			} else {
				rbtree.head = child
			}
			if child != nil {
				child.parent = rbnode.parent
			}
		}
	}
}

func (rbtree *RedBlackTree) rebalanceDeleteCase1(node *redBlackNode) {
	if node.parent != nil {
		rbtree.rebalanceDeleteCase2(node)
	}
}

func (rbtree *RedBlackTree) rebalanceDeleteCase2(node *redBlackNode) {
	subling := node.getSibling()
	if subling.isRed() {
		node.parent.colour = red
		subling.colour = black
		if node == node.parent.leftNode {
			rbtree.rotateLeft(node.parent)
		} else {
			rbtree.rotateRight(node.parent)
		}
	}
	rbtree.rebalanceDeleteCase3(node)
}

func (rbtree *RedBlackTree) rebalanceDeleteCase3(node *redBlackNode) {
	sibling := node.getSibling()
	if sibling != nil && node.parent.isBlack() &&
		sibling.isBlack() &&
		sibling.leftNode.isBlack() &&
		sibling.rightNode.isBlack() {
		sibling.colour = red
		rbtree.rebalanceDeleteCase1(node.parent)
	} else {
		rbtree.rebalanceDeleteCase4(node)
	}
}

func (rbtree *RedBlackTree) rebalanceDeleteCase4(node *redBlackNode) {
	sibling := node.getSibling()
	if sibling != nil && node.parent.isRed() &&
		sibling.isBlack() &&
		sibling.leftNode.isBlack() &&
		sibling.rightNode.isBlack() {
		sibling.colour = red
		node.parent.colour = black
	} else {
		rbtree.rebalanceDeleteCase5(node)
	}
}

func (rbtree *RedBlackTree) rebalanceDeleteCase5(node *redBlackNode) {
	sibling := node.getSibling()
	if sibling != nil && sibling.isBlack() {
		if node == node.parent.leftNode &&
			sibling != nil &&
			sibling.rightNode.isBlack() &&
			sibling.leftNode.isRed() {
			sibling.colour = red
			sibling.leftNode.colour = black
			rbtree.rotateRight(sibling)
		} else if node == node.parent.rightNode &&
			sibling.leftNode.isBlack() &&
			sibling.rightNode.isRed() {
			sibling.colour = red
			sibling.rightNode.colour = black
			rbtree.rotateLeft(sibling)
		}
	}
	rbtree.rebalanceDeleteCase6(node)
}

func (rbtree *RedBlackTree) rebalanceDeleteCase6(node *redBlackNode) {
	sibling := node.getSibling()
	sibling.colour = node.parent.colour
	node.parent.colour = black
	if node == node.parent.leftNode {
		sibling.rightNode.colour = black
		rbtree.rotateLeft(node.parent)
	} else {
		sibling.leftNode.colour = black
		rbtree.rotateRight(node.parent)
	}
}

func (rbnode *redBlackNode) getChild() *redBlackNode {
	var child *redBlackNode
	if rbnode.leftNode != nil {
		child = rbnode.leftNode
	} else {
		child = rbnode.rightNode
	}
	return child
}

func (rbnode *redBlackNode) isBlack() bool {
	return rbnode == nil || rbnode.colour == black
}

func (rbnode *redBlackNode) isRed() bool {
	return rbnode != nil && rbnode.colour == red
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

func (rbnode *redBlackNode) getSibling() *redBlackNode {
	parent := rbnode.parent
	if parent.leftNode == rbnode {
		return parent.rightNode
	} else if parent.rightNode == rbnode {
		return parent.leftNode
	}
	return nil
}
