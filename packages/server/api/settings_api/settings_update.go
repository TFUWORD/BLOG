package settings_api

import (
	"web/config"
	"web/core"
	"web/global"
	"web/models/res"

	"github.com/gin-gonic/gin"
)

// @Tags 系统管理
// @Summary 更新系统信息
// @Description 根据name字段更新相关系统配置信息
// @Param name path string true "系统配置的字段名称" enum("site", "email", "qq", "qiniu", "jwt")
// @Param updateRequest body interface{} true "更新的信息"
// @Router /api/settings/{name} [put]
// @Accept json
// @Produce json
// @Success 200 {object} res.Response
func (SettingsApi) SettingsInfoUpdateView (c *gin.Context) {

	var cr SettingsUri

	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	
	switch cr.Name {
	case "site":
		var info config.SiteInfo
		err := c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.SiteInfo = info

	case "email":
		var email config.Email
		err := c.ShouldBindJSON(&email)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.Email = email

	case "qq":
		var qq config.QQ
		err := c.ShouldBindJSON(&qq)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.QQ = qq

	case "qiniu":
		var qiniu config.QiNiu
		err := c.ShouldBindJSON(&qiniu)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.QiNiu = qiniu

	case "jwt":
		var jwt config.Jwt
		err := c.ShouldBindJSON(&jwt)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.Jwt = jwt

	default:
		res.FailWithMessage("没有对应配置信息", c)
	}

	core.SetYaml()
	res.OkWith(c)
}
