package ctype

import "encoding/json"

type SignStatus int

const (
	SignQQ 		SignStatus = 1  // QQ登录
	SignSite	SignStatus = 2	// Site登录
	SignEmail SignStatus = 3  // Email登录
)

func (s SignStatus) MaeshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s SignStatus) String() string{
	var str string

	switch s {
	case SignQQ :
		str = "QQ登录"
	case SignSite :
		str = "Site登录"
	case SignEmail :
		str = "Email登录"
	default:
		str = "其他"
	}

	return str
}
