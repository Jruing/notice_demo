package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"net/http"
	"notice_demo/models"
)

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tmpl", gin.H{})
}
func UserLogin(c *gin.Context) {
	log_uuid := uuid.New()
	json := make(map[string]interface{})
	err := c.BindJSON(&json)
	logger.Info("用户查询函数", zap.String("uuid", log_uuid.String()), zap.String("request_data", fmt.Sprintf("%+v", json)))

	if err != nil {
		logger.Error("请求报文解析失败", zap.String("uuid", log_uuid.String()), zap.String("error", err.Error()))
		c.JSON(200, gin.H{
			"Code": -1,
			"Msg":  "数据解析失败",
		})
	}
	session := eng.NewSession()
	defer session.Close()
	user := []models.User{}
	session = session.Where("1=1")
	err = session.Where("username=? and passwd=?", json["username"], json["passwd"]).Find(&user)
	if err != nil {
		logger.Error("用户登录失败", zap.String("uuid", log_uuid.String()), zap.String("error", err.Error()))
		c.JSON(200, gin.H{
			"Code": -2,
			"Msg":  "用户登录失败",
		})
	} else {
		fmt.Println(user, len(user))
		if len(user) == 0 {
			logger.Error("用户登录失败", zap.String("uuid", log_uuid.String()))
			c.JSON(200, gin.H{
				"Code": -2,
				"Msg":  "用户密码错误",
			})
		} else {
			logger.Info("用户查询成功", zap.String("uuid", log_uuid.String()))
			c.JSON(200, gin.H{
				"Code": 200,
				"Msg":  "用户登录成功",
			})
		}

	}
}
