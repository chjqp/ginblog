package model

import (
	"ginblog/utils/errmsg"

	"github.com/jinzhu/gorm"
)


type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:varchar(20);not  null" json:"password"`
	Role     int    `gorm:"type:int" json:"role"` //	1为管理员，2为普通用户
}
//查询用户是否存在
func CheckUser(name string) (code int) {
	var users User
	db.Select("id").Where("username = ?", name).First(&users)
	if users.ID > 0 {
		return errmsg.ERROR_USERNAME_USED //1001
	}
	return errmsg.SUCCESS //200
}

func CreateUser(data *User) int{
	err:=db.Create(&data).Error
	if err!=nil{
		return errmsg.ERROR//500
	}
	return errmsg.SUCCESS//200
}