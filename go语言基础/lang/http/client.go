/*
1. 使用http客户端发送请求
2. 使用http.Client控制请求头部等
3. 使用httputil简化工作

基本的请求
func main() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil{
		panic(err)
	}
	defer resp.Body.Close()

	s, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n",s) //打印请求头部以及内容
}

接下来对http请求进行一些控制, 自定义header
func main() {
	request, err := http.NewRequest(http.MethodGet,"http://www.imooc.com",nil)
	//添加request header  请求手机版的慕课网
	request.Header.Add("User-Agent","Mozilla/5.0 (iPhone; CPU iPhone OS 10_3 like Mac OS X) AppleWebKit/602.1.50 (KHTML, like Gecko) CriOS/56.0.2924.75 Mobile/14E5239e Safari/602.1")

	resp, err := http.DefaultClient.Do(request)
	if err != nil{
		panic(err)
	}
	defer resp.Body.Close()

	s, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n",s) //打印请求头部以及内容
}

我们还可以自己去生成http.client
func main() {
	request, err := http.NewRequest(http.MethodGet,"http://www.imooc.com",nil)
	//添加request header  请求手机版的慕课网
	request.Header.Add("User-Agent","Mozilla/5.0 (iPhone; CPU iPhone OS 10_3 like Mac OS X) AppleWebKit/602.1.50 (KHTML, like Gecko) CriOS/56.0.2924.75 Mobile/14E5239e Safari/602.1")

	//http.Client是一个struct
	//type Client struct {
	//	Transport RoundTripper   代理服务器之类的
	//	CheckRedirect func(req *Request, via []*Request) error  如果有重定向会从这个地方过
	//	Jar CookieJar   cookie，在做爬虫项目时模拟用户登录会用到
	//}
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			//via是一个slice，重定向可以有好几次，via里存了重定向的路径; 每一次的目标放在req里
			fmt.Println("Redirect: ", req) 把重定向的地址打印出来  Redirect: &{GET https://www.baidu.com}
			return nil

		}
	}      
	resp, err := client.Do(request)
	if err != nil{
		panic(err)
	}
	defer resp.Body.Close()

	s, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n",s) //打印请求头部以及内容
}


---------------------------------------------------------------------------------

http服务器的性能分析

import _ "net/http/pprof"
访问 /debug/pprof
使用go tool pprof分析性能

在http server中加入 import _ "net/http/pprof" 就可以分析了；在通过访问http://x.x.x.x:8888/debug/pprof就可以访问

还有一个功能，在命令行执行

# go tool pprof http://localhost:8888/debug/pprof/profile  看30秒的CPU负载
等30秒后出现
(pprof) web   输入web，就可以查看个函数的调用时间


---------------------------------------------------------------------------------------
其他标准库

bufio   打开文件或者读文件，使用bufio加一层缓冲，一次性可以写很多数据到缓冲区，最后在刷新到磁盘，类似异步非阻塞
log
encoding/json    结构体可以按照格式encoding
regexp  正则表达式
time    time.Sleep(100 * time.Second)  等待100秒
strings/math/rand 

查看标准库的用法
1. godoc -http :8888   访问后，去查询Packages
2. http://docs.studygolang.com/pkg/   这里也可以查看package的用法


-------------------------------------------------------------------









使用说明可以在pprof.go文件中查看
*/
package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	request, err := http.NewRequest(
		http.MethodGet,
		"http://www.imooc.com", nil)
	request.Header.Add("User-Agent",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 10_3 like Mac OS X) AppleWebKit/602.1.50 (KHTML, like Gecko) CriOS/56.0.2924.75 Mobile/14E5239e Safari/602.1")

	client := http.Client{
		CheckRedirect: func(
			req *http.Request,
			via []*http.Request) error {
			fmt.Println("Redirect:", req)
			return nil
		},
	}
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	s, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", s)
}
