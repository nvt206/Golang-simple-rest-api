package dto

import (
	"demo/models"
	"demo/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserController interface {
	GetAll(ctx *gin.Context)
	GetById(ctx *gin.Context)
	Register(ctx *gin.Context)
	Update(ctx *gin.Context)
}

type userController struct {
	 service services.UserService
}

func (u userController) Update(ctx *gin.Context) {

	id,_ := strconv.Atoi(ctx.Param("id"))

	jwtService := services.NewJWTService()

	if ! jwtService.IsPermit(uint(id),ctx){
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error": "Can not access",
		})
		return
	}

	findUser := u.service.FindById(uint(id))

	if findUser==nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error": "Not exists",
		})
		return
	}

	var user models.User
	if err := ctx.ShouldBindJSON(&user);err!=nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error": err.Error(),
		})
		return
	}
	user.ID = uint(id)
	res := u.service.Update(&user)
	ctx.JSON(http.StatusNoContent,gin.H{
		"user":res,
	})

}

func (u userController) GetById(ctx *gin.Context) {

	id,_ := strconv.Atoi(ctx.Param("id"))
	var user models.User
	user.ID = uint(id)

	resUser := u.service.FindOne(&user)
	if resUser==nil ||id==0{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":"Not exits",
		})
		return
	}

	ctx.JSON(http.StatusOK,gin.H{
		"user": resUser,
	})


}

func (u userController) Register(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user);err!=nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error": err.Error(),
		})
		return
	}


	res := u.service.Register(&user)
	if res == nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":"Invalid input",
		})
		return
	}

	ctx.JSON(http.StatusCreated,gin.H{
		"user":res,
	})
}

func (u userController) GetAll(ctx *gin.Context) {
	users := u.service.GetAll()

	ctx.JSON(http.StatusOK,gin.H{
		"users":users,
	})

}
func NewUserController() UserController{
	return &userController{service: services.NewUserService()}
}
