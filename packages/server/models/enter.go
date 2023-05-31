package models

import "time"

type MODEL struct {
	ID        uint      `gorm:"primarykey" json:"id"` // 主键id
	CreatedAt time.Time `json:"created_at"`           // 创建时间
	UpdatedAt time.Time `json:"-"`                    // 更新时间
}

// 删除图片
type RemoveRequest struct {
	IDList []uint `json:"id_list"`
}

// 搜索id请求
type ESIDRequest struct {
	ID string `json:"id" form:"id" uri:"id"`
}

// 搜索id列表请求
type ESIDListRequest struct {
	IDList []string `json:"id_list" binding:"required"`
}

// 页数
type PageInfo struct {
	Page  int    `form:"page"`
	Key   string `form:"key"`
	Limit int    `form:"limit"`
	Sort  string `form:"sort"`
}
