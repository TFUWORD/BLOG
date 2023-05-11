package utils

import (
	"reflect"

	"github.com/go-playground/validator/v10"
)

// 信息源结构体
func GetValidMsg(err error, obj any) string {
	getObj := reflect.TypeOf(obj)

	if errs, ok := err.(validator.ValidationErrors); ok {

		for _, e := range errs {

			if f, exits := getObj.Elem().FieldByName(e.Field()); exits {
				msg := f.Tag.Get("msg")
				return msg
			}
		}
	}
	return err.Error()
}
