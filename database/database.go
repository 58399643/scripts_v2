// @File:  database.go
// @Time:  2022/1/4 6:02 PM
// @Author: ClassmateLin
// @Email: classmatelin.site@gmail.com
// @Site: https://www.classmatelin.top
// @Description:

package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"scripts/global"
	"scripts/models"
)

var DB *gorm.DB

func init() {
	var err error

	DB, err = gorm.Open("sqlite3", "sqlite3.db")

	if err != nil {
		global.Log.Errorln("Can't not connect database.")
	}
	DB.AutoMigrate(&models.Code{})
}
