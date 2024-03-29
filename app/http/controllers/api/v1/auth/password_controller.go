package auth

import (
	"github.com/gin-gonic/gin"
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/app/models/user"
	"gohub/app/requests"
	"gohub/pkg/response"
)

// 用户控制器
type PasswordController struct {
	v1.BaseAPIController
}

// 使用手机和验证码重置密码
func (pc *PasswordController) ResetByPhone(c *gin.Context) {

	// 1.验证表单
	request := requests.RestByPhoneRequest{}
	if ok := requests.Validate(c, &request, requests.RestByPhone); !ok {
		return
	}

	// 2.更新密码
	userModel := user.GetByPhone(request.Phone)
	if userModel.ID == 0 {
		response.Abort404(c)
	} else {
		userModel.Password = request.Password
		userModel.Save()

		response.Success(c)
	}
}

// 使用邮箱重置密码
func (pc *PasswordController) ResetByEmail(c *gin.Context) {
	// 1.验证表单
	request := requests.RestByEmailRequest{}
	if ok := requests.Validate(c, &request, requests.RestByEmail); !ok {
		return
	}

	// 2.更新密码
	userModel := user.GetByEmail(request.Email)
	if userModel.ID == 0 {
		response.Abort404(c)
	} else {
		userModel.Password = request.Password
		userModel.Save()
		response.Success(c)
	}
}
