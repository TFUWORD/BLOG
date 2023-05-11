package images_api

import (
	"web/global"
	"web/models"
	"web/models/res"

	"github.com/gin-gonic/gin"
)

type ImageUpdateRequest struct {
	ID 	 uint 	`json:"id" 	 binding:"required" msg:"输入文件id"`
	Name string `json:"name" binding:"required" msg:"输入文件名称"`
}

// @Tags 图片管理
// @Summary 图片编辑
// @Description 传入待编辑的图片id及修改后的名称name，返回修改成功
// @Param cr body ImageUpdateRequest true "修改后的图片名称"
// @Router /api/images [put]
// @Accept json
// @Produce json
// @Success 200 {object} res.Response
func (ImagesApi)ImageUpdateView(c *gin.Context) {
	var cr ImageUpdateRequest
	err := c.ShouldBindJSON(&cr)

	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	var imageModel models.BannerModel
	err = global.DB.Take(&imageModel, cr.ID).Error
	if err != nil {
		res.FailWithMessage("文件不存在", c)
		return
	}

	err = global.DB.Model(&imageModel).Update("name", cr.Name).Error
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}

	res.OkWithMessage("图片名称修改成功", c)
	return
}
