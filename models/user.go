package models

import "time"

type User struct {
	Id       int64     `xorm:"int(32) notnull unique 'id' comment('用户id')"`
	UserName string    `xorm:"varchar(25) notnull unique 'username' comment('姓名')"`
	PassWord string    `xorm:"varchar(25) notnull 'passwd' comment('密码')"`
	UserType int64     `xorm:"int(32) notnull 'usertype' comment('用户类型')"`
	Remarks  string    `xorm:"varchar(255) 'remarks' comment('备注')"`
	Created  time.Time `xorm:"created"`
}
