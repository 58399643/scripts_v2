// @File:  base.go
// @Time:  2022/3/1 5:04 PM
// @Author: ClassmateLin
// @Email: classmatelin.site@gmail.com
// @Site: https://www.classmatelin.top
// @Description:
// @Cron: * */1 * * *

package models

import "github.com/jinzhu/gorm"

type Base struct {
	gorm.Model
}
