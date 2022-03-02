// @File:  test.go
// @Time:  2022/3/2 4:05 PM
// @Author: ClassmateLin
// @Email: classmatelin.site@gmail.com
// @Site: https://www.classmatelin.top
// @Description:
// @Cron: * */1 * * *
package main

import "scripts/repositories"

func main() {
	repositories.CodeRepository{}.GetCodeList("jd_farm")
}
