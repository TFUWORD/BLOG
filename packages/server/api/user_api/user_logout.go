package user_api

import (
	"web/global"
	"web/models/res"
	"web/service"
	"web/utils/jwts"

	"github.com/gin-gonic/gin"
)

// @Tags 用户管理
// @Summary 用户注销
// @Description 用户注销请求
// @Router /api/logout [post]
// @Accept json
// @Produce json
// @Success 200 {object} res.Response
func (UserApi) LogoutView(c *gin.Context) {
  _claims, _ := c.Get("claims")
  claims := _claims.(*jwts.CustomClaims)

  token := c.Request.Header.Get("token")

  err := service.ServiceApp.UserService.Logout(claims, token)

  if err != nil {
    global.Log.Error(err)
    res.FailWithMessage("注销失败", c)
    return
  }

  res.OkWithMessage("注销成功", c)

}
