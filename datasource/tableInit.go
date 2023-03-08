package datasource

import (
	"iris-script/models"
)

// 初始化表 如果不存在该表 则自动创建
func Createtable() {
	GetDB().AutoMigrate(
		&models.User{},
		&models.Book{},
	)
}
