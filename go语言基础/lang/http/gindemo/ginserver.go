/*
第三方 gin框架介绍

gin-gonic/gin
middleware的使用
context的使用

下载gin
go get -u github.com/gin-gonic/gin

导入gin
import github.com/gin-gonic/gin

官网demo:

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":"pong",
		})
	})
	r.GET("/hello", func(c *gin.Context) {
		c.string(200, "hello")
	})

	r.Run()
}


当GET请求http://xxxx:8080/ping, gin会返回一个JSON {"message":"pong"}


为gin增加middleware，意思就是无论访问是/ping还是/hello，都要先过middleware
这里通过middleware，加上自定义的日志信息

func main() {
	r := gin.Default()
	logger, err := zap.NewProduction()  //使用zap包来记录日志
	if err != nil {
		fmt.Println(err)
	}

	r.Use(func(c *gin.Context) {
		s := time.Now()
		c.Next()

		//打印日志 输出JSON格式 {} 记录 PATH CODE EXEC_TIME {"path":"/ping","status":200,"elapsed":0.01}
		logger.Info("Incoming request", zap.String("path", c.Request.URL.Path),
										zap.Int("status", c.Writer.Status()),
										zap.Duration("elapsed", time.Now().sub(s))) 
		
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":"pong",
		})
	})
	r.GET("/hello", func(c *gin.Context) {
		c.string(200, "hello")
	})

	r.Run()
}


*/
package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"math/rand"
	"time"
)

const keyRequestId = "requestId"

func main() {
	r := gin.Default()
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	r.Use(func(c *gin.Context) {
		s := time.Now()

		c.Next()

		logger.Info("incoming request",
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.Duratio"elapsed", time.Now().Sub(s)))
	}, func(c *gin.Context) {
		c.Set(keyRequestId, rand.Int()) //为每次请求set一个requestId

		c.Next()
	})

	r.GET("/ping", func(c *gin.Context) {
		h := gin.H{
			"message":   "pong",
		}
		if rid, exists := c.Get(keyRequestId); exists { //取requestId
			h[keyRequestId] = rid
		}
		c.JSON(200, h)
	})
	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "hello")
	})
	r.Run()
}
