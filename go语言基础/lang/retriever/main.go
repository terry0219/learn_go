/*

写代码小技巧: 先把大的框架写好，在不断完善;  先把抽象interface的内容写好(type i interface{Get()})，在写func Test(t i){t.Get()}将接口类型传给形参, 然后在实现struct中实现具体的Get方法，
最后实例化这个struct，赋給func

接口: 定义好抽象的方法
函数: 传一个形参，类型为接口
结构体: 实例化结构体&具体实现这个方法的逻辑

代码顺序:
1. 写好接口interface，定义好方法
2. 写好函数func，将interface传入进来
3. 定义好结构体struct, 并为这个结构体实现具体的方法
4. 实例化struct，传入初始值
5. 调用第2步的函数func，传入实例化后的对象
6. 执行完成

接口的组合在系统库中经常会用到

比如io.ReadWriteCloser就是一个接口组合
type ReadWriteCloser interface {
	Reader
	Writer
	Closer
}

Reader Writer Closer分别都是接口类型，然后在组合一起
func test(r ReadWriteCloser) {
	//变量r就拥有了这些接口的方法
}


type i interface{
	Get()
	Post()
}
接口类型的变量一般都是赋給一个func用的

-------------------------------------------------------------------------------------

定义接口类型，以及它拥有的方法
type testInterface interface {
	Get(string) string
}

定义函数，形参是接口
func testFunc(t testInterface) string {
	s := t.Get("test string")
	return s
}

定义结构体
type testStruct struct {
	x,y string
}

为结构体实现Get方法
func (t testStruct) Get(string) string {
	return t.x
}

初始化结构体
t := testStruct{
	"x":"golang",
	"y":"python",
}

将初始化的结构体赋給i接口
var i testInterface
i = t

将接口i变量传入函数并调用
s := testFunc(i)



*/
package main

import (
	"fmt"

	"time"

	"imooc.com/ccmouse/learngo/lang/retriever/mock"
	"imooc.com/ccmouse/learngo/lang/retriever/real"
)

type Retriever interface { //定义Retriever是一个接口类型，它有一个Get方法
	Get(url string) string
}

type Poster interface { //在定义个Poster接口，有一个Post方法，有两个形参，一个url字符串,另外一个map[string]string字典
	Post(url string,
		form map[string]string) string
}

const url = "http://www.imooc.com"

func download(r Retriever) string { //定义一个download函数，接收一个接口类型的形参
	return r.Get(url) //这个接口定义了Get方法
}

func post(poster Poster) { //定义post函数，形参是Poster类型的接口
	poster.Post(url, //调用Poster接口里的Post方法
		map[string]string{
			"name":   "ccmouse",
			"course": "golang",
		})
}

//接口的组合 将Retriever和Poster两个接口组合起来
type RetrieverPoster interface {
	Retriever
	Poster
}

//RestrieverPoster是一个组合接口，它既有Get方法还有Post方法
func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{ //调用Post方法，传两个参数，一个Url，一个map
		"contents": "another faked imooc.com",
	})
	return s.Get(url)
}

func main() {
	var r Retriever //定义变量r为retriever类型

	mockRetriever := mock.Retriever{ //实例化结构体
		Contents: "this is a fake imooc.com"}
	//只要实现接口里的方法就可以了
	// mock.Retriever和real.Retriever都实现了get方法，可以赋給r这个接口
	// r = mock.Retriever{"this is fake imooc.com"} 将实例化后的对象赋給r这个接口
	// r = real.Retriever{}
	// fmt.Printf("%T %v\n",r,r)  // {real.Retriever Mozilla/5.0 0s}   看下r肚子里有什么东西
	// fmt.Println(download(r))

	r = &mockRetriever
	inspect(r)

	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	inspect(r)

	//通过.(类型的名字)，可以取得这个interface肚子里面具体的类型
	realRetriever := r.(*real.Retriever)
	fmt.Println(realRetriever.TimeOut)

	// Type assertion
	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("r is not a mock retriever")
	}

	fmt.Println(
		"Try a session with mockRetriever")
	fmt.Println(session(&mockRetriever)) //调用session函数，传&mockRetriever
}

//
func inspect(r Retriever) {
	fmt.Println("Inspecting", r)
	fmt.Printf(" > Type:%T Value:%v\n", r, r)
	fmt.Print(" > Type switch: ")
	switch v := r.(type) { //r.(type)查看r的肚子里是什么类型的
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
	fmt.Println()
}

/*
接口变量里有什么？ 实现者的类型和实现者的值

接口变量自带指针
接口变量同样采用值传递，几乎不需要使用接口的指针

接口的值类型:

表示任何类型: interface{}

type Queue []interface{}   代表Queue是slice类型，这个slice里面可以有int string等任何类型

func (q *Queue) Push(v interface{}) {  // Push方法接收一个任意类型的形参 q.Push("abc") 也可以q.Push(1)

}

func (q *Queue) Pop() interface{} {   返回任意类型

}

如果我们想在上层限制只能返回int类型的话
func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head.(int)    因为head是interface类型，因为我们限制了返回int类型，所以需要转换一下， head.(int)
}

func (q *Queue) Push(v interface{}) {
	*q = append(*q, v.(int))    通过v.(int) 强制将v的interface类型转换为int
}


常用的系统接口
fmt/print.go
type Stringer interface {
	String() string
}

io/io.go
type Reader interface {
	Read(p []byte) (n int,err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}



*/




