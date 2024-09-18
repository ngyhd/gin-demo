package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model // https://gorm.io/zh_CN/docs/models.html#%E5%88%9B%E5%BB%BA-x2F-%E6%9B%B4%E6%96%B0%E6%97%B6%E9%97%B4%E8%BF%BD%E8%B8%AA%EF%BC%88%E7%BA%B3%E7%A7%92%E3%80%81%E6%AF%AB%E7%A7%92%E3%80%81%E7%A7%92%E3%80%81Time%EF%BC%89
	Username   string
	Password   string
	Email      string
}
