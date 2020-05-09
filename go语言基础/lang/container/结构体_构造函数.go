//结构体的构造函数
/*
Go语言的类型或结构体没有构造函数的功能，但是我们可以使用结构体初始化的过程来模拟实现构造函数

*/
package main

import "fmt"

type Cat struct {
	Name  string
	Color string
}

func GetCatName(name string) *Cat {
	return &Cat{
		Name: name,
	}
}

type Person struct { //定义Person结构，包含姓名、性别、年龄
	Name string
	Sex  string
	Age  int
}

func GetPersonInfo(name, sex string, age int) *Person { //定义构造Person结构的函数，返回Person指针
	return &Person{ //取地址实例化Person结构
		Name: name,
		Sex:  sex,
		Age:  age,
	}
}

/*
带父子关系的结构体构造和初始化

type Cat struct {
    Color string
    Name  string
}
type BlackCat struct {       //定义 BlackCat 结构，并嵌入了 Cat 结构体，BlackCat 拥有 Cat 的所有成员，实例化后可以自由访问 Cat 的所有成员
    Cat  // 嵌入Cat, 类似于派生
}

// “构造基类”
func NewCat(name string) *Cat {     //NewCat() 函数定义了 Cat 的构造过程，使用名字作为参数，填充 Cat 结构体
    return &Cat{
        Name: name,
    }
}
// “构造子类”
func NewBlackCat(color string) *BlackCat {  //NewBlackCat() 使用 color 作为参数，构造返回 BlackCat 指针
    cat := &BlackCat{}      //实例化 BlackCat 结构，此时 Cat 也同时被实例化
    cat.Color = color      //填充 BlackCat 中嵌入的 Cat 颜色属性，BlackCat 没有任何成员，所有的成员都来自于 Cat。
    return cat
}
*/

func main() {
	var i *Cat

	i = GetCatName("mimi")
	fmt.Printf("value: %+v; type: %T;\n", i, i) //value: &{Name:mimi Color:}; type: *main.Cat;

	p := GetPersonInfo("alex", "nan", 18)
	fmt.Println(p) //&{alex nan 18}
}
