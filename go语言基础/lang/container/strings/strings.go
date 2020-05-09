/*
go是如何处理国际化多语言的，支持的关键就是要理解rune这个类型

rune 相当与go的char

*/

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes我爱慕课网!" // UTF-8
	fmt.Println(s)   // 19
	//fmt.Printf("%s\n", []byte(s)) //Yes我爱慕课网!

	for _, b := range []byte(s) {
		fmt.Printf("%X ", b) // 16进制 acsii码  59 65 73 E6 88 91 .....  59代表Y 65代表e 73代表s E6 88 91代表我； 一个中文字符占3个字节(这就是UTF8编码)；英文字符占1个字节
	}
	fmt.Println()

	for i, ch := range s { // ch is a rune  ch是int32类型的，也就是rune类型; (3 6211)代表"我",占用了2个字节; range s对utf8字符进行解码unicode
		fmt.Printf("(%d %X) ", i, ch) // (0 59) (1 65) (2 73) (3 6211) (6 7231) (9 6155)
		//E6 88 91 进行unicode解码后为 6211
	}
	fmt.Println()

	fmt.Println("Rune count:",
		utf8.RuneCountInString(s)) //9

	bytes := []byte(s) //拿到字节
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch) // Y e s 我 爱 慕 课 网 ！
	}
	fmt.Println()

	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c) ", i, ch)
	}
	fmt.Println()
}
