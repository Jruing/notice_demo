package models

import "time"

type Application struct {
	Id      int64     `xorm:"int(32) notnull pk autoincr unique 'id' comment('id')"`
	AppID   string    `xorm:"varchar(25) notnull unique 'appid' comment('应用ID')"`
	AppName string    `xorm:"varchar(25) notnull 'appname' comment('应用名称')"`
	AppType int64     `xorm:"int(32) notnull 'apptype' comment('用户类型,1:钉钉 2:企业微信')"`
	Remarks string    `xorm:"varchar(255) 'remarks' comment('备注')"`
	Created time.Time `xorm:"created"`
}
