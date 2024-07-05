package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"notice_demo/models"
)

func RobotAdd(c *gin.Context) {
	log_uuid := uuid.New()
	session := eng.NewSession()
	defer session.Close()
	json := make(map[string]interface{})
	err := c.BindJSON(&json)
	logger.Info("机器人新增函数", zap.String("uuid", log_uuid.String()), zap.String("request_data", fmt.Sprintf("%+v", json)))

	if err != nil {
		logger.Error("请求报文解析失败", zap.String("uuid", log_uuid.String()), zap.String("error", err.Error()))
		c.JSON(200, gin.H{
			"Code": -1,
			"Msg":  "数据解析失败",
		})
	}
	robot := new(models.Robot)
	robot.RobotName = json["robotname"].(string)
	robot.RobotAddr = json["robotaddr"].(string)
	robot.RobotType = int64(json["robottype"].(float64))
	robot.Remarks = json["remarks"].(string)
	_, err = session.Insert(robot)
	if err != nil {
		logger.Error("机器人新增失败", zap.String("uuid", log_uuid.String()), zap.String("error", err.Error()))
		c.JSON(200, gin.H{
			"Code": 0,
			"Msg":  "机器人新增失败",
		})
	} else {
		logger.Info("机器人新增成功", zap.String("uuid", log_uuid.String()))
		c.JSON(200, gin.H{
			"Code": 1,
			"Msg":  "机器人新增成功",
		})
	}

}

func RobotDelete(c *gin.Context) {
	log_uuid := uuid.New()
	session := eng.NewSession()
	defer session.Close()
	json := make(map[string]interface{})
	err := c.BindJSON(&json)
	logger.Info("机器人删除函数", zap.String("uuid", log_uuid.String()), zap.String("request_data", fmt.Sprintf("%+v", json)))

	if err != nil {
		logger.Error("请求报文解析失败", zap.String("uuid", log_uuid.String()), zap.String("error", err.Error()))
		c.JSON(200, gin.H{
			"Code": -1,
			"Msg":  "数据解析失败",
		})
	}
	id := int64(json["id"].(float64))
	robot := new(models.Robot)
	_, err = session.Where("id = ?", id).Delete(robot)
	if err != nil {
		logger.Error("机器人删除失败", zap.String("uuid", log_uuid.String()), zap.String("error", err.Error()))
		c.JSON(200, gin.H{
			"Code": 0,
			"Msg":  "机器人删除失败",
		})
	} else {
		logger.Info("机器人删除成功", zap.String("uuid", log_uuid.String()))
		c.JSON(200, gin.H{
			"Code": 1,
			"Msg":  "机器人删除成功",
		})
	}
}

func RobotQuery(c *gin.Context) {
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
	robot := []models.Robot{}
	page := int(json["page"].(float64))
	pageSize := int(json["pageSize"].(float64))
	offset := (page - 1) * pageSize
	err = session.Limit(pageSize, offset).Find(&robot)
	if err != nil {
		logger.Error("机器人查询失败", zap.String("error", err.Error()))
		c.JSON(200, gin.H{
			"Code": -2,
			"Msg":  "机器人查询失败",
		})
	} else {
		logger.Info("机器人查询成功", zap.String("uuid", log_uuid.String()))
		c.JSON(200, gin.H{
			"Code": 200,
			"Msg":  "机器人查询成功",
			"Data": robot,
		})
	}
}

func RobotUpdate(c *gin.Context) {
	log_uuid := uuid.New()
	session := eng.NewSession()
	defer session.Close()
	json := make(map[string]interface{})
	err := c.BindJSON(&json)
	logger.Info("机器人修改函数", zap.String("uuid", log_uuid.String()), zap.String("request_data", fmt.Sprintf("%+v", json)))
	if err != nil {
		logger.Error("请求报文解析失败", zap.String("uuid", log_uuid.String()), zap.String("error", err.Error()))
		c.JSON(200, gin.H{
			"Code": -1,
			"Msg":  "数据解析失败",
		})
	}
	robot := new(models.Robot)
	robot.Id = int64(json["id"].(float64))
	robot.RobotAddr = json["robotaddr"].(string)
	robot.RobotName = json["robotname"].(string)
	robot.Remarks = json["remarks"].(string)
	robot.RobotType = int64(json["robottype"].(float64))
	_, err = session.Where("id = ?", robot.Id).Update(robot)
	if err != nil {
		logger.Error("机器人修改失败", zap.String("uuid", log_uuid.String()), zap.String("error", err.Error()))
		c.JSON(200, gin.H{
			"Code": -2,
			"Msg":  "机器人修改失败",
		})
	} else {
		logger.Info("机器人修改成功", zap.String("uuid", log_uuid.String()))
		c.JSON(200, gin.H{
			"Code": 200,
			"Msg":  "机器人修改成功",
		})
	}
}
