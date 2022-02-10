package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB
var err error

func init() {
	// 第一个参数是连接的数据库类型，第二个参数依次是用户名，密码，数据库名，后面的使用默认的即可
	// 返回两个参数，第二个参数就是连接失败时返回的错误
	Db, err = gorm.Open("mysql", "root:Lswmysql123.@/gorm_project?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err) // 发送错误时报错并终止程序
	}
	defer Db.Close() // 关闭连接
	// defer用来延迟执行代码，它后面的语句会延迟到函数结束时再调用

	// Db.LogMode(true) // 开启日志记录，会打印所有gorm对应的sql语句，可以在配置文件中谁知

	Db.DB().SetMaxOpenConns(100) // 设置数据库的最大连接数
	Db.DB().SetMaxIdleConns(50)  // 设置数据库的最大空闲数

}
