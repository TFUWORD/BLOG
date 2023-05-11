package images_api

import (
	"fmt"
	"web/global"
	"web/models"
	"web/models/res"

	"github.com/gin-gonic/gin"
)

// @Tags 图片管理
// @Summary 图片删除
// @Description 传入待删除的图片ID集合id_list，根据ID执行删除操作
// @Param cr body models.RemoveRequest true "图片ID列表"
// @Router /api/images [delete]
// @Accept json
// @Produce json
// @Success 200 {object} res.Response
func (ImagesApi)ImageRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var imageList []models.BannerModel
	count := global.DB.Find(&imageList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("待删除文件不存在", c)
		return
	}
	global.DB.Delete(&imageList)
	res.OkWithMessage(fmt.Sprintf("共删除图片 %d 张", count), c)
}
