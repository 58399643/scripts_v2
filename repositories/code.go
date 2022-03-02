// @File:  code.go
// @Time:  2022/3/1 5:04 PM
// @Author: ClassmateLin
// @Email: classmatelin.site@gmail.com
// @Site: https://www.classmatelin.top
// @Description:
// @Cron: * */1 * * *

package repositories

import (
	"fmt"
	"scripts/database"
	"scripts/models"
)

// CodeRepository
// @Description:
type CodeRepository struct {
	Base
	Model models.Code
}

// Create
// @Description: 创建或更新助力码
// @receiver c
// @param data
func (c CodeRepository) Create(data map[string]interface{}) models.Code {
	account, _ := data["account"]
	key, _ := data["key"]
	var code models.Code

	database.DB.Model(&c.Model).Where("key = ? and account = ?", key, account).First(&code)

	if code.ID == 0 {
		database.DB.Create(&models.Code{
			Account: data["account"].(string),
			Key:     data["key"].(string),
			Val:     data["val"].(string),
			Sort:    data["sort"].(int),
			Type:    uint(data["type"].(int)),
		}).Assign(&code)
	} else {
		database.DB.Model(&c.Model).Where("key = ? and account = ?", key, account).Select("Val", "Sort").Updates(data).Assign(&code)
	}

	return code
}

// GetCodeList
// @Description: 获取助力码列表
// @receiver c
// @param key 助力码健名
func (c CodeRepository) GetCodeList(key string) {
	codeList := map[string]interface{}{}
	database.DB.Model(&c.Model).Where("key = ?", key).Order("sort asc").Take(&codeList)
	fmt.Println(codeList)

}
