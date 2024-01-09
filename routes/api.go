// package routes 注册路由
package routes

import (
	"github.com/gin-gonic/gin"
	"gohub/app/http/controllers/api/v1/auth"
)

// 注册项目相关路由
func RegisterAPIRoutes(r *gin.Engine) {
	// 测试一个v1路由组

	v1 := r.Group("/v1")
	{

		authGroup := v1.Group("/auth")
		{
			suc := new(auth.SignupController)

			// 判断手机是否已注册
			authGroup.POST("/signup/phone/exist", suc.IsPhoneExist)
			// 判断 Email 是否已注册
			authGroup.POST("/signup/email/exist", suc.IsEmailExist)
			// 手机号注册
			authGroup.POST("/signup/using-phone", suc.SignupUsingPhone)
			// 邮箱注册
			authGroup.POST("/signup/using-email", suc.SignupUsingEmail)

			// 发送验证码
			vcc := new(auth.VerifyCodeController)
			authGroup.POST("/verify-codes/captcha", vcc.ShowCaptcha)
			authGroup.POST("/verify-codes/phone", vcc.SendUsingPhone)
			authGroup.POST("/verify-codes/email", vcc.SendUsingEmail)
		}
	}
}
