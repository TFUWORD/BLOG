package user_api

import (
	"fmt"
	"web/global"
	"web/models"
	"web/models/res"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Tags 用户管理
// @Summary 用户删除
// @Description 传入待删除用户id_list后根据id删除每一个用户
// @Param cr body models.RemoveRequest true "删除请求体"
// @Router /api/users [delete]
// @Accept json
// @Produce json
// @Success 200 {object} res.Response
func (UserApi) UserRemoveView(c *gin.Context) {
  var cr models.RemoveRequest
  err := c.ShouldBindJSON(&cr)
  if err != nil {
    res.FailWithCode(res.ArgumentError, c)
    return
  }

  var userList []models.UserModel
  count := global.DB.Find(&userList, cr.IDList).RowsAffected
  if count == 0 {
    res.FailWithMessage("用户不存在", c)
    return
  }

  // 事务
  err = global.DB.Transaction(func(tx *gorm.DB) error {
    // 删除用户，消息表，评论表，用户收藏的文章，用户发布的文章
    err = global.DB.Delete(&userList).Error
    if err != nil {
      global.Log.Error(err)
      return err
    }
    return nil
  })
  if err != nil {
    global.Log.Error(err)
    res.FailWithMessage("删除用户失败", c)
    return
  }
  res.OkWithMessage(fmt.Sprintf("共删除 %d 个用户", count), c)

}
