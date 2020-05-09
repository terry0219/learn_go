/*
空接口是指没有定义任何接口方法的接口。
没有定义任何接口方法，意味着Go中的任意对象都可以实现空接口(因为没方法需要实现)，任意对象都可以保存到空接口实例变量中。

更常见的，会直接使用interface{}作为一种类型，表示空接口。例如：

// 声明一个空接口实例
var i interface{}
再比如函数使用空接口类型参数：

func myfunc(i interface{})
在Go中很多地方都使用空接口类型的参数，用的最多的fmt中的Print类方法：

$ go doc fmt Println
func Println(a ...interface{}) (n int, err error)


某个struct中，如果有一个字段想存储任意类型的数据，就可以将这个字段的类型设置为空接口：
type my_struct struct {
    anything interface{}
    anythings []interface{}
}

---------------------------------------------------------------------------------

*/
/*
package main

import "fmt"

func main() {
	//var s []interface{}
	s := make([]interface{}, 5) //创建一个空接口的slice
	s[0] = 1
	s[1] = "golang"
	s[2] = true

	fmt.Println(s) //[1 golang true <nil> <nil>]
	s = append(s, 2, 3)
	fmt.Println(s) //[1 golang true <nil> <nil> 2 3]
	//通过空接口类型，Go也能像其它动态语言一样，在数据结构中存储任意类型的数据

	fmt.Printf("%T %+v\n", s, s) //[]interface {} [1 golang true <nil> <nil> 2 3]
	// ------------------------------------------------------------------------

	//接口转回成具体类型
	var a int = 1
	var ins interface{}
	// 接口实例ins中保存的是int类型
	ins = a
	fmt.Printf("%T %v\n", ins, ins) //int 1

	// 接口转回int类型的实例
	x := ins.(int)
	fmt.Printf("%T %v\n", x, x)                   //int 1
	fmt.Println("a addr: ", &a, "\nx addr: ", &x) //a addr:  0xc0000100c8   x addr:  0xc0000100f8

	/*
				如果ins不是*int类型，会panic
				y := ins.(*int)
				fmt.Println(y)  //panic: interface conversion: interface {} is int, not *int

				注意，接口实例转回时，接口实例中存放的是什么类型，才能转换成什么类型。同类型的值类型实例和指针类型实例不能互转，不同类型更不能互转。
				在不能转换时，Golang将直接以Panic的方式终止程序。但可以处理转换失败时的panic，这时需要类型断言，也即类型检测。
				/*

				// ------------------------------------------------------------------------------------
				//接口类型断言
				/*
				类型探测的方式和类型转换的方式都是ins.(Type)和ins.(*Type)。
				当处于单个返回值上下文时，做的是类型转换，当处于两个返回值的上下文时，做的是类型探测。
				类型探测的第一个返回值是类型转换之后的类型实例，第二个返回值是布尔型的ok返回值。

				x := ins.(Type)  单个返回值的就是做类型转换
				x, ok := ins.(Type) 多个返回值的就是类型断言

				// 如果ins保存的是值类型的Type，则输出
					if t, ok := ins.(Type); ok {
						fmt.Printf("%T\n", v)
					}

					// 如果ins保存的是指针类型的*Type，则输出
					if t, ok := ins.(*Type); ok {
						fmt.Printf("%T\n", v)
		}


*/

//}
//*/

package main

import "fmt"

type testInterface interface {
	Get(s string) string
}

func testFunc1(t testInterface) {
	fmt.Println(t.(testStruct))
	val, ok := t.(testStruct) //接口类型断言,判断t的类型是否为testStruct
	fmt.Println(val, ok)      //{aaa bbb} true

	t.Get("a")
}

type testStruct struct {
	x, y string
}

func (t testStruct) Get(s string) string {
	return t.x
}

func main() {
	var i testInterface

	a := testStruct{"aaa", "bbb"}

	i = a //将结构体赋給接口i
	testFunc1(i)
}

/*
// Shaper 接口类型
type Shaper interface {
    Area() float64
}

// Square struct类型
type Square struct {
    length float64
}

// Square类型实现Shaper中的方法Area()
func (s Square) Area() float64 {
    return s.length * s.length
}

func main() {
    var ins1, ins2 Shaper

    // 指针类型的实例
    s1 := new(Square)
    s1.length = 3.0
    ins1 = s1                //因为s1是指针类型的，所以ins1是指针类型的实例
    if v, ok := ins1.(*Square); ok {
        fmt.Printf("ins1: %T\n", v)
    }

    // 值类型的实例
    s2 := Square{4.0}
    ins2 = s2                 //因为s2是值类型的，所以ins2也是值类型的实例
    if v, ok := ins2.(Square); ok {
        fmt.Printf("ins2: %T\n", v)
    }
}

输出:
ins1: *main.Square
ins2: main.Square

-------------------------------------------------------------------------

type Switch结构

直接用if v,ok := ins.(Type);ok {}的方式做类型探测在探测类型数量多时不是很方便，需要重复写if结构。

Golang提供了switch...case结构用于做多种类型的探测，所以这种结构也称为type-switch。
这是比较方便的语法，比如可以判断某接口如果是A类型，就执行A类型里的特有方法，如果是B类型，就执行B类型里的特有方法。

用法如下：

switch v := ins.(type) {
case *Square:
    fmt.Printf("Type Square %T\n", v)
case *Circle:
    fmt.Printf("Type Circle %T\n", v)
case nil:
    fmt.Println("nil value: nothing to check?")
default:
    fmt.Printf("Unexpected type %T", v)
}
其中ins.(type)中的小写type是固定的词语。

例子:

package main

import (
    "fmt"
)

// Shaper 接口类型
type Shaper interface {
    Area() float64
}

// Circle struct类型
type Circle struct {
    radius float64
}

// Circle类型实现Shaper中的方法Area()
func (c *Circle) Area() float64 {
    return 3.14 * c.radius * c.radius
}

// Square struct类型
type Square struct {
    length float64
}

// Square类型实现Shaper中的方法Area()
func (s Square) Area() float64 {
    return s.length * s.length
}

func main() {
    s1 := &Square{3.3}
    whichType(s1)

    s2 := Square{3.4}
    whichType(s2)

    c1 := new(Circle)
    c1.radius = 2.3
    whichType(c1)
}

func whichType(n Shaper) {
    switch v := n.(type) {
    case *Square:
        fmt.Printf("Type Square %T\n", v)
    case Square:
        fmt.Printf("Type Square %T\n", v)
    case *Circle:
        fmt.Printf("Type Circle %T\n", v)
    case nil:
        fmt.Println("nil value: nothing to check?")
    default:
        fmt.Printf("Unexpected type %T", v)
    }
}

上面的type-switch中，之所以没有加上case Circle，是因为Circle只实现了指针类型的receiver，根据Method Set对接口的实现规则，只有指针类型的Circle示例才算是实现了接口Shaper，所以将值类型的示例case Circle放进type-switch是错误的

*/
