/*
如何实现统一的错误处理逻辑

if err != nil {
	panic(err)
} else{
	...
}

把这段代码提出来，封装起来

实现一个打印文件内容的webserver，比如访问http://x.x.x.x:8888/list/fib.txt 就打印出fib.txt的内容

func main() {
	//第一个参数是固定写好的访问地址 比如/list/fib.txt,第二个参数是一个函数
	http.HandleFunc("/list/", func(writer http.ResponseWriter, request *http.Request) {
			path := request.URL.Path[len("/list/"):]   //request.URL.Path是全路径/list/fix.txt; 
													   //request.URL.Path[len("/list/"):] 是取出后面的文件名fib.txt
			file, err := os.Open(path)
			if err != nil {
				http.Error(writer,             http.ResponseWriter
						   err.Error(),        内部的出错信息(这样不好，会返回给用户看了) 要把这些错误信息包装一下在返回给用户
						   http.StatusInternalServerError) 返回状态码
				return 
			}
			defer file.Close()

			all, err := ioutil.ReadAll(file)   用ioutil.ReadAll去读这个file; ReadAll返回([]byte, error);[]byte是内容
			if err != nil {
				panic(err)
			}

			writer.Write(all)  把内容写到http.ResponseWriter里面去
		}
		err := http.ListenAndServe(":8888", nil)  监听端口
		if err != nil {
			panic(err)
		}
}


上面定义的函数func(writer http.ResponseWriter, request *http.Request){} 里面主要是业务逻辑，
我们可以把他们提出来放到filelisting/handler.go文件里

技巧: 业务逻辑尽量不要写在main函数中，要写在其他package中

提出来后
func main() {
	http.HandleFunc("/list/",filelisting.HandleFileList)  filelisting是包名,HandleFileList是提出来的函数
}



*/
package main

import (
	"log"
	"net/http"
	 _ "net/http/pprof"
	"os"

	"imooc.com/ccmouse/learngo/lang/errhandling/filelistingserver/filelisting"
)

//http.HandleFunc("/list/",filelisting.HandleFileList) 因为我们要把filelisting.HandleFileList包装一下
//filelisting.HandleFileList的类型是func(writer http.ResponseWriter,request *http.Request)
//所以定义appHandler也是这个函数类型，返回error
type appHandler func(writer http.ResponseWriter,
	request *http.Request) error

/*
在定义errWrapper函数，形参是appHandler类型的函数，返回一个函数 func(http.ResponseWriter, *http.Request)
然后将handler这个函数返回的error, 传给func(http.ResponseWriter, *http.Request)函数用，这是一个闭包的

func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request){}
*/
func errWrapper(
	handler appHandler) func(
	http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter,
		request *http.Request) {
		// panic
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic: %v", r)
				http.Error(writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()

		err := handler(writer, request)

		if err != nil {
			log.Printf("Error occurred "+
				"handling request: %s",
				err.Error())

			// user error
			if userErr, ok := err.(userError); ok {
				http.Error(writer,
					userErr.Message(),
					http.StatusBadRequest)
				return
			}

			// system error
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer,
				http.StatusText(code), code)
		}
	}
}

type userError interface {
	error
	Message() string
}

func main() {
	http.HandleFunc("/",
		errWrapper(filelisting.HandleFileList))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}

/*
封装一个函数的步骤

比如想封装filelisting.HandleFileList这个函数(自己实现的函数), 这个函数返回func(writer http.ResponseWriter,request *http.Request)

1. 先定义一个和filelisting.HandeFileList返回一样的函数，并自定义返回值

type appHandler func(writer http.ResponseWriter,request *http.Request) error

2. 在定义具体的封装函数errWrapper，最终还是要返回func(writer http.ResponseWriter,request *http.Request)类型

func errWrapper(apphandler appHandler) func(writer http.ResponseWriter,request *http.Request) {}

apphandler函数返回的error,可以传给func(writer http.ResponseWriter,request *http.Request)函数里面用
返回的error是一个自由变量
根据apphandler返回具体error，在去func(writer http.ResponseWriter,request *http.Request)函数里面去做判断返回什么


3. errWrapper这个函数的形参是一个函数，返回值也是一个函数；是一个典型闭包的例子

4. 封装后的函数errWrapper(filelisting.HandlerFilelist)，通过这个方式来调用

*/