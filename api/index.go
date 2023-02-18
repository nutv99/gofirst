
package handler

import (
	"fmt"
	"net/http"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	. "github.com/tbxark/g4vercel"
)
var Db *gorm.DB
var err error

type User struct {
	gorm.Model
	UserName string
	Password string
	FullName string
	Avatar   string
}

func Handler(w http.ResponseWriter, r *http.Request) {
	server := New()
	
	GOLANGMYSQL_DNS :="lbg5pjees347lrun2wdl:pscale_pw_CgPWWfYLYTy3ziSn28WizLQ3fpt3dTTU24kgX82qxNA@tcp(ap-southeast.connect.psdb.cloud)/it_asset?tls=true"
	dsn := GOLANGMYSQL_DNS
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	Db.AutoMigrate(&User{})

	server.GET("/", func(context *Context) {
		context.JSON(200, H{
			"message": "hello go from vercel !!!!",
		})
	})
	server.GET("/hello888", func(context *Context) {
		name := context.Query("name")
		if name == "" {
			context.JSON(400, H{
				"message": "name not found",
			})
		} else {
			context.JSON(200, H{
				"data": fmt.Sprintf("Hello %s!", name),
			})
		}
	})
	server.GET("/user/:id", func(context *Context) {
		context.JSON(400, H{
			"data": H{
				"id": context.Param("id"),
			},
		})
	})
	server.GET("/long/long/long/path/*test", func(context *Context) {
		context.JSON(200, H{
			"data": H{
				"url": context.Path,
			},
		})
	})
	server.Handle(w, r)
}
