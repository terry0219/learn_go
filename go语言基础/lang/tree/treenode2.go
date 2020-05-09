package main

import "fmt"

//tree结构
type treenode struct {
	value       int
	left, right *treenode
}

func (node *treenode) print() {
	if node == nil {
		fmt.Println("nil=node")
	} else {
		fmt.Println(node.value)
	}

}
func (node *treenode) setVal(val int) {
	node.value = val
}

//递归函数：数据结构是递归的，操作方法是递归的
func (node *treenode) recu() {
	if node != nil {
		node.print()
		node.left.recu()
		fmt.Println("------------")
		node.right.recu()

	} else {
		return
	}

}

func main() {
	root := treenode{}
	root.setVal(0)
	root.left = &treenode{value: 1}
	root.right = &treenode{value: 2}
	root.left.left = &treenode{value: 3}
	root.left.right = &treenode{value: 4}
	root.right.left = &treenode{value: 5}
	root.right.right = &treenode{value: 6}

	fmt.Printf("Type:%T ; Value:%+v", root, root) //Type:main.treenode ; Value:{value:0 left:0xc0000044a0 right:0xc0000044c0}
	fmt.Println()

	root.recu()
	/*
		树形结构:
									0
				1											2
			3		4									5		6


	*/

	/*
		0
		1
		3
		------------
		------------
		4
		------------
		------------
		2
		5
		------------
		------------
		6
		------------
	*/
}
