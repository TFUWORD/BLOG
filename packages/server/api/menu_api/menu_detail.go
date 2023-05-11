package menu_api

import (
	"web/global"
	"web/models"
	"web/models/res"

	"github.com/gin-gonic/gin"
)

// @Tags 菜单管理
// @Summary 菜单详情
// @Description 查询编号id的菜单图片信息
// @Router /api/menus/{id} [get]
// @Param id path uint true "菜单id"
// @Accept json
// @Produce json
// @Success 200 {object} res.Response
func (MenuApi) MenuDetailView(c *gin.Context) {
  // 先查菜单
  id := c.Param("id")
  var menuModel models.MenuModel
  err := global.DB.Take(&menuModel, id).Error
  if err != nil {
    res.FailWithMessage("菜单不存在", c)
    return
  }
  // 查连接表
  var menuBanners []models.MenuBannerModel
  global.DB.Preload("BannerModel").Order("sort desc").Find(&menuBanners, "menu_id = ?", id)
  var banners = make([]Banner, 0)
  for _, banner := range menuBanners {
    if menuModel.ID != banner.MenuID {
      continue
    }
    banners = append(banners, Banner{
      ID:   banner.BannerID,
      Path: banner.BannerModel.Path,
    })
  }
  menuResponse := MenuResponse{
    MenuModel: menuModel,
    Banners:   banners,
  }
  res.OkWithData(menuResponse, c)
  return
}
