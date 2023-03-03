package rest

import (
	"github.com/gin-gonic/gin"
	"m/interface/rest/v1"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	userRouter := router.Group("/v1/user")
	{
		userRouter.GET("/:name/:param", rest.GetUser)
		userRouter.GET("", rest.QueryUser)
		userRouter.POST("", rest.CreatUser)
		userRouter.PUT("/:id", rest.UpdateUser)
	}

	return router
}
