/*
go语言闭包的应用

例子1: 斐波那契数列

例子2: 为函数实现接口   

	函数也是可以实现接口的
	可以看出斐波那契数列中调用的f(), 像一个生成器一样,每次调用一次就会生成下一个斐波那契数，如下:
	f := fibonacci()
	f()
	f()
	f()

	这个跟读文件有点像，可以包装成io.Reader接口,就可以当文件一样来读斐波那契数
	type Reader interface {
		Read(p []byte) (n int,err error)
	}

	像这样的,原来的用途是用来打印文件的
	func printFileContents(reader io.Reader) {
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}

	如果生成器实现了Reader接口,那么就可以打印斐波那契数了



*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"imooc.com/ccmouse/learngo/lang/functional/fib"
)

type intGen func() int  //定义intGen为函数类型，这个函数没有参数，返回int

func (g intGen) Read(         //为intGen函数实现Read方法
	p []byte) (n int, err error) {
	next := g()        //调用一次函数
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)   //把next转成字符串类型

	// TODO: incorrect if p is too small!
	return strings.NewReader(s).Read(p)  //把s字符串写到p这个byte类型的slice
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	var f intGen = fib.Fibonacci()
	printFileContents(f)        //f这个函数实现了Read方法，所以可以赋給printFileContents(reader io.Reader)这个函数，io.Reader是一个接口，它有Read方法; 
								//只要传进去的变量实现了Read方法就行
}
