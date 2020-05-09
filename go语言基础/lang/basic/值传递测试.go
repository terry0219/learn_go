/*
参数传递测试
*/

package main

import "fmt"

type Data struct{    //将 Data 声明为结构体类型，结构体是拥有多个字段的复杂结构
	complax []int    //complax 为整型切片类型
	instance InnerData //instance 成员以 InnerData 类型作为 Data 的成员
	ptr *InnerData //将 ptr 声明为 InnerData 的指针类型
}

type InnerData struct {   //声明一个内嵌的结构 InnerData
	a int
}

//// 值传递测试函数
//passByValue() 函数用于值传递的测试，该函数的参数和返回值都是 Data 类型，在调用过程中，Data 的内存会被复制后传入函数
//传入结构体, 返回同类型的结构体
func passByValue(infunc Data) Data {
	fmt.Printf("inFunc value: %+v\n", infunc)
	fmt.Printf("inFunc ptr: %p\n", &infunc)
	return infunc
}


func main() {

	in := Data{                //创建一个 Data 结构的实例 in
		complax: []int{1,2,3},
		instance: InnerData{
			1,
		},
		ptr: &InnerData{2},
	}

	fmt.Printf("in value: %+v\n", in)
	fmt.Printf("in ptr: %p\n",&in)

	out := passByValue(in)
	fmt.Printf("out value: %+v\n", out)
	fmt.Printf("out ptr: %p\n", &out)

	/*
	输出:
	in value: {complax:[1 2 3] instance:{a:1} ptr:0xc0000100d0}
	in ptr: 0xc00006a330
	inFunc value: {complax:[1 2 3] instance:{a:1} ptr:0xc0000100d0}
	inFunc ptr: 0xc00006a3c0
	out value: {complax:[1 2 3] instance:{a:1} ptr:0xc0000100d0}
	out ptr: 0xc00006a390
	*/

}