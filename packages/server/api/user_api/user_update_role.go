package user_api

import (
	"web/global"
	"web/models"
	"web/models/ctype"
	"web/models/res"

	"github.com/gin-gonic/gin"
)

type UserRole struct {
  Role     ctype.Role `json:"role" binding:"required,oneof=1 2 3 4" msg:"权限参数错误"`
  NickName string     `json:"nick_name"` // 防止用户昵称非法，管理员有能力修改
  UserID   uint       `json:"user_id" binding:"required" msg:"用户id错误"`
}

// @Tags 用户管理
// @Summary 用户权限更新
// @Description 传入选中待修改权限的用户进行权限修改
// @Param cr body UserRole true "更新请求体"
// @Router /api/user_role [put]
// @Accept json
// @Produce json
// @Success 200 {object} res.Response
func (UserApi) UserUpdateRoleView(c *gin.Context) {
  var cr UserRole
  if err := c.ShouldBindJSON(&cr); err != nil {
    res.FailWithError(err, &cr, c)
    return
  }
  var user models.UserModel
  err := global.DB.Take(&user, cr.UserID).Error
  if err != nil {
    res.FailWithMessage("用户id错误，用户不存在", c)
    return
  }
  err = global.DB.Model(&user).Updates(map[string]any{
    "role":      cr.Role,
    "nick_name": cr.NickName,
  }).Error
  if err != nil {
    global.Log.Error(err)
    res.FailWithMessage("修改权限失败", c)
    return
  }
  res.OkWithMessage("修改权限成功", c)
}
