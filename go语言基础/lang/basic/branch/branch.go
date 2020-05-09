package main

import (
	"fmt"
	"io/ioutil"
)

func readfile() {
	const filename = 'abc.txt'
	contents, err := ioutil.ReadFile(filename) //读取文件内容, contents返回的是一个[]byte数组
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n",contents)
	}
}

/*
if简化写法:

if contents, err := ioutil.ReadFile(filename); err != nil {
	fmt.Println(err)
} else {
	fmt.Printf("%s\n",contents)
}

if的条件里可以赋值
if的条件里赋值的变量作用域就在这个if语句里
*/

/*
switch后面可以没有表达式
panic使程序中断执行
*/
func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf(
			"Wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

func main() {
	// If "abc.txt" is not found,
	// please check what current directory is,
	// and change filename accordingly.
	const filename = "abc.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	fmt.Println(
		grade(0),
		grade(59),
		grade(60),
		grade(82),
		grade(99),
		grade(100),
		// Uncomment to see it panics.
		// grade(-3),
	)
}
