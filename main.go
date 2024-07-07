package main

import (
	"github.com/gin-gonic/gin"
	"notice_demo/api"
)

func main() {
	//tools.Initdb()
	// 强制日志颜色化
	gin.ForceConsoleColor()
	// 初始化路由
	router := gin.Default()
	// 用户路由
	user := router.Group("/user")
	{
		user.POST("/add", api.UserAdd)
		user.POST("/delete", api.UserDelete)
		user.POST("/update", api.UserUpdate)
		user.POST("/query", api.UserQuery)
	}
	// 机器人路由
	robot := router.Group("/robot")
	{
		robot.POST("/add", api.RobotAdd)
		robot.POST("/delete", api.RobotDelete)
		robot.POST("/update", api.RobotUpdate)
		robot.POST("/query", api.RobotQuery)
	}

	// 应用路由路由
	application := router.Group("/application")
	{
		application.POST("/add", api.ApplicationAdd)
		application.POST("/delete", api.ApplicationDelete)
		application.POST("/update", api.ApplicationUpdate)
		application.POST("/query", api.ApplicationQuery)
	}

	// 邮件路由
	email := router.Group("/email")
	{
		email.POST("/add", api.EmailAdd)
		email.POST("/delete", api.EmailDelete)
		email.POST("/update", api.EmailUpdate)
		email.POST("/query", api.EmailQuery)
	}

	// 运行服务
	router.Run(":8000")
}
