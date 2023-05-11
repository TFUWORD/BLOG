package images_api

import (
	"web/models"
	"web/models/res"
	"web/service/common"

	"github.com/gin-gonic/gin"
)


// ImageListView 图片列表查询
// @Tags 图片管理
// @Summary 图片分页查询
// @Description 传入待查询的图片页数几每页限制几张图片，返回查询结果
// @Param page query int false "页码"
// @Param limit query int false "每页的图片输"
// @Router /api/images [get]
// @Accept json
// @Produce json
// @Success 200 {object} res.Response
func (ImagesApi) ImageListView(c *gin.Context) {
	var cr models.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	imageList, count, err := common.ComList(models.BannerModel{}, common.Option{
		PageInfo: cr,
		Debug: 		true,
	})

	res.OkWithList(imageList, count, c)
	
}
