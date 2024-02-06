package model

import "gorm.io/plugin/soft_delete"

type User struct {
	Id        int
	Username  string
	Password  string
	Email     string
	CreatedAt int                   // gorm GORM 在创建、更新时会自动填充 当前时间 参考:https://gorm.io/zh_CN/docs/models.html#%E5%88%9B%E5%BB%BA-x2F-%E6%9B%B4%E6%96%B0%E6%97%B6%E9%97%B4%E8%BF%BD%E8%B8%AA%EF%BC%88%E7%BA%B3%E7%A7%92%E3%80%81%E6%AF%AB%E7%A7%92%E3%80%81%E7%A7%92%E3%80%81Time%EF%BC%89
	UpdatedAt int                   // gorm GORM 在创建、更新时会自动填充 当前时间
	DeletedAt soft_delete.DeletedAt // 模型包含了 gorm.DeletedAt（datetime格式）字段，那么该模型将会自动获得软删除的能力！ soft_delete.DeletedAt(时间戳格式) 参考:https://gorm.io/zh_CN/docs/delete.html#Unix-%E6%97%B6%E9%97%B4%E6%88%B3
}
