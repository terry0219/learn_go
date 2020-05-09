package real

import (
	"net/http"
	"net/http/httputil"
	"time"
)

type Retriever struct {          //定义结构体
	UserAgent string
	TimeOut   time.Duration
}

func (r *Retriever) Get(url string) string {  //为结构体定义Get方法
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	result, err := httputil.DumpResponse(
		resp, true)

	resp.Body.Close()

	if err != nil {
		panic(err)
	}

	return string(result)
}
