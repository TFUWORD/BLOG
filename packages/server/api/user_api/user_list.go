package user_api

import (
	"fmt"
	"web/models"
	"web/models/ctype"
	"web/models/res"
	"web/service/common"
	"web/utils/desens"
	"web/utils/jwts"

	"github.com/gin-gonic/gin"
)

// @Tags 用户管理
// @Summary 用户列表
// @Description 传入页面参数获取用户信息
// @Param page query models.PageInfo false "页面信息"
// @Router /api/users [get]
// @Accept json
// @Produce json
// @Success 200 {object} res.Response
func (UserApi) UserListView(c *gin.Context) {
  // 如何判断是管理员
  _claims, _ := c.Get("claims")
  claims := _claims.(*jwts.CustomClaims)

  var page models.PageInfo
  if err := c.ShouldBindQuery(&page); err != nil {
    res.FailWithCode(res.ArgumentError, c)
    return
  }
  var users []models.UserModel
  list, count, _ := common.ComList(models.UserModel{}, common.Option{
    PageInfo: page,
  })
  for _, user := range list {
    if ctype.Role(claims.Role) != ctype.PermissionAdmin {
      // 非管理员
      user.UserName = ""
    }
    user.Tel = desens.DesensitizationTel(user.Tel)
    user.Email = desens.DesensitizationEmail(user.Email)
    // 脱敏
    users = append(users, user)
    fmt.Printf(user.NickName)
  }

  res.OkWithList(users, count, c)
}

