/*
slice是对array的view, 如果修改了slice，相应的也是修改了array;
slice本身是没有数据的，是对底层array的一个view

arr := [10]int{0,1,2,3,4,5,6,7,8,9}
s := arr[:]  s就变成了slice

Reslice就是在slice的基础上在slice
s := arr[2:6]
s = s[:3]
s = s[1:]
s = arr[:]

slice的实现: 包含ptr,len(slice的长度), cap(底层array的容量)
slice可以向后扩展，不能向前扩展，但是向后扩展不可以超越底层数组的cap()
*/
package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	fmt.Println("arr[2:6] =", arr[2:6])
	fmt.Println("arr[:6] =", arr[:6])
	s1 := arr[2:]
	fmt.Println("s1 =", s1)
	s2 := arr[:]
	fmt.Println("s2 =", s2)

	fmt.Println("After updateSlice(s1)")
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	fmt.Println("After updateSlice(s2)")
	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println(arr)

	fmt.Println("Reslice")
	fmt.Println(s2)
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)

	fmt.Println("Extending slice")
	arr[0], arr[2] = 0, 2
	fmt.Println("arr =", arr) //arr是一个数组[0, 1, 2, 3, 4, 5, 6, 7]
	s1 = arr[2:6]             // s1是切片[2,3,4,5]
	s2 = s1[3:5]              // [s1[3], s1[4]]    s2是s1的切片[5,6]
	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n",
		s1, len(s1), cap(s1))
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n",
		s2, len(s2), cap(s2))

	s3 := append(s2, 10)     // 向变量为s2的slice类型添加元素 [5,6,10]
	s4 := append(s3, 11)     // [5,6,10,11]
	s5 := append(s4, 12)     // [5,6,10,11,12]
	fmt.Println("s3, s4, s5 =", s3, s4, s5) 
	// s4 and s5 no longer view arr.
	fmt.Println("arr =", arr) // arr的值为[0, 1, 2, 3, 4, 5, 6, 10],最后一个值变更了，因为引用它的slice修改的; arr还是一个长度为8的数组
							  //s3 s4的值因为超过了arr的长度，所以s3 s4指向的是一个新的array,是由系统分配的
							// 添加元素时，如果超越cap,系统会重新分配一个更大的底层数组; 如果旧的arr没在用的话，系统会垃圾回收
							// 由于值传递的原因，必须接受append返回值  s = append(s,val)
	// Uncomment to run sliceOps demo.
	// If we see undefined: sliceOps
	// please try go run slices.go sliceops.go
	fmt.Println("Uncomment to see sliceOps demo")
	// sliceOps()
}
