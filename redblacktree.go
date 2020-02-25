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
			}
		}
	} else {
		if rbnode.leftNode != nil {
			rbnode.leftNode.addToSubtree(value)
		} else {
			rbnode.leftNode = &redBlackNode{
				value:  value,
				colour: red,
			}
		}
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
