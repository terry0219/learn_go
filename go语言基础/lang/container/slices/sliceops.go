package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("%v, len=%d, cap=%d\n", //打印slice的len长度和cap容量; len是当前slice的长度,cap是slice底层数组的容量
		s, len(s), cap(s))
}

func sliceOps() {
	fmt.Println("Creating slice")
	var s []int // Zero value for slice is nil

	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)

	s1 := []int{2, 4, 6, 8} //初始化slice，流程: 先创建好数组，然后在创建一个slice，指向这个数组
	printSlice(s1)

	s2 := make([]int, 16)     //初始化一个长度为16的slice
	s3 := make([]int, 10, 32) //初始化一个长度为10，容量为32的slice, 底层数组的容量是32
	printSlice(s2)
	printSlice(s3)

	fmt.Println("Copying slice")
	copy(s2, s1) //s1是src, s2是dst。 将s1的值拷贝到s2;  s2的值为[2,4,6,8,0,0,0,0,0,0,0,0,0,0,0,0]
	printSlice(s2)

	fmt.Println("Deleting elements from slice") //删除slice中的一个元素，
	s2 = append(s2[:3], s2[4:]...)              //删除s2的中值为8的值，8的下标为3，所以只要append的时候，跳过下标为3的值即可
	//append的第二个参数是可变长参数，通常是这样的append(s2[:3], 0,0,0)，如果想要把s2[4:]的值传出去的话，就用s2[4:]...
	printSlice(s2) //append(slice []Type, elems ...Type) []Type

	fmt.Println("Popping from front")
	front := s2[0] //取s2的第一个元素
	s2 = s2[1:]

	fmt.Println(front)
	printSlice(s2)

	fmt.Println("Popping from back")
	tail := s2[len(s2)-1] // 取s2的最后一个元素
	s2 = s2[:len(s2)-1]

	fmt.Println(tail)
	printSlice(s2)
}
