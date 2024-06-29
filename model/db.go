package model

import (
	"fmt"
	"ginblog/utils"
	"github.com/jinzhu/gorm"
	 _"github.com/go-sql-driver/mysql"
	"time"
)
var db *gorm.DB
var err error
func InitDb(){
   db,err=gorm.Open(utils.Db,fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",utils.DbUser,utils.DbPassWord,utils.DbHost,utils.DbPort,utils.DbName))
	 if err!=nil{
		 fmt.Println("连接数据库失败，请检查参数：",err)
	 }

	 db.SingularTable(true)
	 
	 db.AutoMigrate(&User{},&Article{},&Category{})

	 db.DB().SetMaxIdleConns(10)
	 db.DB().SetMaxOpenConns(100)
	 db.DB().SetConnMaxLifetime(time.Second*10)


	 //db.Close()

}