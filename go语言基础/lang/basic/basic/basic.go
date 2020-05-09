package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	aa = 3      //如果没有定义变量类型，编译器会自动推算; var aa int = 3 这种就是明确定义了
	ss = "kkk"
	bb = true
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s)
}

func euler() {
	fmt.Printf("%.3f\n",
		cmplx.Exp(1i*math.Pi)+1)
}

func triangle() {
	var a, b int = 3, 4
	fmt.Println(calcTriangle(a, b))
}

func calcTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

//常量
func consts() {
	const (
		filename = "abc.txt"
		a, b     = 3, 4 //如果常量没有定义类型，那么它是不确定的，这里a b既可以当int，也可以当float; const a int = 3 这种就是确定类型的
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

//自增枚举
func enums() {
	const (           //使用const定义枚举, iota代表是可以自增值的(cpp=0 python=2 golang=3 javascript=4)
		cpp = iota
		_
		python
		golang
		javascript
	)

	const (
		b = 1 << (10 * iota)  //iota也可以参与运算
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(cpp, javascript, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("Hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, ss, bb)

	euler()
	triangle()
	consts()
	enums()
}

/*
要点回顾:

变量类型写在变量名之后
编译器可推测变量类型
没有char,只有rune
原生支持复数类型

*/