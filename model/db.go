package model

import (
	"fmt"
	"time"

	_"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB //= InitDb()
var err error = nil

func InitDb() {
	// fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", utils.DbUser, utils.DbPassWord, utils.DbHost, utils.DbPort, utils.DbName)
	db, err = gorm.Open("mysql", "root:root@tcp(localhost:3306)/ginblog?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("连接数据库失败，请检查参数：", err)
		return
	}

	// db.SingularTable(true)
	user := User{}
	Article := Article{}
	Category := Category{}
	if err := db.AutoMigrate(&user, &Article, &Category); err != nil {
		fmt.Println("自动迁移表失败：", err)
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

// 改成db,err=gorm.Open(umysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",utils.DbUser,utils.DbPassWord,utils.DbHost,utils.DbPort,utils.DbName)))
// 最后把db return 了就行
