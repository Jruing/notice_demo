package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"notice_demo/models"
)

func ApplicationAdd(c *gin.Context) {
	log_uuid := uuid.New()
	session := eng.NewSession()
	defer session.Close()
	json := make(map[string]interface{})
	err := c.BindJSON(&json)
	logger.Info("应用新增函数", zap.String("uuid", log_uuid.String()), zap.String("request_data", fmt.Sprintf("%+v", json)))

	if err != nil {
		logger.Error("请求报文解析失败", zap.String("uuid", log_uuid.String()), zap.String("error", err.Error()))
		c.JSON(200, gin.H{
			"Code": -1,
			"Msg":  "数据解析失败",
		})
	}
	app := new(models.Application)
	app.AppID = json["appid"].(string)
	app.AppName = json["appname"].(string)
	app.AppType = int64(json["apptype"].(float64))
	app.Remarks = json["remarks"].(string)
	_, err = session.Insert(app)
	if err != nil {
		logger.Error("应用新增失败", zap.String("uuid", log_uuid.String()), zap.String("error", err.Error()))
		c.JSON(200, gin.H{
			"Code": 0,
			"Msg":  "应用新增失败",
		})
	} else {
		logger.Info("应用新增成功", zap.String("uuid", log_uuid.String()))
		c.JSON(200, gin.H{
			"Code": 1,
			"Msg":  "应用新增成功",
		})
	}

}

func ApplicationDelete(c *gin.Context) {
	log_uuid := uuid.New()
	session := eng.NewSession()
	defer session.Close()
	json := make(map[string]interface{})
	err := c.BindJSON(&json)
	logger.Info("应用删除函数", zap.String("uuid", log_uuid.String()), zap.String("request_data", fmt.Sprintf("%+v", json)))

	if err != nil {
		logger.Error("请求报文解析失败", zap.String("uuid", log_uuid.String()), zap.String("error", err.Error()))
		c.JSON(200, gin.H{
			"Code": -1,
			"Msg":  "数据解析失败",
		})
	}
	id := int64(json["id"].(float64))
	app := new(models.Application)
	_, err = session.Where("id = ?", id).Delete(app)
	if err != nil {
		logger.Error("应用删除失败", zap.String("uuid", log_uuid.String()), zap.String("error", err.Error()))
		c.JSON(200, gin.H{
			"Code": 0,
			"Msg":  "应用删除失败",
		})
	} else {
		logger.Info("应用删除成功", zap.String("uuid", log_uuid.String()))
		c.JSON(200, gin.H{
			"Code": 1,
			"Msg":  "应用删除成功",
		})
	}
}

func ApplicationQuery(c *gin.Context) {
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
	app := []models.Application{}
	page := int(json["page"].(float64))
	pageSize := int(json["pageSize"].(float64))
	offset := (page - 1) * pageSize
	err = session.Limit(pageSize, offset).Find(&app)
	if err != nil {
		logger.Error("应用查询失败", zap.String("error", err.Error()))
		c.JSON(200, gin.H{
			"Code": -2,
			"Msg":  "应用查询失败",
		})
	} else {
		logger.Info("应用查询成功", zap.String("uuid", log_uuid.String()))
		c.JSON(200, gin.H{
			"Code": 200,
			"Msg":  "应用查询成功",
			"Data": app,
		})
	}
}

func ApplicationUpdate(c *gin.Context) {
	log_uuid := uuid.New()
	session := eng.NewSession()
	defer session.Close()
	json := make(map[string]interface{})
	err := c.BindJSON(&json)
	logger.Info("应用修改函数", zap.String("uuid", log_uuid.String()), zap.String("request_data", fmt.Sprintf("%+v", json)))
	if err != nil {
		logger.Error("请求报文解析失败", zap.String("uuid", log_uuid.String()), zap.String("error", err.Error()))
		c.JSON(200, gin.H{
			"Code": -1,
			"Msg":  "数据解析失败",
		})
	}
	app := new(models.Application)
	app.Id = int64(json["id"].(float64))
	app.AppName = json["appname"].(string)
	app.AppID = json["appid"].(string)
	app.Remarks = json["remarks"].(string)
	app.AppType = int64(json["apptype"].(float64))
	_, err = session.Where("id = ?", app.Id).Update(app)
	if err != nil {
		logger.Error("应用修改失败", zap.String("uuid", log_uuid.String()), zap.String("error", err.Error()))
		c.JSON(200, gin.H{
			"Code": -2,
			"Msg":  "应用修改失败",
		})
	} else {
		logger.Info("应用修改成功", zap.String("uuid", log_uuid.String()))
		c.JSON(200, gin.H{
			"Code": 200,
			"Msg":  "应用修改成功",
		})
	}
}
