package images_api

import (
	"web/global"
	"web/models"
	"web/models/res"

	"github.com/gin-gonic/gin"
)

type ImageResponse struct{
	ID 	 uint   `json:"id"`
	Path string `json:"path"`                        // 图片路径
	Name string `gorm:"size:38" json:"name"`         // 图片名称
}

// @Tags 图片管理
// @Summary 全部图片展示
// @Description 无需输入参数，直接返回查询图片库所有图片的id，path，name 
// @Router /api/image_names [get]
// @Accept json
// @Produce json
// @Success 200 {object} res.Response
func (ImagesApi) ImageNameListView(c *gin.Context) {
	var imageList []ImageResponse
	global.DB.Model(models.BannerModel{}).Select("id", "path", "name").Scan(&imageList)
	res.OkWithData(imageList, c)
}
