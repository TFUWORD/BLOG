package api

import (
	"web/api/article_api"
	"web/api/digg_api"
	"web/api/images_api"
	"web/api/menu_api"
	"web/api/settings_api"
	"web/api/user_api"
)

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
	ImagesApi   images_api.ImagesApi
	MenuApi     menu_api.MenuApi
	UserApi     user_api.UserApi
	ArticleApi  article_api.ArticleApi
	DiggApi     digg_api.DiggApi
}

var ApiGroupApp = new(ApiGroup)
