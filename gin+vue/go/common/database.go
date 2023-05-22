package common

import (
	"fmt"
	"gin/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


func InitDB() *gorm.DB {
	//driverName := "mysql"
	host := "127.0.0.1"
	port := "3306"
	username := "root"
	password := "root"
	DBName := "gin_demo"
	charset := "utf8"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True",
		username,
		password,
		host,
		port,
		DBName,
		charset,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.User{})
	return db
}
