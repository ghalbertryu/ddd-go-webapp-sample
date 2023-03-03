package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"m/infrastructure/utils"
)

func Logger() gin.HandlerFunc {
	return func(context *gin.Context) {
		host := context.Request.Host
		url := context.Request.URL
		log.Printf("%s %s %s\n", utils.Tag("HTTP Logger"), host, url)
		context.Next()
	}
}

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		log.Println("授權驗證...")
		context.Next()
	}
}
