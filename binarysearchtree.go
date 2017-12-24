package main

import (
	"fmt"
)

type Stack struct {
	list []*Node
	top  int
}

func (stack *Stack) push(x *Node) {

	if stack.top == len(stack.list)-1 {

		list := make([]*Node, len(stack.list)*2)

		for i, x := range stack.list {

			list[i] = x

		}

		stack.list = list
		stack.list[stack.top] = x

	} else {

		stack.top++
		stack.list[stack.top] = x

	}

}

func (stack *Stack) pop() {

	if stack.top >= 0 {

		stack.top--
	}

}

func (stack *Stack) peek() *Node {

	return stack.list[stack.top]

}

func (stack *Stack) isEmpty() bool {

	if stack.top == -1 {

		return true
	} else {

		return false
	}

}

type Queue struct {
	stack1 *Stack
	stack2 *Stack
}

func (que *Queue) enque(x *Node) {

	que.stack1.push(x)

}

func (que *Queue) deque() *Node {

	if que.stack2.isEmpty() {

		for !(que.stack1.isEmpty()) {

			que.stack2.push(que.stack1.peek())
			que.stack1.pop()

		}

		node := que.stack2.peek()
		que.stack2.pop()
		return node

	} else {

		node := que.stack2.peek()
		que.stack2.pop()
		return node

	}

}

func (que *Queue) isEmpty() bool {

	if que.stack1.isEmpty() && que.stack2.isEmpty() {

		return true

	} else {

		return false

	}

}

type Node struct {
	data  int
	left  *Node
	right *Node
}

type BinarySearchTree struct {
	root *Node
}



func (bst *BinarySearchTree) addNode(root *Node, data int) *Node {

	if root == nil {

		node := Node{}
		node.data = data
		root = &node
		return root

	} else if data <= root.data {

		if root.left == nil {

			node := Node{}
			node.data = data
			root.left = &node
			return root

		} else {

			bst.addNode(root.left, data)
			return root

		}

	} else {

		if root.right == nil {

			node := Node{}
			node.data = data
			root.right = &node
			return root

		} else {

			bst.addNode(root.right, data)
			return root

		}

	}

}

func (bst *BinarySearchTree) postOrder(root *Node) {

	if root == nil {

		return

	}

	bst.postOrder(root.left)
	bst.postOrder(root.right)
	fmt.Println(root.data)

}

func (bst *BinarySearchTree) inOrder(root *Node) {

	if root == nil {

		return

	}

	bst.inOrder(root.left)
	fmt.Println(root.data)
	bst.inOrder(root.right)

}

func (bst *BinarySearchTree) breadthFirstSearch(root *Node) {

	if root == nil {

		return

	}

	stack1 := Stack{list: make([]*Node, 10), top: -1}
	stack2 := Stack{list: make([]*Node, 10), top: -1}

	que := Queue{stack1: &stack1, stack2: &stack2}

	que.enque(root)

	for !(que.isEmpty()) {

		node := que.deque()

		fmt.Println(node.data)

		if node.left != nil {

			que.enque(node.left)

		}

		if node.right != nil {

			que.enque(node.right)

		}

	}

}

func (bst *BinarySearchTree) mirrorImage(root *Node) {

	if root == nil {

		return

	}

	stack1 := Stack{list: make([]*Node, 10), top: -1}
	stack2 := Stack{list: make([]*Node, 10), top: -1}

	que := Queue{stack1: &stack1, stack2: &stack2}

	que.enque(root)

	for !(que.isEmpty()) {

		node := que.deque()

		temp := node.left
		node.left = node.right
		node.right = temp

		if node.left != nil {

			que.enque(node.left)

		}

		if node.right != nil {

			que.enque(node.right)

		}

	}

}

func main() {


	bst := BinarySearchTree{}

	bst.root = bst.addNode(nil, 50)
	bst.root = bst.addNode(bst.root, 30)
	bst.root = bst.addNode(bst.root, 20)
	bst.root = bst.addNode(bst.root, 40)
	bst.root = bst.addNode(bst.root, 70)
	bst.root = bst.addNode(bst.root, 60)
	bst.root = bst.addNode(bst.root, 80)

        fmt.Println("Inorder Traversal")

	bst.inOrder(bst.root)

	fmt.Println("Breadth First Traversal")

	bst.breadthFirstSearch(bst.root)

	bst.mirrorImage(bst.root)

	fmt.Println("Breadth First Traversal of Mirror Image")

	bst.breadthFirstSearch(bst.root)

}
