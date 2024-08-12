package model

import (
	"fmt"
	"ginblog/utils"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB //= InitDb()
var err error = nil

// func InitDb() {
// 	db, err = gorm.Open("mysql", "root:root@tcp(localhost:3306)/ginblog?charset=utf8&parseTime=True&loc=Local")
// 	if err != nil {
// 		fmt.Println("连接数据库失败，请检查参数：", err)
// 		return
// 	}

// 	// 使用类型而不是变量实例进行自动迁移
// 	if err := db.Debug().AutoMigrate(&User{}, &Article{}, &Category{}); err != nil {
// 		fmt.Println("自动迁移表失败：", err)
// 	}
// }

func InitDb() {
	// "root:root@tcp(localhost:3306)/ginblog?charset=utf8&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", utils.DbUser, utils.DbPassWord, utils.DbHost, utils.DbPort, utils.DbName)
	log.Println(dsn)
	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println("连接数据库失败，请检查参数：", err)
		return
	}
	// db.SingularTable(true)

	// db.Debug().AutoMigrate(&User{}, &Article{}, &Category{})
	result := db.AutoMigrate(&User{}, &Article{}, &Category{})
	if result.Error != nil {
		fmt.Println("自动迁移表失败：", result.Error)
		return
	}
	//db.AutoMigrate(&User{}, &Article{}, &Category{})
	// , &Article{}, &Category{}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Second * 10)
	// db.Create(&User{Username: "admin", Password: "admin", Role: 1})
	//	return db
	//db.Close()

}
