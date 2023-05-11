package user_ser

import (
	"fmt"
	"web/utils/pwd"
	"web/global"
	"web/models"
	"web/models/ctype"
	"time"
	"web/service/redis_ser"
	"web/utils/jwts"
  "errors"
)

type UserService struct {
}

func (UserService) Logout(claims *jwts.CustomClaims, token string) error {
  exp := claims.ExpiresAt
  now := time.Now()
  diff := exp.Time.Sub(now)
  return redis_ser.Logout(token, diff)
}

func (UserService) CreateUser(userName, nickName, password string, role ctype.Role, email string, ip string) error {
  var userModel models.UserModel
  err := global.DB.Take(&userModel, "user_name = ?", userName).Error
  if err == nil {
    fmt.Println(err)
    return errors.New("用户名已经存在")
  }

  hashPwd := pwd.HashPwd(password)

  // 入库
  err = global.DB.Create(&models.UserModel{
    NickName:   nickName,
    UserName:   userName,
    Password:   hashPwd,
    Avatar:     global.Config.Avatar.Pic,
    Email:      email,
    Addr:       "内网地址",
    IP:         ip,
    Role:       role,
    SignStatus: ctype.SignSite,
  }).Error
  if err != nil {
    return err
  }
  return nil
}
