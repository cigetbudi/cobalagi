package db

import (
	"cobalagi/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	conf := config.GetConfig()

	dsn := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME + "?charset=utf8&parseTime=true&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})

	// dsn := "host=" + conf.DB_HOST + " user=" + conf.DB_USERNAME + " password=" + conf.DB_PASSWORD + " dbname=" + conf.DB_NAME + " port=" + conf.DB_PORT + " sslmode=disable TimeZone=Asia/Shanghai"
	// DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Cannot connect database")
	}
	DB.AutoMigrate()
}
