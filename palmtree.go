package main

//PalmTree type
type PalmTree struct {
	order int
	root  *palmNode
}

type palmNode struct {
	countOfKeys      int
	isLeaf           bool
	internalNodeHead *palmTreePointer //only for internal node
	leafHead         *palmTreeKey     //only for leaf node
}

type palmTreeKey struct {
	value           int64
	nextPointer     *palmTreePointer
	nextKey         *palmTreeKey
	previousKey     *palmTreeKey
	previousPointer *palmTreePointer
}

type palmTreePointer struct {
	nextKey     *palmTreeKey
	previousKey *palmTreeKey
	childNode   *palmNode
}

//Insert function
func (ptree *PalmTree) Insert(key int64, value string) {
	if ptree.root.leafHead == nil && ptree.root.internalNodeHead == nil {
		ptree.root.leafHead = &palmTreeKey{
			value: key,
		}
		ptree.root.countOfKeys = 1
	} else {
		rootPointer := palmTreePointer{childNode: ptree.root}
		ptree.insert(key, value, &rootPointer)
		if rootPointer.nextKey != nil {
			newNode := palmNode{
				internalNodeHead: &rootPointer,
				countOfKeys:      1,
			}
			ptree.root = &newNode
		}
	}
}

//Delete function
func (ptree *PalmTree) Delete(key int64) bool {
	if ptree.delete(key, ptree.root) {
		if ptree.root.internalNodeHead != nil && ptree.root.internalNodeHead.nextKey == nil {
			ptree.root = ptree.root.internalNodeHead.childNode
		}
		return true
	}
	return false
}

//Find function
func (ptree *PalmTree) Find(key int64) bool {
	return ptree.find(key, ptree.root)
}

//function returns true, if cutting has occured
func (ptree *PalmTree) insert(key int64, value string, pointerToNode *palmTreePointer) bool {
	if pointerToNode.childNode.isLeaf {
		pointerToNode.childNode.insertToLeafNode(key, value)
		if pointerToNode.childNode.countOfKeys > 2*ptree.order-1 {
			cutIfPossible(pointerToNode)
			return true
		}
	} else {
		suitablePointer := pointerToNode.childNode.getPointer(key)
		if ptree.insert(key, value, suitablePointer) {
			pointerToNode.childNode.countOfKeys = pointerToNode.childNode.countOfKeys + 1
		}
		if pointerToNode.childNode.countOfKeys > 2*ptree.order-1 {
			cutIfPossible(pointerToNode)
			return true
		}
	}
	return false
}

func (pnode *palmNode) getPointer(key int64) *palmTreePointer {
	if !pnode.isLeaf {
		currentPointer := pnode.internalNodeHead
		nextKeyValueMoreThanKey := false
		nextKeyIsNil := false
		for {
			nextKeyIsNil = currentPointer.nextKey == nil
			if !nextKeyIsNil {
				nextKeyValueMoreThanKey = currentPointer.nextKey.value > key
			}
			if nextKeyValueMoreThanKey || nextKeyIsNil {
				break
			} else {
				currentPointer = currentPointer.nextKey.nextPointer
			}
		}
		if nextKeyIsNil && !nextKeyValueMoreThanKey {
			return currentPointer
		} else if nextKeyValueMoreThanKey {
			return currentPointer
		} else {
			panic("Operation is not allowed!!!")
		}
	} else {
		panic("Operation is not allowed!!!")
	}
}

func (pnode *palmNode) insertToLeafNode(key int64, value string) {
	newLeaf := palmTreeKey{value: key}
	if pnode.leafHead == nil {
		pnode.leafHead = &newLeaf
	} else {
		leafBeforeNewLeaf := pnode.leafHead
		for leafBeforeNewLeaf.nextKey != nil && leafBeforeNewLeaf.nextKey.value <= key {
			leafBeforeNewLeaf = leafBeforeNewLeaf.nextKey
		}
		newLeaf = palmTreeKey{
			value: key,
		}
		if pnode.leafHead.value < key {
			if leafBeforeNewLeaf.nextKey == nil {
				leafBeforeNewLeaf.nextKey = &newLeaf
				newLeaf.previousKey = leafBeforeNewLeaf
			} else {
				newLeaf.nextKey = leafBeforeNewLeaf.nextKey
				newLeaf.previousKey = leafBeforeNewLeaf
				leafBeforeNewLeaf.nextKey.previousKey = &newLeaf
				leafBeforeNewLeaf.nextKey = &newLeaf
			}
		} else {
			newLeaf.nextKey = pnode.leafHead
			pnode.leafHead.previousKey = &newLeaf
			pnode.leafHead = &newLeaf
		}
	}
	pnode.countOfKeys = pnode.countOfKeys + 1
}

