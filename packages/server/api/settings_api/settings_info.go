package settings_api

import (
	"web/global"
	"web/models/res"

	"github.com/gin-gonic/gin"
)

type SettingsUri struct {
	Name string `uri:"name"`
}

// @Tags 系统管理
// @Summary 查看系统信息
// @Description 根据name字段查看相关系统配置信息
// @Param name path string true "系统配置的字段名称" enum("site", "email", "qq", "qiniu", "jwt")
// @Router /api/settings/{name} [get]
// @Accept json
// @Produce json
// @Success 200 {object} res.Response
func (SettingsApi) SettingsInfoView (c *gin.Context) {
	var cr SettingsUri

	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	switch cr.Name {
	case "site":
		res.OkWithData(global.Config.SiteInfo, c)
	case "email":
		res.OkWithData(global.Config.Email, c)
	case "qq":
		res.OkWithData(global.Config.QQ, c)
	case "qiniu":
		res.OkWithData(global.Config.QiNiu, c)
	case "jwt":
		res.OkWithData(global.Config.Jwt, c)
	default:
		res.FailWithMessage("没有对应配置信息", c)
	}

}
