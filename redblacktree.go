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
	value     int
	colour    int
	leftNode  *redBlackNode
	rightNode *redBlackNode
	parent    *redBlackNode
}

//Add func
func (rbtree *RedBlackTree) Add(value int) {
	if rbtree.head == nil {
		rbtree.head = &redBlackNode{
			colour: black,
			value:  value,
		}
	} else {
		rbtree.head.addToSubtree(value)
	}
}

func (rbnode *redBlackNode) addToSubtree(value int) {
	if rbnode.value == value {
		panic("Key already exists in red-black tree")
	}
	if value > rbnode.value {
		if rbnode.rightNode != nil {
			rbnode.rightNode.addToSubtree(value)
		} else {
			rbnode.rightNode = &redBlackNode{
				value:  value,
				colour: red,
				parent: rbnode,
			}
			rebalance(rbnode.rightNode)
		}
	} else {
		if rbnode.leftNode != nil {
			rbnode.leftNode.addToSubtree(value)
		} else {
			rbnode.leftNode = &redBlackNode{
				value:  value,
				colour: red,
				parent: rbnode,
			}
			rebalance(rbnode.leftNode)
		}
	}
}

func rebalance(node *redBlackNode) {
	uncle := node.getUncle()
	if uncle != nil {
		grandparent := node.getGrandparent()
		if uncle.colour == red {
			node.parent.colour = black
			uncle.colour = black
			grandparent.colour = red
			rebalance(grandparent)
		} else if uncle.colour == black {

		} else {
			panic("Impossible situation")
		}
	}
}

func (rbtree *RedBlackTree) leftLeftCase(grandparent *redBlackNode) {
	parent := grandparent.leftNode
	rotateRight(grandparent)
	//swap colours of parent and grandparent
	parent.colour = black
	grandparent.colour = red
}

func (rbtree *RedBlackTree) leftRightCase(grandparent *redBlackNode) {

}

func (rbtree *RedBlackTree) rightRightCase(grandparent *redBlackNode) {

}

func (rbtree *RedBlackTree) rightLeftCase(grandparent *redBlackNode) {

}

func rotateLeft(grandparent *redBlackNode) {
	//change structure of nodes
	parent := grandparent.rightNode
	changeStructureOfNodes(grandparent)
	grandparent.rightNode = parent.leftNode
	if parent.leftNode != nil {
		parent.leftNode.parent = grandparent
	}
	parent.leftNode = grandparent
	grandparent.parent = parent
}

func rotateRight(grandparent *redBlackNode) {
	//change structure of nodes
	parent := grandparent.leftNode
	changeStructureOfNodes(grandparent)
	grandparent.leftNode = parent.rightNode
	if parent.rightNode != nil {
		parent.rightNode.parent = grandparent
	}
	parent.rightNode = grandparent
	grandparent.parent = parent
}

func changeStructureOfNodes(grandparent *redBlackNode) {
	grandgrandparent := grandparent.parent
	parent := grandparent.leftNode
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

//Delete func
func (rbtree *RedBlackTree) Delete(value int) {

}

//Find func
func (rbtree *RedBlackTree) Find(value int) bool {
	if rbtree.head == nil {
		return false
	}
	return rbtree.head.find(value)
}

func (rbnode *redBlackNode) find(value int) bool {
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
