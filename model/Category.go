package model

import (
	"ginblog/utils/errmsg"
	"log"

	"github.com/jinzhu/gorm"
)

type Category struct {
	//gorm.Model
	ID uint `gorm:"primary_key;auto_increment" json:"id"`

	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

// r=gin.Default()
// 查询分类是否存在
func CheckCategory(name string) (code int) {
	var cate Category
	if db == nil {
		log.Println("Database connection is not initialized.")
		return errmsg.ERROR // 或者是一个更具体的错误代码
	}

	db.Select("id").Where("name = ?", name).First(&cate)
	if cate.ID > 0 {
		return errmsg.ERROR_CATENAME_USED //1001
	}
	return errmsg.SUCCESS //200

}

// 新增分类
func CreateCate(data *Category) int {
	//data.Password=ScryptPw(data.Password)

	err := db.Create(&data).Error
	if err != nil {
		log.Println(err)
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS //200
}
func GetCate(pageSize int, pageNum int) ([]Category, int) {
	var cate []Category
	var total int
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cate).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return cate, total
}

// 编辑用户
func EditCate(id int, data *Category) int {
	var cate Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name

	err = db.Model(&cate).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS //200
}

// 删除分类
func DeleteCate(id int) int {
	var cate Category
	err = db.Where("id = ?", id).Delete(&cate).Error
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS //200
}
