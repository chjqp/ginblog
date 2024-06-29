package model   


import (
	"github.com/jinzhu/gorm"

	
)
type Article struct{
	Category Category
	gorm.Model
	Title string `gorm:"type:varchar(100);not null" json:"title"`
	 
	Cid int `gorm:"type:int;not null" json:"cid"`
	Desc string `gorm:"type:varchar(200);not null" json:"desc"`
	Contenet string `gorm:"type:longtext;not null" json:"content"`
	Img string `gorm:"type:varchar(100);not null" json:"img"`
	


}