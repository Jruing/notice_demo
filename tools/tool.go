package tools

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"notice_demo/models"
	"xorm.io/xorm"
)

func Initdb() *xorm.Engine {
	// 初始化pg数据库连接
	pgEngine, err := xorm.NewEngine("mysql", "root:test1234@/notice_demo?charset=utf8")
	if err != nil {
		panic(err)
	}
	err = pgEngine.Sync(new(models.User))
	if err != nil {
		fmt.Println(err)
		panic("用户表同步失败")
	}
	err = pgEngine.Sync(new(models.Robot))
	if err != nil {
		fmt.Println(err)
		panic("用户表同步失败")
	}

	err = pgEngine.Sync(new(models.Application))
	if err != nil {
		fmt.Println(err)
		panic("用户表同步失败")
	}

	return pgEngine
}
