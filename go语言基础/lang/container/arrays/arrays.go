package main

import "fmt"

/*
数组的几种定义方法

var arr1 [3]int
arr2 := [3]int{1,2,3}
arr3 := [...]int{1,2,3,4,5,6}

var gade [4][5]int     //4行5列 多维数组 [ [0 0 0 0 0] [0 0 0 0 0] [0 0 0 0 0] [0 0 0 0 0]] 

数组是值类型，就是拷贝; 当定义好一个变量a后，传给函数b，变量a实际上拷贝了新地址传给b的，所以函数b无论怎么改a，只能在函数内生效。
如果想改变原来数组的内容，就需要给函数传一个指针地址进去，调用函数时，传数组的地址
func printArray(arr *[5]int){}
a := [3]int{1,2,3}
printArray(&a)


[10]int 和 [20]int 是不同的类型
调用函数 func f(arr [10]int) 会拷贝数组
*/

func printArray(arr [5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}
	var grid [4][5]int

	fmt.Println("array definitions:")
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	fmt.Println("printArray(arr1)")
	printArray(arr1)

	fmt.Println("printArray(arr3)")
	printArray(arr3)

	fmt.Println("arr1 and arr3")
	fmt.Println(arr1, arr3)
}
