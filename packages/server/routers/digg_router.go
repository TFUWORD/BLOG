package routers

import (
	"web/api"
)

func (router RouterGroup) DiggRouter() {
	app := api.ApiGroupApp.DiggApi
	router.POST("digg/article", app.DiggArticleView)
}
