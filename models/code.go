// @File:  code.go
// @Time:  2022/3/1 4:50 PM
// @Author: ClassmateLin
// @Email: classmatelin.site@gmail.com
// @Site: https://www.classmatelin.top
// @Description:
// @Cron: * */1 * * *

package models

// Code
// @Description: 助力码表定义
type Code struct {
	Type    uint   `gorm:"column:type;type:tinyint;default 0;"json:"type"` // 助力码类型
	Account string `gorm:"column:account;type:varchar(100);not null;index;default:'';"json:"account"`
	Key     string `gorm:"column:key;type:varchar(30);not null;index;default:'';"json:"key"`
	Val     string `gorm:"column:val;type:varchar(255);not null;default:'';"json:"val"`      // 助力码健值
	Sort    int    `gorm:"column:sort;type:unsigned int(11);not null;default 0;"json:"sort"` // 助力码排序
	Base
}

// TableName
// @Description: 定义表名
// @receiver c
// @return string
func (c Code) TableName() string {
	return "code"
}
