package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1、创建路由
	r := gin.Default()  //返回默认的路由引擎
	// 2、绑定路由规则、执行的函数
	r.GET("/",func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK,gin.H{
			"message": "hello gin",
			"username": "zhangsan",
		})
	})
	// 3、监听端口 不写端口默认是8080
	r.Run(":8888")

}