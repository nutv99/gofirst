package main

import (
	"fmt"
	"net/http"
         "database/sql"
         "github.com/go-sql-driver/mysql"
	."github.com/tbxark/g4vercel"
)

func main(w http.ResponseWriter, r *http.Request) {
	server := New()
	

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
				"idnew": connectDB,
				"conncetdb": connectDB,
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
