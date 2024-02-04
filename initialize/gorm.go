package initialize

import (
	"fmt"
	"openapi/global"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// Mysql init mysql connect
func Mysql() {
	m := global.Config.Mysql
	fmt.Printf("mysql config: username: %s, password:%s, url:%s\n", m.Username, m.Password, m.Url)
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/mall?charset=utf8mb4&parseTime=True", m.Username, m.Password, m.Url)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		fmt.Printf("mysql error:%s", err)
		return
	}
	sqlDb, err := db.DB()
	if err != nil {
		fmt.Printf("mysql error:%s", err)
		return
	}
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(time.Hour)

	var timestamp string
	db.Raw("select DATE_FORMAT(now(),'%Y-%m-%d %H:%i:%s')").Scan(&timestamp)
	fmt.Printf("mysql connect :%s", timestamp)

	global.Db = db
}
