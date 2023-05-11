package config

import "fmt"

type QQ struct {
  AppID    string `yaml:"app_id"  json:"app_id"`
  Key      string `yaml:"key"     json:"key"`
  Redirect string `yaml:"redirect"json:"redirect"` //登录后的回调地址
}

func (qq QQ) GetPath() string {
  if qq.Key == "" || qq.AppID == "" || qq.Redirect == "" {
    return ""
  }
  return fmt.Sprintf("https://graphqq.com/oauth2.0/show?which=Login&display=pc&respone_type=code&client_id=%s&redirect_url=%s", qq.AppID, qq.Redirect)
}
