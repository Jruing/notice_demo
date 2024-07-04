package api

import "github.com/gin-gonic/gin"

func RobotAdd(c *gin.Context) {
	c.JSON(200, gin.H{
		"Code": 200,
		"Msg":  "添加用户成功",
	})
}

func RobotDelete(c *gin.Context) {
	c.JSON(200, gin.H{
		"Code": 200,
		"Msg":  "添加用户成功",
	})
}

func RobotQuery(c *gin.Context) {
	c.JSON(200, gin.H{
		"Code": 200,
		"Msg":  "添加用户成功",
	})
}

func RobotUpdate(c *gin.Context) {
	c.JSON(200, gin.H{
		"Code": 200,
		"Msg":  "添加用户成功",
	})
}
