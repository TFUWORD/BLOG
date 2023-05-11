package service

import (
	"web/service/image_ser"
	"web/service/user_ser"
)

type ServiceGroup struct{
	ImageService image_ser.ImageService
	UserService  user_ser.UserService
}

var ServiceApp = new(ServiceGroup)
