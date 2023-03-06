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

// e.g. http://localhost:8080/v1/user/1
func GetUser(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	userPo, _ := repo.GDB.GetUser(context.Background(), id)

	c.JSON(http.StatusOK, gin.H{
		"msg": user.ConvertUser(userPo),
	})
}

// e.g. http://localhost:8080/v1/user?Account=Ryu
func QueryUser(c *gin.Context) {
	var queryUser user.User
	c.ShouldBindQuery(&queryUser)
	log.Printf("%s queryUser=%v\n", utils.Tag("queryUser"), queryUser)

	userPo, _ := repo.GDB.GetUserByAccount(context.Background(), queryUser.Account)
	c.JSON(http.StatusOK, gin.H{
		"msg": user.ConvertUser(userPo),
	})
}

// e.g. curl -X POST http://localhost:8080/v1/user -H 'Content-Type: application/json' -d '{"acc":"Albert", "pwd":"abc123"}'
func CreatUser(c *gin.Context) {
	var userPost user.User
	if c.ShouldBind(&userPost) == nil {
		result, _ := repo.GDB.CreateUser(context.Background(),
			repo.CreateUserParams{userPost.Account, userPost.Password})

		id, _ := result.LastInsertId()
		userPo, _ := repo.GDB.GetUser(context.Background(), id)
		c.JSON(http.StatusOK, gin.H{
			"msg": user.ConvertUser(userPo),
		})
	}
}

// e.g. curl -X PUT http://localhost:8080/v1/user/2 -H 'Content-Type: application/json' -d '{"acc":"Ryu", "pwd":"ddd333"}'
func UpdateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	var userPut user.User
	if c.ShouldBind(&userPut) == nil {

		repo.GDB.UpdateUser(context.Background(),
			repo.UpdateUserParams{userPut.Account, userPut.Password, id})
		userPo, _ := repo.GDB.GetUser(context.Background(), id)
		c.JSON(http.StatusOK, gin.H{
			"msg": user.ConvertUser(userPo),
		})
	}

}
