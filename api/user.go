package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"notice_demo/models"
	"notice_demo/tools"
)

var (
	eng    = tools.Initdb()
	logger = tools.InitLog("notice_info.log", 0, 0, 0)
)

// 用户新增
func UserAdd(c *gin.Context) {
	log_uuid := uuid.New()
	session := eng.NewSession()
	defer session.Close()
	json := make(map[string]interface{})
	err := c.BindJSON(&json)
	logger.Info("用户新增函数", zap.String("uuid", log_uuid.String()), zap.String("request_data", fmt.Sprintf("%+v", json)))

	if err != nil {
		logger.Error("请求报文解析失败", zap.String("uuid", log_uuid.String()), zap.String("error", err.Error()))
		c.JSON(200, gin.H{
			"Code": -1,
			"Msg":  "数据解析失败",
		})
	}
	user := new(models.User)
	user.UserName = json["username"].(string)
	user.PassWord = json["passwd"].(string)
	user.Remarks = json["remarks"].(string)
	user.Id = int64(json["id"].(float64))
	user.UserType = int64(json["usertype"].(float64))
	_, err = session.Insert(user)
	if err != nil {
		logger.Error("用户添加失败", zap.String("uuid", log_uuid.String()), zap.String("error", err.Error()))
		c.JSON(200, gin.H{
			"Code": 0,
			"Msg":  "用户添加失败",
		})
	} else {
		logger.Info("用户添加成功", zap.String("uuid", log_uuid.String()))
		c.JSON(200, gin.H{
			"Code": 1,
			"Msg":  "用户添加成功",
		})
	}

}

// 用户删除
func UserDelete(c *gin.Context) {
	log_uuid := uuid.New()
	session := eng.NewSession()
	defer session.Close()
	json := make(map[string]interface{})
	err := c.BindJSON(&json)
	logger.Info("用户删除函数", zap.String("uuid", log_uuid.String()), zap.String("request_data", fmt.Sprintf("%+v", json)))

	if err != nil {
		logger.Error("请求报文解析失败", zap.String("uuid", log_uuid.String()), zap.String("error", err.Error()))
		c.JSON(200, gin.H{
			"Code": -1,
			"Msg":  "数据解析失败",
		})
	}
	id := int64(json["id"].(float64))
	user := new(models.User)
	_, err = session.Where("id = ?", id).Delete(user)
	if err != nil {
		logger.Error("删除失败", zap.String("uuid", log_uuid.String()), zap.String("error", err.Error()))
		c.JSON(200, gin.H{
			"Code": 0,
			"Msg":  "删除用户失败",
		})
	} else {
		logger.Info("删除用户成功", zap.String("uuid", log_uuid.String()))
		c.JSON(200, gin.H{
			"Code": 1,
			"Msg":  "删除用户成功",
		})
	}
}

// 用户查询
func UserQuery(c *gin.Context) {
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
	session = session.Where("1=1")
	for key, value := range json {
		if value != nil && key == "id" {
			session = session.Where("id = ?", value)
		}
	}
	user := []models.User{}
	page := int(json["page"].(float64))
	pageSize := int(json["pageSize"].(float64))
	offset := (page - 1) * pageSize
	err = session.Limit(pageSize, offset).Find(&user)
	if err != nil {
		logger.Error("用户查询失败", zap.String("error", err.Error()))
		c.JSON(200, gin.H{
			"Code": -2,
			"Msg":  "用户数据查询失败",
		})
	} else {
		logger.Info("用户查询成功", zap.String("uuid", log_uuid.String()))
		c.JSON(200, gin.H{
			"Code": 200,
			"Msg":  "用户数据查询成功",
			"Data": user,
		})
	}
}

// 用户修改
func UserUpdate(c *gin.Context) {
	log_uuid := uuid.New()
	session := eng.NewSession()
	defer session.Close()
	json := make(map[string]interface{})
	err := c.BindJSON(&json)
	logger.Info("用户修改函数", zap.String("uuid", log_uuid.String()), zap.String("request_data", fmt.Sprintf("%+v", json)))
	if err != nil {
		logger.Error("请求报文解析失败", zap.String("uuid", log_uuid.String()), zap.String("error", err.Error()))
		c.JSON(200, gin.H{
			"Code": -1,
			"Msg":  "数据解析失败",
		})
	}
	user := new(models.User)
	user.Id = int64(json["id"].(float64))
	user.UserName = json["username"].(string)
	user.PassWord = json["passwd"].(string)
	user.Remarks = json["remarks"].(string)
	user.Id = int64(json["id"].(float64))
	user.UserType = int64(json["usertype"].(float64))
	_, err = session.Where("id = ?", user.Id).Update(user)
	if err != nil {
		logger.Error("用户修改失败", zap.String("uuid", log_uuid.String()), zap.String("error", err.Error()))
		c.JSON(200, gin.H{
			"Code": -2,
			"Msg":  "修改用户失败",
		})
	} else {
		logger.Info("用户修改成功", zap.String("uuid", log_uuid.String()))
		c.JSON(200, gin.H{
			"Code": 200,
			"Msg":  "修改用户成功",
		})
	}
}
