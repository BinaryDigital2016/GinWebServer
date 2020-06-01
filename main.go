package main

import (
	"gin_webserver/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	// Engin
	router := gin.Default()
	//router := gin.New()

	router.LoadHTMLFiles("./html/home.html")

	router.GET("/hello", handlers.Hello)
	router.GET("/home", handlers.Home)
	// 指定地址和端口号
	router.Run(":9090")
}
