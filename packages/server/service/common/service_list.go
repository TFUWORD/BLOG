package common

import (
	"web/global"
	"web/models"

	"gorm.io/gorm"
)

type Option struct {
	models.PageInfo
	Debug bool
}

func ComList[T any](model T, option Option) (list []T, count int64, err error) {

	DB := global.DB
	if option.Debug {
		DB = global.DB.Session(&gorm.Session{Logger: global.MysqlLog})
	}

	if option.Sort =="" {
		option.Sort = "created_at desc"
	}
	query := DB.Where(model)
	count = query.Select("id").Find(&list).RowsAffected
	query = DB.Where(model)
	offset := (option.Page - 1) * option.Limit
	if offset < 0 {
		offset = 0
	}
	// Find(&list) 方法用于执行查询并将结果存储到 list 变量中
	err = query.Limit(option.Limit).Offset(offset).Order(option.Sort).Find(&list).Error

	return list, count, err
}
