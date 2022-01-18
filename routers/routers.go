package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//此文件定义所有的路由信息

func SetupRouter() *gin.Engine {
	//创建一个路由实例
	r := gin.Default()

	//路由
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	//将路由实例返回
	return r
}
