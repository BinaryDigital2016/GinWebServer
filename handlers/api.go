package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Hello(ctx *gin.Context) {
	log.Println(">>>> hello gin start <<<<")
	ctx.JSON(200, gin.H{
		"code":    200,
		"success": true,
	})
}

func Home(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "home.html", nil)
}
