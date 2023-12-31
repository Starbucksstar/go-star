package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	. "star/src/entity"
	. "star/src/service"
	"star/src/util"
	"strconv"
)

type UserController struct {
	userService UserService
}

func NewUserController(us UserService) UserController {
	return UserController{
		userService: us,
	}
}

// Login @BasePath /users
// Login godoc
// @Summary User login
// @Schemes
// @Description login
// @Tags login
// @Accept json
// @Produce json
// @Success 200 {string} User
// @Router /users/login [get]
func (uc *UserController) Login(context *gin.Context) {
	var user User
	if err := context.ShouldBindQuery(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	uc.userService.FindUser(&user)
	token, _ := util.GenToken(user.Name, user.Password, strconv.Itoa(int(user.ID)))
	context.JSON(http.StatusOK, gin.H{"user": user, "token": token})
}

func (uc *UserController) SignUp(context *gin.Context) {
	var user User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	saveUser, _ := uc.userService.SaveUser(user)
	context.JSON(http.StatusCreated, saveUser)
}

func (uc *UserController) SignOut(context *gin.Context) {
	var user User
	id := context.Params.ByName("id")
	parseUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		panic(err)
	}
	user.ID = uint(parseUint)
	_ = uc.userService.DeleteUser(user)
	context.JSON(http.StatusOK, "delete success")
}
