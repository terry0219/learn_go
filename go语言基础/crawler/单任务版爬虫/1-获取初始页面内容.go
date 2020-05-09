/*
获取并打印所有城市第一页用户的详细信息

封装函数的步骤:
1. 先看下被包装的函数返回的类型
2. 定义封装函数，返回的类型要和被包装函数返回的类型一样
3. 被包装的函数要在包装函数里调用执行
*/

/*
package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func main() {
	//先获取城市列表
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code",resp.StatusCode)
		return
	}

	//go get github.com/golang.org/x/text
	//resp.Body是源reader，simplifiedchinese.GBK.NewDecoder是转换后的编码
	utf8Reader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())//把resp.Body转换为gbk
	all, err := ioutil.ReadAll(utf8Reader) //这样就可以读取gbk的编码了；但是这里有一个问题是通用性很差，不适合其他网站

	if err != nil {
		panic(err)
	}
	//fmt.Printf("%s\n",all) //如果这里输出有乱码的话，看下编码如果是gbk的话，需要转为utf-8
}

*/
-------------------------------------------------------------

package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func main() {
	//先获取城市列表
	/*
		获取html页面的编码有两种方式
		1. 根据返回的<meta charset="gbk">,但是这个不一定准确

		2. github.com/golang.org/x/net/html有一个包可以自动检测字符集 
	*/
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code",resp.StatusCode)
		return
	}

	//这个函数会去读前1024个byte去猜是什么字符集, 返回encoding.Encoding
	//charset.DetermineEncoding(content []byte, contentType string) (e encoding.Encoding, name string, cerain bool) {}
	//调用包装的函数,返回endcoding
	e := determineEncoding(resp.Body)
	//go get github.com/golang.org/x/text
	//resp.Body是源reader，simplifiedchinese.GBK.NewDecoder是转换后的编码
	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
	all, err := ioutil.ReadAll(utf8Reader) //这样就可以读取gbk的编码了；但是这里有一个问题是通用性很差，不适合其他网站

	if err != nil {
		panic(err)
	}
	//fmt.Printf("%s\n",all) //如果这里输出有乱码的话，看下编码如果是gbk的话，需要转为utf-8
}

//包装一下charset.DetermineEncoding函数
func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024) //读1024Byte
	if err != nil {
		panic(err)
	}
	e,_,_ := charset.DetermineEncoding(bytes,"")
	return e
}

