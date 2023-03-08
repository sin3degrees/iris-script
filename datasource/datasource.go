package datasource

import (
	"errors"
	"iris-script/conf"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func init() {
	path := ""
	dbType := conf.Sysconfig.DB.Type
	switch dbType {
	case "mysql":
		path = strings.Join([]string{conf.Sysconfig.DB.UserName, ":", conf.Sysconfig.DB.Password, "@(", conf.Sysconfig.DB.Ip,
			":", conf.Sysconfig.DB.Port, ")/", conf.Sysconfig.DB.DBName, "?charset=utf8&parseTime=true"}, "")
	case "postgres":
		path = strings.Join([]string{"host=", conf.Sysconfig.DB.Ip, " port=", conf.Sysconfig.DB.Port, " user=", conf.Sysconfig.DB.UserName,
			" dbname=", conf.Sysconfig.DB.DBName, " password=", conf.Sysconfig.DB.Password, " sslmode=disable"}, "")
	case "mssql":
		path = strings.Join([]string{"server=", conf.Sysconfig.DB.Ip, ";port=", conf.Sysconfig.DB.Port, ";database=",
			conf.Sysconfig.DB.DBName, ";user id=", conf.Sysconfig.DB.UserName, ";password=", conf.Sysconfig.DB.Password}, "")
	case "sqlite3":
		path = conf.Sysconfig.DB.DBName
	default:
		panic(errors.New("不支持的数据库类型"))
	}
	var err error
	db, err = gorm.Open(dbType, path)
	if err != nil {
		panic(err)
	}
	db.DB().SetConnMaxLifetime(time.Duration(conf.Sysconfig.DB.MaxLife) * time.Second)
	db.DB().SetMaxIdleConns(conf.Sysconfig.DB.MaxIdle) //最大打开的连接数
	db.DB().SetMaxOpenConns(conf.Sysconfig.DB.MaxOpen) //设置最大闲置个数
	db.SingularTable(true)                             //表生成结尾不带s
	// 启用Logger，显示详细日志
	db.LogMode(true)
	Createtable()
}
