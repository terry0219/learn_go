/*
斐波那契数列, 返回值是前面两次的值相加起来的

编写代码顺序

1. 先定义函数结构

func fibonacci() func() int {

}

2. 定义main

func main() {
	f := fibonacci()
	f() //1
	f() //1
	f() //2
	f() //3
	f() //5
	f() //8

}

3. 最后在填充fibonacci函数的具体内容


*/
package fib

// 1, 1, 2, 3, 5, 8, 13, ...
func Fibonacci() func() int { //Fibonacci函数返回一个无参数的函数
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
