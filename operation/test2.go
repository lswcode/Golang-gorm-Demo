package test1

import (
	"gorm_project/db"
	"gorm_project/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func TestFun() {
	db.Db, _ = gorm.Open("mysql", "root:Lswmysql123.@/gorm_project?charset=utf8&parseTime=True&loc=Local")

	var user models.User

	db.Db.First(&user) // 按主键升序查找符合条件的第一条数据

	db.Db.Last(&user) // 按主键升序查找最后一条数据

	db.Db.Find(&user) // 获取所有数据

	db.Db.First(&user, 10) // 指定主键获取数据

	// ----------------------------------
	//  配合简单的SQL语句: Where 查询条件

	// 根据指定条件查询，获取所有符合条件的数据
	db.Db.Find(&user, "name = ?", "jinzhu")

	// 和上面的语句是等价的，获取所有符合条件的数据
	db.Db.Where("name = ?", "jinzhu").Find(&user)

	// 获取第一条符合条件的数据
	db.Db.Where("name = ?", "jinzhu").First(&user)

	// like模糊查询 %里面写内容%
	db.Db.Where("name LIKE ?", "%jin%").Find(&user)

	// -----------------------------------------
	// Not不等于
	db.Db.Not("name", "jinzhu").First(&user)
	// select * from users where name <> "jinzhu" limit 1

	// Or或者
	db.Db.Where("role = ?", "admin").Or("role = ?", "super_admin").Find(&user)
	// select * from users where role = 'admin' or role = 'super_admin'

	// Order指定查询数据的顺序
	db.Db.Order("age desc, name").Find(&user) // 注意: order要写在find前面
	// select * from users order by age desc, name;

	// ---------------------------------------------
	// Scopes：将当前数据库连接传递到参数函数中，然后在函数中就可以接收到db，写具体的操作
	db.Db.Scopes(AmountGreaterThan1000, PaidWithCreditCard).Find(&user)
	// 查找所有信用卡订单和金额大于1000

	// --错误处理-------------------------------
	// 如果执行Sql查询出现错误，会保存在*gorm.DB的Error字段
	result := db.Db.Where("name = ?", "jinzhu").First(&user)
	if result.Error != nil {
		println("111")
	}
	// 或者
	err := db.Db.Where("name = ?", "tizi365").First(&user).Error
	if err != nil {
		println("111")
	}

	// 处理多个错误，使用GetErrors()
	errors := db.Db.First(&user).Limit(10).Find(&user).GetErrors()
	println(len(errors)) // 打印错误数量

	// 遍历错误内容
	for _, err := range errors {
		println(err)
	}

}

func AmountGreaterThan1000(db *gorm.DB) *gorm.DB {
	return db.Where("amount > ?", 1000)
}

func PaidWithCreditCard(db *gorm.DB) *gorm.DB {
	return db.Where("pay_mode_sign = ?", "C")
}
