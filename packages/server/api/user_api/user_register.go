package user_api

import(
	"web/models/ctype"
	"web/models/res"
	"web/service/user_ser"
	"web/global"
	"fmt"
	
	"github.com/gin-gonic/gin"
)

type UserCreateRequest struct {
	NickName   string			`json:"nick_name" binding:"required" msg:"请输入昵称"`
	UserName   string			`json:"user_name" binding:"required" msg:"请输入用户名"`
	Password   string			`json:"password"  binding:"required" msg:"请输入密码"`
	Role       ctype.Role `json:"role" 			binding:"required" msg:"请输入权限"`	// 1 管理员   2 普通用户   3 游客
}

// @Tags 用户管理
// @Summary 用户注册
// @Description 传入注册信息注册新用户
// @Param cr body UserCreateRequest true "注册信息"
// @Router /api/register [post]
// @Accept json
// @Produce json
// @Success 200 {object} res.Response
func (UserApi) UserCreateView(c *gin.Context) {
	var cr UserCreateRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	err := user_ser.UserService{}.CreateUser(cr.UserName, cr.NickName, cr.Password, cr.Role, "", c.ClientIP())
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithMessage(fmt.Sprintf("用户%s创建完成", cr.UserName), c)
	return
}
