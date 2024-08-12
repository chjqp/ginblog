package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//添加分类

//查询分类是否存在

//查询单个分类下的文章

//查询分类列表

//编辑分类

//删除分类

// 添加用户

func AddCategory(c *gin.Context) {
	var data model.Category
	// if err := c.ShouldBindJSON(&data); err != nil {
	// 	// 处理错误

	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	_ = c.ShouldBindJSON(&data)
	code := model.CheckCategory(data.Name)
	if code == errmsg.SUCCESS {
		model.CreateCate(&data)

	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询分类下的所有分类列表
func GetCate(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data, total := model.GetCate(pageSize, pageNum)
	code := errmsg.SUCCESS // 修正这里
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// 编辑分类名
func EditCate(c *gin.Context) {
	var data model.Category
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON((&data))
	code := model.CheckCategory(data.Name)
	if code == errmsg.SUCCESS {
		model.EditCate(id, &data)
	}
	if code == errmsg.ERROR_CATENAME_USED {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code)})
}

// 删除用户
func DeleteCate(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DeleteCate(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
