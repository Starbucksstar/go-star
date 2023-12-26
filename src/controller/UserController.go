package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	. "star/src/entity"
	. "star/src/repository"
	"star/src/util"
	"strconv"
)

func Login(context *gin.Context) {
	var user User
	if err := context.ShouldBindQuery(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_ = QueryUserByNameAndPassword(&user)
	token, _ := util.GenToken(user.Name, user.City, strconv.Itoa(int(user.ID)))
	context.JSON(http.StatusOK, gin.H{"user": user, "token": token})
}

func SignUp(context *gin.Context) {
	var user User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	saveUser, _ := SaveUser(user)
	context.JSON(http.StatusCreated, saveUser)
}

func SignOut(context *gin.Context) {
	var user User
	id := context.Params.ByName("id")
	parseUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		panic(err)
	}
	user.ID = uint(parseUint)
	_ = DeleteUser(user)
	context.JSON(http.StatusOK, "delete success")
}
