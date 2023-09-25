package logic

import (
	"bluebell/dao/mysql"
	"bluebell/pkg/snowflake"
)

func SignUp() {
	//1.判断用户存在
	mysql.QueryUserByName()

	//2.生成uid
	snowflake.GenID()
	//密码加密

	//2.保存进数据库
	mysql.InsertUser()
}
