package user_api

import (
	"web/global"
	"web/models"
	"web/models/res"
	"web/utils/jwts"
	"web/utils/pwd"

	"github.com/gin-gonic/gin"
)

type UpdatePasswordRequest struct {
  OldPwd string `json:"old_pwd" binding:"required" msg:"请输入旧密码"` // 旧密码
  Pwd    string `json:"pwd" 		binding:"required" msg:"请输入新密码"`     // 新密码
}

// @Tags 用户管理
// @Summary 修改用户密码
// @Description 传入验证信息修改密码
// @Param cr body UpdatePasswordRequest true "密码更新请求体"
// @Router /api/user_password [put]
// @Accept json
// @Produce json
// @Success 200 {object} res.Response
func (UserApi) UserUpdatePassword(c *gin.Context) {
  _claims, _ := c.Get("claims")
  claims := _claims.(*jwts.CustomClaims)
  var cr UpdatePasswordRequest
  if err := c.ShouldBindJSON(&cr); err != nil {
    res.FailWithError(err, &cr, c)
    return
  }
  var user models.UserModel
  err := global.DB.Take(&user, claims.UserID).Error
  if err != nil {
    res.FailWithMessage("用户不存在", c)
    return
  }
  // 判断密码是否一致
  if !pwd.CheckPwd(user.Password, cr.OldPwd) {
    res.FailWithMessage("密码错误", c)
    return
  }
  hashPwd := pwd.HashPwd(cr.Pwd)
  err = global.DB.Model(&user).Update("password", hashPwd).Error
  if err != nil {
    global.Log.Error(err)
    res.FailWithMessage("密码修改失败", c)
    return
  }
  res.OkWithMessage("密码修改成功", c)
  // return
}
