// package routes 注册路由
package routes

import (
	"gohub/app/http/controllers/api/v1/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 注册项目相关路由
func RegisterAPIRoutes(r *gin.Engine) {
	// 测试一个v1路由组

	v1 := r.Group("/v1")
	{
		// 注册一个路由
		v1.GET("/", func(c *gin.Context) {
			// 以json格式响应
			c.JSON(http.StatusOK, gin.H{
				"hello": "world!",
			})
		})

		authGroup := v1.Group("/auth")
		{
			suc := new(auth.SignupController)

			// 判断手机是否已注册
			authGroup.POST("/signup/phone/exist", suc.IsPhoneExist)
			// 判断 Email 是否已注册
			authGroup.POST("/signup/email/exist", suc.IsEmailExist)
		}
	}
}
