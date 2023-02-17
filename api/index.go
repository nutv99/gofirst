package handler

import (
	"fmt"
	
	
	"gorm.io/driver/mysql"	
       ."github.com/tbxark/g4vercel"
)

var Db *gorm.DB
var err error

func Handler(w http.ResponseWriter, r *http.Request) {
	server := New()
	
	//GOLANGMYSQL_DNS="lbg5pjees347lrun2wdl:pscale_pw_CgPWWfYLYTy3ziSn28WizLQ3fpt3dTTU24kgX82qxNA@tcp(ap-southeast.connect.psdb.cloud)/it_asset?tls=true"
	//Db, err = gorm.Open(mysql.Open(GOLANGMYSQL_DNS), &gorm.Config{})
	//if err != nil {
	//	panic("failed to connect database")
	//}
	
	
	 
	
	

	server.GET("/", func(context *Context) {
		context.JSON(200, H{
			"message": "hello go from vercel by nutv99 !!!!",
		})
	})
	server.GET("/hello", func(context *Context) {
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
