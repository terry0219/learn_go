/*
资源管理与出错处理

资源管理:
		比如打开文件后，要关闭; 连接数据库后要释放；

当程序执行过程中遇到报错，可能会导致打开的连接没有及时释放； 即使代码中写了关闭连接的语句，但是也要看程序运行时中间有没有中断，导致连接不释放

defer调用
	1. 确保调用在函数结束时发生

	func tryDefer() {
		defer fmt.Println(1)
		defer fmt.Println(2)
		fmt.Println(3)
	}
	输出: 3 2 1
	defer里面相当于是一个栈，是先进后出的，所以是 3 2 1

	加了defer后，不管程序运行时遇到return还是panic, defer后面的语句都会被执行

	2. defer列表为先进后出


何时使用defer
1. Open/Close
2. Lock/Unlock
3. PrintHeader/PrintFooter



错误处理

if err != nil {
	fmt.Println("error:", err.Error())
	return
}


*/
package main

import (
	"fmt"
	"os"

	"bufio"

	"imooc.com/ccmouse/learngo/lang/functional/fib"
)

//输出 30 29 28 27 26 25 24 23 22 21 20 ... 0  倒过来输出的
//defer相当于是一个栈，先进后出
func tryDefer() {
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			// Uncomment panic to see
			// how it works with defer
			// panic("printed too many")
		}
	}
}

func writeFile(filename string) {
	//file, err := os.Create(filename) 也是打开文件
	file, err := os.OpenFile(filename,
		os.O_EXCL|os.O_CREATE|os.O_WRONLY, 0666)

	// err = errors.New("this is custom error")  自定义错误
	/*
	type error interface {
		Error() string
	}

	error是一个接口，也可以通过自定义error实现Error方法，返回string就可以
	*/

	
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok { //通过err.(*PathError)来获取错误; 如果err变量不是*PathError类型，就执行panic
			panic(err)
		} else {
			fmt.Printf("%s, %s, %s\n",
				pathError.Op,         //open
				pathError.Path,       //fib.txt
				pathError.Err)        //file exists
		}
		return
	}
	defer file.Close()	//关闭打开的文件

	//因为直接写file会比较慢，所以用到bufio。这个writer是带有buffer的
	writer := bufio.NewWriter(file)
	defer writer.Flush()   //因为是Buffer，所以需要Flush刷新到磁盘上

	f := fib.Fibonacci()      //斐波那契数  fib是package名, Fibonacci是fib包下的函数; 返回的f是一个函数
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())  //将f()生成的内容，写入到writer的buffer中
	}
}

func main() {
	tryDefer()
	writeFile("fib.txt")
}
