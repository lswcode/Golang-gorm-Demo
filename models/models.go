package models

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

// 模型名和表名的映射规则
// 1  第一个大写字母变为小写
// 2  遇到其他大写字母变为小写并且在前面加下划线
// 3  复数形式，自动加s

// User --> users
// UserInfo --> user_infos

//  gin-gorm常用的数据库API
//  First: 按主键升序查找符合条件的第一条数据
//  Last: 按主键升序查找最后一条数据
//  FirstOrCreate: 如果未查找到指定数据，则创建
//  Find: 查询符合条件的所有数据
//  Where: 配合sql语句进行查询
//  Select: 指定要检索的字段
//  Create: 插入单条数据
//  Save: 将会更新所有的字段，即使该字段没有修改
//  Update, Updates: 更新被修改过的字段
//  Delete: 删除指定数据
//  Not: 不等于，sql语句一般使用<>表示不等于
//  Or: 或者
//  Order: 指定查询数据的顺序
//  Limit: 指定查询数据的最大数量
//  Offset: 指定要跳过的数据数量
//  Scan: 将查询到的数据，保存到指定的结构体变量中
//  Count: 获取指定模型的数据条数
//  Group: Group将查询到的数据按规则分组
//  Having: 配合Group使用，过滤Group中的数据
//  Join: 用于多表联合查询
//  FirstOrInit: 获取第一条匹配的数据，如果不存在，则初始化一条新数据（仅适用于 struct，map条件）
//  Scopes: 将当前数据库连接传递到参数函数中，然后在函数中就可以接收到db，写具体的操作
