package main

import (
	"fmt"

	"imooc.com/ccmouse/learngo/lang/tree"   //目录 $GOPATH/imooc.com/ccmouse/learngo/lang/tree
)

//扩展已有类型;为tree包扩展方法，继承tree.Node
type myTreeNode struct {
	node *tree.Node        //*tree.Node 指针类型; tree是package名称,Node是这个tree包下的struct
}

//为myTreeNode类型扩展方法postOrde
func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}

	left := myTreeNode{myNode.node.Left}  //实例化myTreeNode
	right := myTreeNode{myNode.node.Right}

	left.postOrder() //调用postOrder方法
	right.postOrder()
	myNode.node.Print()
}

func main() {
	var root tree.Node

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)

	fmt.Print("In-order traversal: ")
	root.Traverse()

	fmt.Print("My own post-order traversal: ")
	myRoot := myTreeNode{&root}
	myRoot.postOrder()
	fmt.Println()

	nodeCount := 0
	root.TraverseFunc(func(node *tree.Node) {
		nodeCount++
	})
	fmt.Println("Node count:", nodeCount)

	c := root.TraverseWithChannel()
	maxNodeValue := 0
	for node := range c {
		if node.Value > maxNodeValue {
			maxNodeValue = node.Value
		}
	}
	fmt.Println("Max node value:", maxNodeValue)
}
