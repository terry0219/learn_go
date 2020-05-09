/*
使用正则表达式，从text字符串中找出terry@163.com

const text = "My email is terry@163.com"

func main() {
	re := regexp.MustCompile("terry@163.com")  定义正则表达式
	match := re.FindString(text)  从text字符串中找出符合正则表达式的字符
	fmt.Println(match)  输出terry@163.com

	regexp.MustCompile(".*@163.com")  .*代表匹配0个或多个  .+代表匹配1个或多个   
	MustCompile(".+@.+\\..+")    \\.是匹配.点的
	MustCompile(`.+@.+\..+`)  建议用这种方式   一个点.代表任何字符
	MustCompile(`[a-zA-Z0-9]+@`)  匹配@前面是a-zA-Z0-9的字符，一个或多个; @前面必须是字母或者数字

	[a-zA-Z0-9.]+   匹配一个或多个字母或者数字或者.
	`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`  正则表达式中加上()就可以单独提取出来了
}

*/
package main

import (
	"fmt"
	"regexp"
)

const text = `
my email is ccmouse@gmail.com@abc.com
email1 is abc@def.org
email2 is    kkk@qq.com
email3 is ddd@abc.com.cn
`

func main() {
	re := regexp.MustCompile(
		`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	match := re.FindAllStringSubmatch(text, -1) //re.FindAllStringSubmatch返回的是一个[][]string二维数组
	//match := re.FindAllString(text, -1) 从text中查询所有符合正则的字符 返回[abc@def.org kkk@qq.com ddd@abc.com]
	for _, m := range match {
		fmt.Println(m)
	}
	/*
	输出:
	[ccmouse@gmail.com ccmouse gmail com]  ccmouse就是第一个括号的内容， gmail是第二个括号  com是第三个括号
	[abc@def.org abc def org]
	*/
}
