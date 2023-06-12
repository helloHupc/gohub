package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	// new一个gin Engine实例
	r := gin.New()

	// 注册中间件
	r.Use(gin.Logger(), gin.Recovery())

	// 注册一个路由
	r.GET("/", func(c *gin.Context) {

		//json格式响应
		c.JSON(http.StatusOK, gin.H{
			"hello": "world!",
		})
	})

	// 处理404请求
	r.NoRoute(func(c *gin.Context) {
		// 获取表头信息 Accept 信息
		acceptString := c.Request.Header.Get("Accept")

		if strings.Contains(acceptString, "text/html") {
			//如果是HTML的话
			c.String(http.StatusNotFound, "页面返回 404")
		} else {
			// 默认返回 JSON
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义，请确认URL和请求方法是否正确。",
			})
		}
	})

	// 运行服务
	r.Run(":8000")
}
