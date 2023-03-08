package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
	Name     string //姓名
	Email    string //邮箱
	Mobile   string //手机
	QQ       string
	Ding     string     //钉钉
	Gender   int        //0男 1女
	Birthday *time.Time //生日
	Remark   string     //备注
	Token    string     `gorm:"-"`
	Session  string     `gorm:"-"`
}