func cutIfPossible(pointer *palmTreePointer) {
	leftPointer := pointer
	newKey := palmTreeKey{
		previousPointer: leftPointer,
	}
	rightPointer := palmTreePointer{
		previousKey: &newKey,
		nextKey:     pointer.nextKey,
	}
	newKey.nextPointer = &rightPointer
	if pointer.nextKey != nil {
		pointer.nextKey.previousPointer = &rightPointer
	}
	leftPointer.nextKey = &newKey
	leftNode := pointer.childNode
	if pointer.childNode.isLeaf {
		rightNode := palmNode{
			isLeaf:      true,
			countOfKeys: leftNode.countOfKeys / 2,
		}
		keyBeforeNextNode := leftNode.leafHead
		for i := 1; i < leftNode.countOfKeys/2; i++ {
			keyBeforeNextNode = keyBeforeNextNode.nextKey
		}
		leftNode.countOfKeys = rightNode.countOfKeys
		rightPointer.childNode = &rightNode
		newKey.value = keyBeforeNextNode.nextKey.value
		rightNode.leafHead = keyBeforeNextNode.nextKey
		rightNode.leafHead.previousKey = nil
		keyBeforeNextNode.nextKey = nil
	} else {
		rightNode := palmNode{
			isLeaf:      false,
			countOfKeys: leftNode.countOfKeys / 2,
		}
		pointerBeforeMiddleKey := leftNode.internalNodeHead
		for i := 1; i < leftNode.countOfKeys/2; i++ {
			pointerBeforeMiddleKey = pointerBeforeMiddleKey.nextKey.nextPointer
		}
		leftNode.countOfKeys = rightNode.countOfKeys - 1
		rightPointer.childNode = &rightNode
		rightNode.internalNodeHead = pointerBeforeMiddleKey.nextKey.nextPointer
		rightNode.internalNodeHead.previousKey = nil
		newKey.value = pointerBeforeMiddleKey.nextKey.value
		pointerBeforeMiddleKey.nextKey = nil
	}
}

func (ptree *PalmTree) find(key int64, node *palmNode) bool {
	if node != nil {
		if node.isLeaf {
			leaf := node.leafHead
			for leaf != nil && leaf.value != key {
				leaf = leaf.nextKey
			}
			if leaf == nil {
				return false
			}
			return true
		}
		pointer := node.getPointer(key)
		return ptree.find(key, pointer.childNode)
	}
	panic("Operation is not allowed")
}

func (ptree *PalmTree) delete(key int64, node *palmNode) bool {
	if node != nil {
		if node.isLeaf {
			return node.deleteFromLeafNode(key)
		}
		currentPointer := node.getPointer(key)
		success := ptree.delete(key, currentPointer.childNode)
		if success {
			if currentPointer.nextKey == nil {
				ptree.redistributeNodesIfPossible(currentPointer.previousKey.previousPointer, node)
			} else {
				ptree.redistributeNodesIfPossible(currentPointer, node)
			}
		}
		return success
	}
	panic("Operation is not allowed")
}

func (ptree *PalmTree) redistributeNodesIfPossible(subtree *palmTreePointer, node *palmNode) {
	leftPointer := subtree
	middleKey := leftPointer.nextKey
	rightPointer := middleKey.nextPointer
	leftPointerChildNodeLessThanOrderMinusOne := leftPointer.childNode.countOfKeys <= ptree.order-1
	rightPointerChildNodeLessThanOrderMinusOne := rightPointer.childNode.countOfKeys <= ptree.order-1
	if leftPointerChildNodeLessThanOrderMinusOne && rightPointerChildNodeLessThanOrderMinusOne {
		merge(subtree)
		node.countOfKeys = node.countOfKeys - 1
	} else if leftPointerChildNodeLessThanOrderMinusOne && !rightPointerChildNodeLessThanOrderMinusOne {
		moveToLeftNode(subtree)
	} else if !leftPointerChildNodeLessThanOrderMinusOne && rightPointerChildNodeLessThanOrderMinusOne {
		moveToRightNode(subtree)
	}
}

func moveToLeftNode(subtree *palmTreePointer) {
	leftPointer := subtree
	middleKey := leftPointer.nextKey
	rightPointer := middleKey.nextPointer
	lowLevelIsLeaves := leftPointer.childNode.isLeaf && rightPointer.childNode.isLeaf
	if lowLevelIsLeaves {
		movedItem := rightPointer.childNode.leafHead
		rightPointer.childNode.leafHead = rightPointer.childNode.leafHead.nextKey
		rightPointer.childNode.leafHead.previousKey = nil
		movedItem.nextKey = nil
		tailKey := leftPointer.childNode.leafHead
		for tailKey.nextKey != nil {
			tailKey = tailKey.nextKey
		}
		tailKey.nextKey = movedItem
		movedItem.previousKey = tailKey
		middleKey.value = rightPointer.childNode.leafHead.value
	} else {
		tailPointer := leftPointer.childNode.internalNodeHead
		for tailPointer.nextKey != nil {
			tailPointer = tailPointer.nextKey.nextPointer
		}
		newKey := palmTreeKey{
			value:           middleKey.value,
			nextPointer:     rightPointer.childNode.internalNodeHead,
			previousPointer: tailPointer,
		}
		tailPointer.nextKey = &newKey
		rightPointer.childNode.internalNodeHead.previousKey = &newKey
		middleKey.value = rightPointer.childNode.internalNodeHead.nextKey.value
		rightPointer.childNode.internalNodeHead = rightPointer.childNode.internalNodeHead.nextKey.nextPointer
		newKey.nextPointer.nextKey = nil
	}
	leftPointer.childNode.countOfKeys = leftPointer.childNode.countOfKeys + 1
	rightPointer.childNode.countOfKeys = rightPointer.childNode.countOfKeys - 1
}

