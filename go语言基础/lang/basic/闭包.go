package main

/*
 内函数对外函数的变量的修改，是对变量的引用。共享一个在堆上的变量。 变量被引用后，它所在的函数结束，这变量也不会马上被销毁。
 相当于变相延长了函数的生命周期

 */
func AntherExFunc(n int) func() {
    n++
    return func() {
        fmt.Println(n)
    }
}

func ExFunc(n int) func() {
    return func() {
        n++
        fmt.Println(n)
    }
}

func main() {
    myAnotherFunc:=AntherExFunc(20)
    fmt.Println(myAnotherFunc)  //0x48e3d0  在这儿已经定义了n=20 ，然后执行++ 操作，所以是21 。
    myAnotherFunc()     //21 后面对闭包的调用，没有对n执行加一操作，所以一直是21
    myAnotherFunc()     //21

    myFunc:=ExFunc(10)
    fmt.Println(myFunc)  //0x48e340   这儿定义了n 为10
    myFunc()       //11  后面对闭包的调用，每次都对n进行加1操作。
    myFunc()       //12

}
