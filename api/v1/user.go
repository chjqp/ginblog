package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"net/http"

	"github.com/gin-gonic/gin"
)

//查询用户是否存在
func UserExist(c *gin.Context){

}

//添加用户
func AddUser(c *gin.Context){
	var data model.User
	// if err := c.ShouldBindJSON(&data); err != nil {
	// 	// 处理错误

	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	_=c.ShouldBindJSON(&data)
	code:=model.CheckUser(data.Username)
	if code==errmsg.SUCCESS{
		model.CreateUser(&data)
	

	}
	if code==errmsg.ERROR_USERNAME_USED{
		code=errmsg.ERROR_USERNAME_USED
	}
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"message":errmsg.GetErrMsg(code),
	})
}


//查询单个用户

//查询用户列表
func GetUsers(c *gin.Context){

}
//编辑用户        
func EditUser(c *gin.Context){

}
//删除用户
func DeleteUser(c *gin.Context){

}