func moveToRightNode(subtree *palmTreePointer) {
	leftPointer := subtree
	middleKey := leftPointer.nextKey
	rightPointer := middleKey.nextPointer
	lowLevelIsLeaves := leftPointer.childNode.isLeaf && rightPointer.childNode.isLeaf
	if lowLevelIsLeaves {
		tailKey := leftPointer.childNode.leafHead
		for tailKey.nextKey != nil {
			tailKey = tailKey.nextKey
		}
		tailKey.previousKey.nextKey = nil
		tailKey.previousKey = nil
		tailKey.nextKey = rightPointer.childNode.leafHead
		rightPointer.childNode.leafHead.previousKey = tailKey
		rightPointer.childNode.leafHead = tailKey
		middleKey.value = rightPointer.childNode.leafHead.value
	} else {
		tailPointer := leftPointer.childNode.internalNodeHead
		for tailPointer.nextKey != nil {
			tailPointer = tailPointer.nextKey.nextPointer
		}
		newKey := palmTreeKey{
			value:           middleKey.value,
			nextPointer:     rightPointer.childNode.internalNodeHead,
			previousPointer: tailPointer,
		}
		middleKey.value = tailPointer.previousKey.value
		tailPointer.nextKey = &newKey
		tailPointer.previousKey.previousPointer.nextKey = nil
		tailPointer.previousKey = nil
		rightPointer.childNode.internalNodeHead.previousKey = &newKey
		rightPointer.childNode.internalNodeHead = newKey.previousPointer
	}
	leftPointer.childNode.countOfKeys = leftPointer.childNode.countOfKeys - 1
	rightPointer.childNode.countOfKeys = rightPointer.childNode.countOfKeys + 1
}

func merge(subtree *palmTreePointer) {
	leftPointer := subtree
	middleKey := leftPointer.nextKey
	rightPointer := middleKey.nextPointer
	lowLevelIsLeaves := leftPointer.childNode.isLeaf && rightPointer.childNode.isLeaf
	leftPointer.childNode.countOfKeys += rightPointer.childNode.countOfKeys
	if lowLevelIsLeaves {
		tailLeaf := leftPointer.childNode.leafHead
		for tailLeaf.nextKey != nil {
			tailLeaf = tailLeaf.nextKey
		}
		tailLeaf.nextKey = rightPointer.childNode.leafHead
		rightPointer.childNode.leafHead.previousKey = tailLeaf
		leftPointer.nextKey = rightPointer.nextKey
		if rightPointer.nextKey != nil {
			rightPointer.nextKey.previousPointer = leftPointer
		}
	} else {
		tailPointer := leftPointer.childNode.internalNodeHead
		for tailPointer.nextKey != nil {
			tailPointer = tailPointer.nextKey.nextPointer
		}
		tailPointer.nextKey = middleKey
		middleKey.previousPointer = tailPointer
		middleKey.nextPointer = rightPointer.childNode.internalNodeHead
		rightPointer.childNode.internalNodeHead.previousKey = middleKey
		leftPointer.nextKey = rightPointer.nextKey
		if rightPointer.nextKey != nil {
			rightPointer.nextKey.previousPointer = leftPointer
		}
		leftPointer.childNode.countOfKeys++
	}
}

func (pnode *palmNode) deleteFromLeafNode(key int64) bool {
	currentLeaf := pnode.leafHead
	for currentLeaf != nil {
		if currentLeaf.value == key {
			if currentLeaf.previousKey == nil {
				pnode.leafHead = currentLeaf.nextKey
				if currentLeaf.nextKey != nil {
					currentLeaf.nextKey.previousKey = nil
				}
			} else if currentLeaf.nextKey == nil {
				currentLeaf.previousKey.nextKey = nil
			} else {
				currentLeaf.previousKey.nextKey = currentLeaf.nextKey
				currentLeaf.nextKey.previousKey = currentLeaf.previousKey
			}
			pnode.countOfKeys = pnode.countOfKeys - 1
			return true
		}
		currentLeaf = currentLeaf.nextKey
	}
	return false
}