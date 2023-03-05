package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"m/domain/user"
	"m/infrastructure/repo"
	"m/infrastructure/utils"
	"net/http"
	"strconv"
)

// e.g. http://localhost:8080/v1/user/id
func GetUser(c *gin.Context) {
	id := c.Param("id")
	log.Println("getUser. id=", id)
	user, _ := repo.GDB.GetUser(context.Background(), 1)
	c.JSON(http.StatusOK, gin.H{
		"msg": user,
	})
}

// e.g. http://localhost:8080/v1/user?acc=Albert
func QueryUser(context *gin.Context) {
	var queryUser user.User
	context.ShouldBindQuery(&queryUser)

	log.Printf("%s queryUser=%v\n", utils.Tag("queryUser"), queryUser)
	context.JSON(http.StatusOK, gin.H{
		"msg": queryUser,
	})
}

// e.g. curl -X POST http://localhost:8080/v1/user -H 'Content-Type: application/json' -d '{"acc":"Albert", "pwd":"abc123"}'
func CreatUser(context *gin.Context) {
	var userPost user.User
	if context.ShouldBind(&userPost) == nil {
		log.Printf("%s userPost=%v\n", utils.Tag("creatUser"), userPost)

		context.JSON(http.StatusOK, gin.H{
			"msg": userPost,
		})
	}
}

// e.g. curl -X PUT http://localhost:8080/v1/user/22 -H 'Content-Type: application/json' -d '{"uid":321, "acc":"Ryu", "pwd":"ddd333"}'
func UpdateUser(context *gin.Context) {
	idStr := context.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err == nil {
		var userPut user.User
		if context.ShouldBind(&userPut) == nil {
			newUser := user.NewUser(userPut.Account)
			newUser.Id = id
			newUser.Password = userPut.Password
			context.JSON(http.StatusOK, gin.H{
				"msg": newUser,
			})
		}
	} else {
		log.Println(err)
	}
}
