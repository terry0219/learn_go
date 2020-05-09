/*
panic和recover

panic不能经常用

1. 停止当前函数执行
2. 一直向上返回, 执行每一层的defer
3. 如果没有遇到recover, 程序退出

如果程序抛出异常，没有人来接的话，程序就挂掉了， panic也是这样，如果没有recover来接的话，程序也是挂掉了


recover
1. 仅在defer调用中使用
2. 获取panic的值 panic(err)  就是获取err的值
3. 如果无法处理，可重新panic

defer后面一个匿名函数, ()表示调用这个函数
func tryRecover() {

	defer func(){
	r := recover()   因为recover返回的值一个interface{}任意类型的
	if err,ok := r.(error); ok{   所以要判断下r的类型是否为error，r.(error)
		fmt.Println("Error occurred: ", err)
	}else{
		panic(r)   如果r 不是error类型，那么就panic
	}
}()

	//模拟程序抛异常
	panic(errors.New("this is an error"))

	//写一段程序抛错, 让recover接住
	a := 0
	b := 5/a
	fmt.Pirntln(b)

	//因为123不是error类型的，所以recover会走else分支
	panic(123)
}



*/
package main

import (
	"fmt"
)

func tryRecover() {
	defer func() {
		r := recover()
		if r == nil {
			fmt.Println("Nothing to recover. " +
				"Please try uncomment errors " +
				"below.")
			return
		}
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred:", err)
		} else {
			panic(fmt.Sprintf(
				"I don't know what to do: %v", r))
		}
	}()

	// Uncomment each block to see different panic
	// scenarios.
	// Normal error
	//panic(errors.New("this is an error"))

	// Division by zero
	//b := 0
	//a := 5 / b
	//fmt.Println(a)

	// Causes re-panic
	//panic(123)
}

func main() {
	tryRecover()
}
