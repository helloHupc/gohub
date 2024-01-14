package v1

import (
	"gohub/app/models/user"
	"gohub/pkg/auth"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	BaseAPIController
}

// CurrentUser 当前登录用户信息
func (ctrl *UsersController) CurrentUser(c *gin.Context) {
	userModel := auth.CurrentUser(c)
	response.Data(c, userModel)
}

func (ctrl *UsersController) Show(c *gin.Context) {
	userModel := user.Get(c.Param("id"))
	if userModel.ID == 0 {
		response.Abort404(c)
		return
	}
	response.Data(c, userModel)
}
