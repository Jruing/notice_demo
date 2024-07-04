package api

import "github.com/gin-gonic/gin"

func EmailAdd(c *gin.Context) {
	c.JSON(200, gin.H{
		"Code": 200,
		"Msg":  "添加用户成功",
	})
}

func EmailDelete(c *gin.Context) {
	c.JSON(200, gin.H{
		"Code": 200,
		"Msg":  "添加用户成功",
	})
}

func EmailQuery(c *gin.Context) {
	c.JSON(200, gin.H{
		"Code": 200,
		"Msg":  "添加用户成功",
	})
}

func EmailUpdate(c *gin.Context) {
	c.JSON(200, gin.H{
		"Code": 200,
		"Msg":  "添加用户成功",
	})
}
