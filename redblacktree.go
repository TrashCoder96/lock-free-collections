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

func (rbtree *RedBlackTree) rotateLeft(grandparent *redBlackNode) {
	parent := grandparent.rightNode
	grandgrandparent := grandparent.parent
	if grandgrandparent == nil {
		rbtree.head = parent
		rbtree.head.parent = nil
	} else {
		if grandgrandparent.leftNode == grandparent {
			grandgrandparent.leftNode = parent
			parent.parent = grandgrandparent
		} else if grandgrandparent.rightNode == grandparent {
			grandgrandparent.rightNode = parent
			parent.parent = grandgrandparent
		} else {
			panic("Impossible situation")
		}
	}
	grandparent.rightNode = parent.leftNode
	if parent.leftNode != nil {
		parent.leftNode.parent = grandparent
	}
	parent.leftNode = grandparent
	grandparent.parent = parent
}

func (rbtree *RedBlackTree) rotateRight(grandparent *redBlackNode) {
	parent := grandparent.leftNode
	grandgrandparent := grandparent.parent
	if grandgrandparent == nil {
		rbtree.head = parent
		rbtree.head.parent = nil
	} else {
		if grandgrandparent.leftNode == grandparent {
			grandgrandparent.leftNode = parent
			parent.parent = grandgrandparent
		} else if grandgrandparent.rightNode == grandparent {
			grandgrandparent.rightNode = parent
			parent.parent = grandgrandparent
		} else {
			panic("Impossible situation")
		}
	}
	grandparent.leftNode = parent.rightNode
	if parent.rightNode != nil {
		parent.rightNode.parent = grandparent
	}
	parent.rightNode = grandparent
	grandparent.parent = parent
}

//Delete func
func (rbtree *RedBlackTree) Delete(value int64) {

}

//Find func
func (rbtree *RedBlackTree) Find(value int64) bool {
	if rbtree.head == nil {
		return false
	}
	return rbtree.head.find(value)
}

func (rbnode *redBlackNode) find(value int64) bool {
	if rbnode.value == value {
		return true
	}
	if rbnode.leftNode != nil {
		return rbnode.leftNode.find(value)
	} else if rbnode.rightNode != nil {
		return rbnode.rightNode.find(value)
	} else {
		return false
	}
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
