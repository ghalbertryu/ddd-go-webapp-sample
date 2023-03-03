package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"m/domain/user"
	"m/infrastructure/utils"
	"net/http"
	"strconv"
)

// e.g. http://localhost:8080/v1/user/Albert/xxx
func GetUser(context *gin.Context) {
	name := context.Param("name")
	param := context.Param("param")

	log.Println("getUser. param=", param)
	context.JSON(http.StatusOK, gin.H{
		"msg": user.User{1, name, 33},
	})
}

// e.g. http://localhost:8080/v1/user?name=Albert
func QueryUser(context *gin.Context) {
	var queryUser user.User
	context.ShouldBindQuery(&queryUser)

	log.Printf("%s queryUser=%v\n", utils.Tag("queryUser"), queryUser)
	context.JSON(http.StatusOK, gin.H{
		"msg": queryUser,
	})
}

// e.g. curl -X POST http://localhost:8080/v1/user -H 'Content-Type: application/json' -d '{"uid":123, "name":"Albert"}'
func CreatUser(context *gin.Context) {
	var userPost user.User
	if context.ShouldBind(&userPost) == nil {
		log.Printf("%s userPost=%v\n", utils.Tag("creatUser"), userPost)

		context.JSON(http.StatusOK, gin.H{
			"msg": userPost,
		})
	}
}

// e.g. curl -X PUT http://localhost:8080/v1/user/22 -H 'Content-Type: application/json' -d '{"uid":123, "name":"Ryu", "age":22}'
func UpdateUser(context *gin.Context) {
	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)
	if err == nil {
		var userPut user.User
		if context.ShouldBind(&userPut) == nil {
			context.JSON(http.StatusOK, gin.H{
				"msg": user.User{id, userPut.Name, userPut.Age},
			})
		}
	}
}
