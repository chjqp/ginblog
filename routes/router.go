package routes

import (
	"ginblog/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	router := r.Group("api/v1")
	{
		router.GET("hello", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"msg": "ok",
			})
		})
	}

	r.Run(utils.HttpPort)
}