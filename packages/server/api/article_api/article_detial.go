package article_api

import (
	"web/models"
	"web/models/res"
	"web/service/es_ser"
	"web/service/redis_ser"

	"github.com/gin-gonic/gin"
)

func (ArticleApi) ArticleDetailView(c *gin.Context) {
	var cr models.ESIDRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	redis_ser.Look(cr.ID)
	model, err := es_ser.CommDetail(cr.ID)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithData(model, c)
}

type ArticleDetailRequest struct {
	Title string `json:"title" form:"title"`
}

func (ArticleApi) ArticleDetailByTitleView(c *gin.Context) {
	var cr ArticleDetailRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	// redis_ser.Look(cr.ID)
	model, err := es_ser.CommDetailByKeyword(cr.Title)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithData(model, c)
}
