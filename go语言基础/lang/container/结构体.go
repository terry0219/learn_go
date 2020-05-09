/*
结构体

Go语言可以通过自定义的方式形成新的类型，结构体就是这些类型中的一种复合类型，结构体是由零个或多个任意类型的值聚合成的实体，每个值都可以称为结构体的成员。

结构体的定义只是一种内存布局的描述，只有当结构体实例化时，才会真正地分配内存，因此必须在定义结构体并实例化后才能使用结构体的字段

type Point struct {
	x int
}
结构体实例化的三种方法:
1、var ins Point
2、ins := new(Point)
3、ins := &Point{}

使用键值对初始化结构体:
type People struct {
    name  string
    child *People          //结构体的结构体指针字段，类型是 *People 代表是People类型的指针
}
relation := &People{       //relation 由 People 类型取地址后，形成类型为 *People 的实例
    name: "爷爷",
    child: &People{
        name: "爸爸",
        child: &People{     //child 在初始化时，需要 *People 类型的值，使用取地址初始化一个 People
                name: "我",
        },
    },
}

使用多个值的列表初始化结构体:
type Address struct {
    Province    string
    City        string
    ZipCode     int
    PhoneNumber string
}
addr := Address{
    "四川",
    "成都",
    610000,
    "0",
}

*/
package main

import (
	"fmt"
)

type Point struct {
	x int
	y int
}

type Point2 struct {
	A int
	B string
}

func main() {
	var ins Point //基本的实例化, 将ins变量定义为Point类型
	ins.x = 10
	ins.y = 20
	fmt.Println(ins) // {10 20}

	ins2 := new(Point2) // 使用new函数进行实例化,返回的是指针类型的结构体
	ins2.A = 100
	ins2.B = "golang"
	fmt.Println(ins2)                                     //&{100 golang}
	fmt.Println(&ins2)                                    //0xc000006030  取变量的内存地址
	p := &ins2                                            // p是指针，指向了ins2变量的地址
	fmt.Printf("p is type: %T; p is value: %+v\n", p, *p) // p is type: **main.Point2; p is value: &{A:100 B:golang}

	ins3 := &Point2{}
	ins3.A = 200
	ins3.B = "python"
	fmt.Println(ins3) // &{200 python}
}
