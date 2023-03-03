package rest

import (
	"github.com/gin-gonic/gin"
	"m/interface/rest/handler/v1"
	"m/interface/rest/middleware"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.Logger(), middleware.Auth(), gin.Recovery())

	userRouter := router.Group("/v1/user")
	{
		userRouter.GET("/:id", handler.GetUser)
		userRouter.GET("", handler.QueryUser)
		userRouter.POST("", handler.CreatUser)
		userRouter.PUT("/:id", handler.UpdateUser)
	}
	return router
}
