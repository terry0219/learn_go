/*
golang中string底层是通过byte数组实现的。中文字符在unicode下占2个字节，在utf-8编码下占3个字节，而golang默认编码正好是utf-8。
Go语言中byte和rune实质上就是uint8和int32类型。byte用来强调数据是raw data，而不是数字；而rune用来表示Unicode的code point

byte 等同于uint8，常用来处理ascii字符
rune 等同于int32,常用来处理unicode或utf-8字符

*/
package main

import (
	"fmt"
)

func main() {
	s := "Yes我爱慕课网!"
	/*
		for k, v := range s {
			fmt.Println(k, v)
		}
		返回:
			0 89         这里返回的89 101 115这些是十进制的ascii码
			1 101
			2 115
			3 25105
			6 29233
			9 24917
			12 35838
			15 32593
			18 33
	*/

	/*
		for k, v := range []byte(s) {
			fmt.Println(k, v)
		}
		返回:
			0 89      
			1 101
			2 115
			3 230
			4 136
			5 145
			6 231
			7 136
			8 177
			9 230
			10 133
			11 149
			12 232
			13 175
			14 190
			15 231
			16 189
			17 145
			18 33
	*/

	/*
		for k, v := range []rune(s) {
			fmt.Println(k, v)
		}

		返回:
		0 89
		1 101
		2 115
		3 25105
		4 29233
		5 24917
		6 35838
		7 32593
		8 33
	*/

	fmt.Println(len(s)) //19
	//[]byte(s) 查看s这个字符原始的字节，看是如何存的
	fmt.Printf("%s\n", []byte(s)) // %s打印具体的内容  Yes我爱慕课网!
	//fmt.Printf("%X\n", []byte(s)) // %X打印字节具体的数字，也就是ascii码

	//%x 输出使用 base-16 编码的字符串，每个字节使用 2 个字符表示; %c 相应Unicode码点所表示的字符
	//%x 十六进制ascii表示，字母形式为小写 a-f; %X 十六进制ascii表示，字母形式为大写 A-F
	//(0 uint8 59) (1 uint8 65) (2 uint8 73)  像59 65 73这种就是ascii码 十六进制的 59代表Y 65代表e 73代表s
	for i, b := range []byte(s) {
		fmt.Printf("(%d %T %X) ", i, b, b) // (0 uint8 59) (1 uint8 65) (2 uint8 73) (3 uint8 E6) (4 uint8 88) (5 uint8 91) (6 uint8 E7) (7 uint8 88) (8 uint8 B1) (9 uint8 E6) (10 uint8 85) (11 uint8 95) (12 uint8 E8) (13 uint8 AF) (14 uint8 BE) (15 uint8 E7) (16 uint8 BD) (17 uint8 91) (18 uint8 21)
	}

	fmt.Println()

	for i2, b2 := range s {
		fmt.Printf("(%d %T %X,%c) ", i2, b2, b2, b2) //(0 int32 59,Y) (1 int32 65,e) (2 int32 73,s) (3 int32 6211,我) (6 int32 7231,爱) (9 int32 6155,慕) (12 int32 8BFE,课) (15 int32 7F51,网) (18 int32 21,!)
	}

	fmt.Println()

	for i3, b3 := range []rune(s) { //%X返回的就是ascii码
		fmt.Printf("(%d %T %X %c) ", i3, b3, b3, b3) // (0 int32 59 Y) (1 int32 65 e) (2 int32 73 s) (3 int32 6211 我) (4 int32 7231 爱) (5 int32 6155 慕) (6 int32 8BFE 课) (7 int32 7F51 网) (8 int32 21 !)
	}

	s2 := "hello world!"
	fmt.Println(len(s2)) //12 英文字符占一个字节
	//fmt.Println(s[0], s[1]) //89 101 返回对应的字节而不是字符
	bytes := []byte(s2)
	fmt.Printf("%X\n", s2)                            //68 65 6C 6C 6F 20 77 6F 72 6C 64 21
	fmt.Printf("type: %T; value: %v\n", bytes, bytes) //type: []uint8; value: [104 101 108 108 111 32 119 111 114 108 100 33]

	//字符串截取和拼接
	fmt.Println(s2[:5])         //hello
	fmt.Println(s2[6:])         //world!
	fmt.Println("hi " + s2[6:]) //hi world!

	fmt.Printf("type: %T", s) // string
}
