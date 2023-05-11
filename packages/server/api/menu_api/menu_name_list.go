package menu_api

import (
	"web/global"
	"web/models"
	"web/models/res"

	"github.com/gin-gonic/gin"
)

type MenuNameResponse struct {
  ID    uint   `json:"id"`
  Title string `json:"title"`
  Path  string `json:"path"`
}

// @Tags 菜单管理
// @Summary 查询菜单
// @Description 查询所有菜单栏并只返回菜单编号id，标题title，跳转链接path
// @Router /api/menu_names [get]
// @Accept json
// @Produce json
// @Success 200 {object} res.Response
func (MenuApi) MenuNameList(c *gin.Context) {
  var menuNameList []MenuNameResponse
  global.DB.Model(models.MenuModel{}).Select("id", "title", "path").Scan(&menuNameList)
  if menuNameList == nil {
    res.OkWithData([]MenuNameResponse{}, c)
    return
  }
  res.OkWithData(menuNameList, c)
}
