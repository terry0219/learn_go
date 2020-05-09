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
	fmt.Printf("%s\n",all) //如果这里输出有乱码的话，看下编码如果是gbk的话，需要转为utf-8

	//提取城市和url
	printCityList(all)  
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

/*
写正则技巧
1. 先把要获取的城市和url，复制到正则表达式里
2. 
*/
func printCityList(contents []byte) {
	//regexp.MustCompile(`<a href="http://www.zhenai.com/zhenghun/yuxi" class="" >玉溪</a>`)
	//regexp.MustCompile(`<a href="http://www.zhenai.com/zhenghun/[0-9a-z]+" class="" >玉溪</a>`)
	//regexp.MustCompile(`<a href="http://www.zhenai.com/zhenghun/[0-9a-z]+"[^>]*>玉溪</a>`)       [^>]*  匹配0个或多个字符，除了>
	re := regexp.MustCompile(`<a href="http://www.zhenai.com/zhenghun/[0-9a-z]+"[^>]*>[^<]+</a>`)       [^<]+ 匹配1个或多个字符，除了<
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`) 
	//上面正则第一个括号是url地址，第二个括号是匹配城市名称
	
	//matches := re.FindAll(contents, -1) 如果要提取城市名称和Url，需要在正则表达式加上对应的括号
	matches := re.FindAllSubmatch(contents,-1)

	for _,m := matches {
		for _, subMatch := m {
			fmt.Printf("%s ",subMatch)
		}
		fmt.Println() //换行
	}

	for _,m := range matches {
		fmt.Printf("%s\n",m) //会输出所有匹配到的字符  <a href="http://www.zhenai.com/zhenghun/???" class="" >???</a>
	}


}