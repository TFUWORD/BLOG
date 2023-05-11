package user_api

import (
	"web/global"
	"web/models"
	"web/models/res"
	"web/utils/jwts"
	"web/utils/pwd"

	"github.com/gin-gonic/gin"
)

type EmailLoginRequest struct {
  UserName string `json:"user_name" binding:"required" msg:"请输入用户名"`  // 用户或者邮箱
  Password string `json:"password" binding:"required" msg:"请输入密码"`
}

// @Tags 用户管理
// @Summary 用户登录
// @Description 传入登录信息登录账号
// @Param cr body EmailLoginRequest true "注册请求体"
// @Router /api/email_login [post]
// @Accept json
// @Produce json
// @Success 200 {object} res.Response
func (UserApi) EmailLoginView(c *gin.Context) {
  var cr EmailLoginRequest
  err := c.ShouldBindJSON(&cr)
  if err != nil {
    res.FailWithError(err, &cr, c)
    return
  }

  var userModel models.UserModel
  err = global.DB.Take(&userModel, "user_name = ? or email = ?", cr.UserName, cr.UserName).Error
  if err != nil {
    // 没找到
    global.Log.Warn("用户名不存在")
    res.FailWithMessage("用户名或密码错误", c)
    return
  }
  // 校验密码
  isCheck := pwd.CheckPwd(userModel.Password, cr.Password)
  if !isCheck {
    global.Log.Warn("用户名密码错误")
    res.FailWithMessage("用户名或密码错误", c)
    return
  }
  // 登录成功，生成token
  token, err := jwts.GenToken(jwts.JwtPayLoad{
    NickName: userModel.NickName,
    Role:     int(userModel.Role),
    UserID:   userModel.ID,
  })
  if err != nil {
    global.Log.Error(err)
    res.FailWithMessage("token生成失败", c)
    return
  }
  res.OkWithData(token, c)

}
