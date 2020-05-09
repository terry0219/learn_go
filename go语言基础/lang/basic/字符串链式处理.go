//字符串的链式处理
package main

import (
	"fmt"
	"strings"
)

//字符串处理函数（ProccessString）需要外部提供数据源，一个字符串切片（list[]string），另外还要提供一个链式处理函数的切片（chainfunc []func(s string) string）
func ProcessString(list []string, chainfunc []func(s string) string) []string {	
	for k,v := range list {
		// 将当前字符串保存到 result 变量中，作为第一个处理函数的参数
		result := v

		// 输入一个字符串进行处理, 返回数据作为下一个处理链的输入(这块当时没写出来，没有想到定义个result变量，然后通过函数一直对result变量进行修改)
		for _,f := range chainfunc{
			result = f(result) //result 变量即是每个处理函数的输入变量，处理后的变量又会重新保存到 result 变量中。
		}
		// 将结果放回切片
		list[k] = result
		
	}
	//fmt.Println(list)	
	return list
}

//定义一个去除字符串空格的函数
func removeSpace(s string) string {
	str := strings.Replace(s, " ", "", -1)
	return str
}

func main() {
	// 待处理的字符串列表
	init := []string{
		"aa aa",
		"bb b",
		"cc ccc",
	}

	// 处理函数链
	chainfunc := []func(s string) string{
		removeSpace,
		strings.ToUpper,
		
	}

	t := ProcessString(init, chainfunc)
	fmt.Println(t)
	
}