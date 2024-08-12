package model

import (
	"ginblog/utils/errmsg"
	"log"

	"github.com/jinzhu/gorm"
)

/*
逻辑外键（Cid）：在你的Article结构体中，Cid字段作为逻辑外键，指向Category表的某个记录的主键。
预加载（Preload）：通过预加载Category，GORM会自动查询每篇文章对应的分类信息，并将其填充到Article结构体的Category字段中。
分页查询：代码中的Limit和Offset方法用于实现分页功能，确保每次只查询和返回指定数量的文章记录。
总的来说，虽然你的代码中没有直接操作物理外键，但通过GORM的foreignkey标签和Preload方法，它实现了类似外键的数据关联和完整性保护的逻辑。
*/
type Article struct {
	Category Category `gorm:"foreignkey:Cid"`
	gorm.Model
	Title string `gorm:"type:varchar(100);not null" json:"title"`

	Cid     int    `gorm:"type:int;not null" json:"cid"`
	Desc    string `gorm:"type:varchar(200);not null" json:"desc"`
	Content string `gorm:"type:longtext;not null" json:"content"`
	Img     string `gorm:"type:varchar(100);not null" json:"img"`
}

// type Article struct {
// 	Category Category `gorm:"foreignkey:Cid"`
// 	gorm.Model
// 	Title        string `gorm:"type:varchar(100);not null" json:"title"`
// 	Cid          int    `gorm:"type:int;not null" json:"cid"`
// 	Desc         string `gorm:"type:varchar(200)" json:"desc"`
// 	Content      string `gorm:"type:longtext" json:"content"`
// 	Img          string `gorm:"type:varchar(100)" json:"img"`
// 	CommentCount int    `gorm:"type:int;not null;default:0" json:"comment_count"`
// 	ReadCount    int    `gorm:"type:int;not null;default:0" json:"read_count"`
// }

// func Base64image() {

// 	file, err := os.Open("ginblog/images/goods12.png")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()

// 	// 解码原始图片文件
// 	img, _, err := image.Decode(file)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// 压缩图片到指定大小
// 	newWidth := 200
// 	newHeight := 0 // 0 表示按比例缩放
// 	resizedImg := resize.Resize(uint(newWidth), uint(newHeight), img, resize.Lanczos3)

// 	// 创建压缩后的图片文件
// 	out, err := os.Create("compressed.jpg")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer out.Close()

// 	// 将压缩后的图片写入文件
// 	jpeg.Encode(out, resizedImg, nil)

// 	log.Println("图片压缩完成")

// }

// 新增文章
func CreateArt(data *Article) int {
	//data.Password=ScryptPw(data.Password)

	err := db.Create(&data).Error
	if err != nil {
		log.Println(err)
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS //200
}

//查询单个文章

// 查询文章列表
func GetArt(pageSize int, pageNum int) ([]Article, int, int) {
	var artlist []Article
	var total int
	err = db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&artlist).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR, 0
	}
	return artlist, errmsg.SUCCESS, total
}

// 编辑用户
func EditArt(id int, data *Article) int {
	var art Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img

	err = db.Model(&art).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		log.Println(err)
		return errmsg.ERROR //500

	}
	return errmsg.SUCCESS //200
}

// 查询单个分类下的所有文章
func GetCateArt(id int, pageSize int, pageNum int) ([]Article, int, int) {
	var cateArtlist []Article
	var total int
	err = db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("cid=?", id).Find(&cateArtlist).Count(&total).Error
	if err != nil {
		return nil, errmsg.ERROR_CATEGORY_NOT_EXIST, 0

	}
	return cateArtlist, errmsg.SUCCESS, total
}

func GetArtInfo(id int) (Article, int) {

	var art Article
	err := db.Preload("Category").Where("id = ?", id).First(&art)
	if err != nil {
		return art, errmsg.ERROR
	}
	return art, errmsg.SUCCESS

}

// 删除分类
func DeleteArt(id int) int {
	var art Article
	err = db.Where("id = ?", id).Delete(&art).Error
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS //200
}
