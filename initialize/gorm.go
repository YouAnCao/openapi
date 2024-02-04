package initialize

import (
	"fmt"
	"openapi/global"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Mysql init mysql connect
func Mysql() {
	m := global.Config.Mysql
	fmt.Printf("mysql config: username: %s, password:%s, url:%s", m.Username, m.Password, m.Url)
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/dbname?charset=utf8mb4&parseTime=True")
	gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
