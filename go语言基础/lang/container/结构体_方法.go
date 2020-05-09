//结构体方法
/*
Go 方法是作用在接收器（receiver）上的一个函数，接收器是某种类型的变量。
接收器类型可以是（几乎）任何类型，不仅仅是结构体类型，任何类型都可以有方法，甚至可以是函数类型，可以是 int、bool、string 或数组的别名类型，但是接收器不能是一个接口类型，因为接口是一个抽象定义，而方法却是具体实现

接收器的格式:
func (接收器变量 接收器类型) 方法名(参数列表) (返回参数) {
    函数体
}
接收器根据接收器的类型可以分为指针接收器、非指针接收器



*/
package main

import "fmt"

type Bag struct {
	items []string
}

//为Bag这个结构体，定义一个Insert方法，这个方法要传一个string类型的参数，返回一个string类型的slice
func (b *Bag) Insert(item string) []string { //Insert(itemid int) 的写法与函数一致，(b*Bag) 表示接收器，即 Insert 作用的对象实例
	b.items = append(b.items, item)
	return b.items
}

/*
理解指针类型的接收器
指针类型的接收器由一个结构体的指针组成
由于指针的特性，调用方法时，修改接收器指针的任意成员变量，在方法结束后，修改都是有效的
*/
type Property struct { //定义一个属性结构，拥有一个整型的成员变量
	value int
}

func (p *Property) SetValue(v int) { //接收器为结构体的指针,定义SetValue方法
	p.value = v //设置属性值方法的接收器类型为指针，因此可以修改成员值，即便退出方法，也有效
}

func (p *Property) GetValue() int { //定义获取值的方法
	return p.value
}

/*
理解非指针类型的接收器
当方法作用于非指针接收器时，Go语言会在代码运行时将接收器的值复制一份，在非指针接收器的方法中可以获取接收器的成员值，但修改后无效
// 定义点结构
type Point struct {
    X int
    Y int
}
// 非指针接收器的加方法
func (p Point) Add(other Point) Point {
    // 成员值与参数相加后返回新的结构
    return Point{p.X + other.X, p.Y + other.Y}
}

// 初始化点
p1 := Point{1, 1}    Point{1, 1}返回的不是指针类型; &Point{1, 1}返回的是指针类型
p2 := Point{2, 2}
// 与另外一个点相加
result := p1.Add(p2)
// 输出结果
fmt.Println(result)


在计算机中，小对象由于值复制时的速度较快，所以适合使用非指针接收器，大对象因为复制性能较低，适合使用指针接收器，在接收器和参数间传递时不进行复制，只是传递指针

*/
func main() {
	bag_obj := new(Bag)                //实例化Bag这个结构体
	bag_list := bag_obj.Insert("book") //调用它的Insert方法
	fmt.Printf("Type: %T; Value: %v\n", bag_list, bag_list)

	pro_obj := &Property{}
	pro_obj.SetValue(100)
	fmt.Println(pro_obj.GetValue())
}

/*
Go语言可以对任何类型添加方法，给一种类型添加方法就像给结构体添加方法一样，因为结构体也是一种类型
在Go语言中，使用 type 关键字可以定义出新的自定义类型，之后就可以为自定义类型添加各种方法了
type MyInt int  //MyInt就是自定义int类型

例如:
if v == 0 {}这个判断，可以对v定义一个IsZero方法来判断, 变为 if v.IsZero(){}

type MyInt int     //使用 type MyInt int 将 int 定义为自定义的 MyInt 类型

func (m MyInt) IsZero() bool {  //为 MyInt 类型添加 IsZero() 方法，该方法使用了 (m MyInt) 的非指针接收器，数值类型没有必要使用指针接收器
	return m == 0
}

var d MyInt
d.IsZero()
*/
