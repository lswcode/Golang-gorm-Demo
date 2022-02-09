package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	Id   int
	Name string
	Age  int
}
type Test struct {
	Id   int
	Name string
	City string
}

func main() {
	// 第一个参数是连接的数据库类型，第二个参数依次是用户名，密码，数据库名，后面的使用默认的即可
	// 返回两个参数，第二个参数就是连接失败时返回的错误
	db, dbErr := gorm.Open("mysql", "root:Lswmysql123.@/gorm_project?charset=utf8&parseTime=True&loc=Local")
	if dbErr != nil {
		panic(dbErr) // 发送错误时报错并终止程序
	}
	defer db.Close() // 关闭连接
	// defer用来延迟执行代码，它后面的语句会延迟到函数结束时再调用

	// 创建数据库表格

	db.CreateTable(&Test{}) // 将Text作为model模型，表名会自动加上s

	db.Table("user").CreateTable(&User{}) //使用User结构体作为模型来来创建一个表格，Table中可以自定义表名

	// 1 使用表名删除表`user`
	db.DropTable("tests")

	// 2 使用模型删除表
	// db.DropTable(&Test{})

	// ----检查表名是否存在-------------------------------
	exist := db.HasTable("tests") // 1 使用表名

	// exist := db.HasTable(&Person{})   // 2 使用模型

	fmt.Println(exist)

	// --自动迁移----------------------------------------
	// 创建表格时推荐使用自动迁移，自动迁移会创建表，添加列/索引，但无法改变现有列的类型，也无法删除列，目的是保护数据
	db.AutoMigrate(&Test{}, &User{}) // 可以用逗号隔开，创建多个表格

	// ---增删改查-------------------------------------
	// ------增----------------
	xiaobai := User{Id: 1, Name: "xiaobai", Age: 18}
	xiaolan := User{Id: 2, Name: "lan", Age: 15}
	xiaohei := User{Id: 3, Name: "hei", Age: 11}
	db.Create(&xiaobai)
	db.Create(&xiaolan)
	db.Create(&xiaohei)

	// ------查----------------
	var xb User
	db.First(&xb, 1) // 查询第一条符合条件的数据，第一个参数接收查询到的数据，第二个参数就是条件，默认为模型的第一个字段
	// 相当于Sql语句 select * from users order by id limit 1
	fmt.Println(xb)
	db.First(&xb, "name=?", "xiaobai") // 查询指定的字段
	fmt.Println(xb)

	// 查询所有数据
	db.Find(&xb) // select * from users;
	fmt.Println(xb)

	// -----改-------需要先查到数据，然后才能修改
	var user User
	db.First(&user, 1)                 // 默认使用id字段
	db.Model(&user).Update("age", 111) // 这里的Model就是查询出来的结构体对象user，不是模型

	// 使用`map`可以同时修改多个属性
	var user2 User
	db.First(&user2, 2)
	db.Model(&user2).Updates(map[string]interface{}{"name": "小白", "age": 77}) // 注意id字段时不能修改的
	// update users set name='hello', age=18, actived=false, updated_at='2013-11-17 21:34:10' where id=111;

	// // 使用`struct`也可以更新多个属性
	// db.Model(&user2).Updates(User{Name: "修改", Age: 77})

	// ---删------也要先查找到数据
	var user3 User
	db.First(&user3, 3) // 默认使用id字段
	db.Delete(&user3)

	// ---自定义表名，可以给表名加前缀---------------------------------------
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string { // 第二个参数是新建表时的默认表名
		return "sys_" + defaultTableName
	}

	db.CreateTable(&Test{}) // 将Text作为model模型，表名会自动加上s

	// --gorm提供了一个模板模型gorm.Model-------------------------------------------
	type Boy struct {
		gorm.Model // 里面有 `Id`, `CreatedAt`, `UpdatedAt`, `DeletedAt`这几个字段
		Name       string
		City       string `gorm:"-"` // 在gorm中忽略该字段
		// 其它常用标签属性:
		// `gorm:"primary_key"` 设置主键
		// `gorm:"not null"` 规定不为空
		// `gorm:"index"`    索引，推荐设置索引，可以加快数据查询速度
		// `gorm:"column:user_name"` 指定列名
	}

	db.Debug().CreateTable(&Boy{}) // 在gorm语句中调用Debug()可以打印出对应的Sql语句
}
