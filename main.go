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
	router.Static("/static", "./templates/static")
	router.LoadHTMLFiles("templates/login.tmpl")
	router.LoadHTMLFiles("templates/index.tmpl")
	// 登录页面
	router.GET("/login", api.LoginPage)
	router.GET("/index", func(context *gin.Context) {
		context.HTML(200, "index.tmpl", gin.H{
			"role": 1,
		})
	})
	// 用户路由
	user := router.Group("/user")
	{
		user.POST("/add", api.UserAdd)
		user.POST("/delete", api.UserDelete)
		user.POST("/update", api.UserUpdate)
		user.POST("/query", api.UserQuery)
		user.POST("/userlogin", api.UserLogin)
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
	err := router.Run(":8000")
	if err != nil {
		return
	}
}
