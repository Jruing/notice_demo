package api

import "github.com/gin-gonic/gin"

func ApplicationAdd(c *gin.Context) {
	c.JSON(200, gin.H{
		"Code": 200,
		"Msg":  "添加用户成功",
	})
}

func ApplicationDelete(c *gin.Context) {
	c.JSON(200, gin.H{
		"Code": 200,
		"Msg":  "添加用户成功",
	})
}

func ApplicationQuery(c *gin.Context) {
	c.JSON(200, gin.H{
		"Code": 200,
		"Msg":  "添加用户成功",
	})
}

func ApplicationUpdate(c *gin.Context) {
	c.JSON(200, gin.H{
		"Code": 200,
		"Msg":  "添加用户成功",
	})
}
