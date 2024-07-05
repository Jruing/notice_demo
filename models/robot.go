package models

import "time"

type Robot struct {
	Id        int64     `xorm:"int(32) notnull pk autoincr unique 'id' comment('id')"`
	RobotAddr string    `xorm:"varchar(255) notnull 'robotaddr' comment('机器人地址')"`
	RobotName string    `xorm:"varchar(25) notnull 'robotname' comment('机器人名称')"`
	RobotType int64     `xorm:"int(32) notnull 'robottype' comment('机器人类型：1:钉钉 2:企业微信')"`
	Remarks   string    `xorm:"varchar(255) 'remarks' comment('备注')"`
	Created   time.Time `xorm:"created"`
}
