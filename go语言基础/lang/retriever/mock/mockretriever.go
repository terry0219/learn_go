package mock

import "fmt"

//Retriever这个结构体实现了Get\Post方法
type Retriever struct {  //定义Retriever为struct类型
	Contents string
}

func (r *Retriever) String() string {  //定义String方法，fmt.Println会打印自定义的string()
	return fmt.Sprintf(
		"Retriever: {Contents=%s}", r.Contents)
}

//为Retriever定义Post方法
func (r *Retriever) Post(url string,
	form map[string]string) string {
	r.Contents = form["contents"]
	return "ok"
}

func (r *Retriever) Get(url string) string { //为Retriever定义Get方法
	return r.Contents
}
