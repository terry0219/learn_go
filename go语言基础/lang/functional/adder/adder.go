/*
函数式编程

go语言对函数式编程的支持主要体现在闭包上面

函数是一等公民: 参数，变量，返回值都可以是函数
高阶函数: 函数的参数可以是函数
闭包

写代码一个简单思路

1. 先定义好函数
func adder() func(int) int {

}

2. 在写如何调用
func main() {
	for i:=0; i<10; i++ {
		a(i)
	}
}

3. 最后在填充func
func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}


python中的闭包

def add():
	sum = 0
	def f(v):
		nonlocal sum   #声明sum为自由变量
		sum += v
		return sum
	return f


*/
package main

import "fmt"

//闭包
func adder() func(int) int {
	sum := 0
	return func(v int) int { //v是一个局部变量
		sum += v //sum不是函数体内的变量，是函数外部的变量，成为自由变量
		return sum
	}
}

type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}

func main() {
	// a := adder() is trivial and also works.
	a := adder2(0)
	for i := 0; i < 10; i++ {
		var s int
		s, a = a(i)
		fmt.Printf("0 + 1 + ... + %d = %d\n",
			i, s)
	}
}
