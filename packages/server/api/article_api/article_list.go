package article_api

import (
	"web/global"
	"web/models"
	"web/models/res"
	"web/service/es_ser"

	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
)

type ArticleSearchRequest struct {
	models.PageInfo
	Tag string `json:"tag" form:"tag"`
}

func (ArticleApi) ArticleListView(c *gin.Context) {
	var cr ArticleSearchRequest
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	list, count, err := es_ser.CommList(es_ser.Option{
		PageInfo: cr.PageInfo,
		Fields:   []string{"title", "content"},
		Tag:      cr.Tag,
	})
	if err != nil {
		global.Log.Error(err)
		res.OkWithMessage("查询失败", c)
		return
	}

	data := filter.Omit("list", list)
	_list, _ := data.(filter.Filter)
	if string(_list.MustMarshalJSON()) == "{}" {
		list = make([]models.ArticleModel, 0)
		res.OkWithList(list, int64(count), c)
		return
	}

	res.OkWithList(filter.Omit("list", list), int64(count), c)
}
