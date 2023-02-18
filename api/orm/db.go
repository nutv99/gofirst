package orm

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB
var err error

func InitDB() {

	// dsn := "root:12345678@tcp(127.0.0.1:3306)/it_asset_server?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := os.Getenv("GOLANGMYSQL_DNS")
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	Db.AutoMigrate(&User{})
}
