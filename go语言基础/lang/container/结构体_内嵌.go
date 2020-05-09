//结构体内嵌
/*

结构体可以包含一个或多个匿名（或内嵌）字段，即这些字段没有显式的名字，只有字段的类型是必须的，此时类型也就是字段的名字。匿名字段本身可以是一个结构体类型，即结构体可以包含内嵌结构体。

在一个结构体中对于每一种数据类型只能有一个匿名字段

结构体内嵌特性:

1. 内嵌的结构体可以直接访问其成员变量
2. 内嵌结构体的字段名是它的类型名

内嵌结构体字段仍然可以使用详细的字段进行一层层访问
*/

package main

import "fmt"

type A struct {
	ax, ay int
}

type B struct {
	A
	bx, by int
	int    // anonymous field  int类型的匿名字段
	string // string类型的匿名字段
	bool   //bool类型的匿名字段
}

func main() {

	b := B{A{1, 2}, 3, 4, 5, "go", true} //实例化结构体B
	// b.ax = 10
	// b.ay = 20
	fmt.Println(b)                       //{{1 2} 3 4}
	fmt.Println(b.ax, b.ay, b.bx, b.by)  // 1 2 3 4
	fmt.Println(b.int, b.string, b.bool) //5 go true  访问匿名字段
	fmt.Println(b.A.ax)                  //1     内嵌结构体字段仍然可以使用详细的字段进行一层层访问
}
