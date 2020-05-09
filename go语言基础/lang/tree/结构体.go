/*
go语言仅支持封装，不支持继承和多态
go语言没有class, 只有struct
不论地址还是结构本身，一律使用.来访问成员

一般来说局部变量是分配在栈上，函数一旦退出后，这个局部变量就会被销毁；如果要传出去用，必须要在堆上分配，但是堆上分配要手动释放(C++的做法)
go语言中变量的分配在哪里是由编译器决定的，如果变量没有取指针地址，也没有返回出去，编译器会认为这个变量不需要给外部，所以在栈上分配；当编译器看到变量取了指针地址并且返回出去给外部用，
那么这个变量就会在堆上分配；堆上分配后，这个变量就要参与垃圾回收；等外部调用这个指针地址结束后，就会被垃圾回收

在go语言，函数退出后，变量不一定就会被销毁，要等外部调用结束后，在回收

nil指针也可以调用方法
*/
package main

import "fmt"

//定义名为treeNode的结构体, 一共有3个值，value为int类型，left right为treeNode类型的指针
type treeNode struct {
	value       int
	left, right *treeNode
}

//为treeNode类型定义print方法；(node treeNode)代表接收器；print()代表方法名; print方法有个接收者是node; 就是一个普通的函数
func (node treeNode) print() { //显示定义和命名方法接收者
	fmt.Println(node.value)
}

/*
另外一种写法
func print(node treeNode) {
	fmt.Println(node.value)
}
print(root)
*/

func (node treeNode) setValue(value int) { //go语言函数的参数传递都是值传递;也就是无论对在函数内部对treeNode如何改变，只在函数内有效。不会影响函数外的treeNode
	node.value = value
}

//下面是使用指针作为方法接收者，只有指针才可以改变结构内容
//func (node *treeNode) setValue(value int) {}  如果类型改为*treeNode就是引用传递了，在函数内修改treeNode后，在函数外也对应的修改

//定义traverse方法，遍历tree
func (node *treeNode) traverse() {
	if node == nil { //nil指针也可以调用方法;这里要判断下node指针是否为nil, 如果为nil的话，程序会抛异常，因为nil的方法取不到值，这时候需要return,以防报错
		return
	}
	node.left.traverse()  //打印左子树
	node.print()          //
	node.right.traverse() //打印右子树
}

//构造函数;使用自定义工厂函数
func createNode(value int) *treeNode {
	return &treeNode{
		value: value,
	}
}

func main() {
	var root treeNode //定义变量root为treeNode类型
	//fmt.Println(root) // {0 nil nil}

	root = treeNode{value: 3} //初始化treeNode
	root.left = &treeNode{}   //因为left值为treeNode类型的指针，所以要加&；
	root.right = &treeNode{5, nil, nil}
	//不论地址还是结构本身，一律使用.来访问成员
	root.right.left = new(treeNode) //使用new返回的是treeNode类型的指针，也可以初始化treeNode；root.right.left 因为right的值为treeNode类型的指针，所以可以一直取值，比如root.right.left.right,可以无限取值下去
	root.left.right = createNode(2) //使用createNode构造函数初始化

	//创建一个slice，里面数据是treeNode类型
	// nodes := []treeNode{
	// 	{value: 3},
	// 	{},
	// 	{6, nil, &root},
	// }
	// fmt.Println(nodes)

	/*
				3
		0       		5
			2		0

	*/
	// root.right.left.setValue(4) //虽然在setValue函数内修改了root.right.left这个treeNode类型的变量value，但是下面打印value值还是为0
	// root.right.left.print()     //0 调用print方法

	// pRoot := &root
	// pRoot.print()       //3
	// pRoot.setValue(200) //因为setValue函数接收器是指传递，所以变量只能函数内生效
	// pRoot.print()       //3

	fmt.Println()
	root.traverse()

	/*
		总结:
			1、要改变内容必须使用指针接收者
			2、结构过大也考虑使用指针接收者
			3、一致性: 如有指针接收者, 最好都是指针接收者

	*/
}
