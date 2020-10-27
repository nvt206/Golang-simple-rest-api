package dto

import (
	"demo/models/dto"
	"demo/services"
	"demo/validations"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
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

	//update
	res,err := u.service.Update(ctx)
	if err!=nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK,gin.H{
		"code":http.StatusOK,
		"data":res,
	})

}

func (u userController) GetById(ctx *gin.Context) {
	user,err := u.service.GetById(ctx)
	if err !=nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK,gin.H{
		"user":user,
	})

}

func (u userController) Register(ctx *gin.Context) {

	var user dto.User
	if err := ctx.ShouldBindJSON(&user);err!=nil{
		ctx.AbortWithStatusJSON(http.StatusBadRequest,gin.H{
			"error":validations.GetErrors(err.(validator.ValidationErrors)),
		})
		return
	}

	res,err := u.service.Register(&user)

	if err!=nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated,gin.H{
		"user":res,
	})
}

func (u userController) GetAll(ctx *gin.Context) {
	users,err := u.service.GetAll(ctx)

	if err !=nil {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK,gin.H{
		"users":users,
	})

}
func NewUserController() UserController{
	return &userController{service: services.NewUserService()}
}
