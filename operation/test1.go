package test1

import (
	"fmt"
	"gorm_project/db"
	"gorm_project/models"

	"github.com/jinzhu/gorm"
)

func TestFun1() {
	// 创建数据库表格

	// 导入连接的数据库db.Db
	db.Db.CreateTable(&models.Test{}) // 将Text作为model模型，表名会自动加上s

	db.Db.Table("user").CreateTable(&models.User{}) //使用User结构体作为模型来来创建一个表格，Table中可以自定义表名

	// 1 使用表名删除表`user`
	db.Db.DropTable("tests")

	// 2 使用模型删除表
	// db.Db.DropTable(&Test{})

	// ----检查表名是否存在-------------------------------
	exist := db.Db.HasTable("tests") // 1 使用表名

	// exist := db.Db.HasTable(&Person{})   // 2 使用模型

	fmt.Println(exist)

	// --自动迁移----------------------------------------
	// 创建表格时推荐使用自动迁移，自动迁移会创建表，添加列/索引，但无法改变现有列的类型，也无法删除列，目的是保护数据
	db.Db.AutoMigrate(&models.Test{}, &models.User{}) // 可以用逗号隔开，创建多个表格

	// ---增删改查-------------------------------------
	// ------增----------------
	xiaobai := models.User{Id: 1, Name: "xiaobai", Age: 18}
	xiaolan := models.User{Id: 2, Name: "lan", Age: 15}
	xiaohei := models.User{Id: 3, Name: "hei", Age: 11}
	db.Db.Create(&xiaobai)
	db.Db.Create(&xiaolan)
	db.Db.Create(&xiaohei)

	// ------查----------------
	var xb models.User
	db.Db.First(&xb, 1) // 查询第一条符合条件的数据，第一个参数接收查询到的数据，第二个参数就是条件，默认为模型的第一个字段
	// 相当于Sql语句 select * from users order by id limit 1
	fmt.Println(xb)
	db.Db.First(&xb, "name=?", "xiaobai") // 查询指定的字段
	fmt.Println(xb)

	// 查询所有数据
	db.Db.Find(&xb) // select * from users;
	fmt.Println(xb)

	// -----改-------需要先查到数据，然后才能修改
	var user models.User
	db.Db.First(&user, 1)                 // 默认使用id字段
	db.Db.Model(&user).Update("age", 111) // 这里的Model就是查询出来的结构体对象user，不是模型

	// 使用`map`可以同时修改多个属性
	var user2 models.User
	db.Db.First(&user2, 2)
	db.Db.Model(&user2).Updates(map[string]interface{}{"name": "小白", "age": 77}) // 注意id字段时不能修改的
	// update users set name='hello', age=18, actived=false, updated_at='2013-11-17 21:34:10' where id=111;

	// // 使用`struct`也可以更新多个属性
	// db.Db.Model(&user2).Updates(User{Name: "修改", Age: 77})

	// ---删------也要先查找到数据
	var user3 models.User
	db.Db.First(&user3, 3) // 默认使用id字段
	db.Db.Delete(&user3)

	// ---自定义表名，可以给表名加前缀---------------------------------------
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "sys_" + defaultTableName
	}

	db.Db.CreateTable(&models.Test{}) // 将Text作为model模型，表名会自动加上s

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

	db.Db.Debug().CreateTable(&Boy{}) // 在gorm语句中调用Debug()可以打印出对应的Sql语句
}
