package dto

import (
	"demo/models/dto"
	"demo/services"
	"demo/validations"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/mitchellh/mapstructure"
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
	services.BaseService
	service services.UserService
}

func (u userController) Update(ctx *gin.Context) {

	//bind user from client request
	var user dto.User
	if err := ctx.ShouldBindJSON(&user);err!=nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"errors": validations.GetErrors(err.(validator.ValidationErrors)),
		})
		return
	}

	//Check role user update
	// just user or admin can update
	if !(u.BaseService.IsAdmin(ctx)||u.BaseService.IsRealUser(ctx,user.ID)){
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":"Can not access",
		})
		return
	}


	//find user by id
	if err := u.BaseService.GetOne(&dto.User{},"ID=?",user.ID).Error;err!=nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}

	//check email exist

	if u.BaseService.GetOne(&dto.User{},"email=?",user.Email).RowsAffected != 0{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":"Email is exists!",
		})
		return
	}


	//update
	res := u.BaseService.Update(&user)
	if err:= res.Error;err!=nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK,gin.H{
		"code":http.StatusOK,
		"data":res.Value,
	})

}

func (u userController) GetById(ctx *gin.Context) {

	//get id
	id,_ := strconv.Atoi(ctx.Query("id"))
	//user,err := u.service.GetById(ctx)
	res := u.BaseService.GetOne(&dto.User{},"ID=?",uint(id))
	if err := res.Error; err !=nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}

	var userinfo dto.UserInfo
	mapstructure.Decode(res.Value,userinfo)
	ctx.JSON(http.StatusOK,gin.H{
		"user":userinfo,
	})

}


// @Summary Register User
// @Tags Auth
// @Security ApiKeyAuth
// @Accept json
// @Produce  json
// @Param   body  body   validations.LoginValidation  true "body"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"message"}"
// @Failure 400 {string} json
// @Router /register [POST]
func (u userController) Register(ctx *gin.Context) {

	var user dto.User
	if err := ctx.ShouldBindJSON(&user);err!=nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"code":http.StatusBadRequest,
			"errors":validations.GetErrors(err.(validator.ValidationErrors)),
		})
		return
	}

	//res,err := u.service.Register(&user)
	res := u.BaseService.Create(&user)
	if err:=res.Error; err!=nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}
	userInfo := u.service.ConvertUserToUserInfo(res.Value.(*dto.User))
	ctx.JSON(http.StatusCreated,gin.H{
		"user":userInfo,
	})
}

func (u userController) GetAll(ctx *gin.Context) {

	if ! u.BaseService.IsAdmin(ctx){
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":"Can not access",
		})
		return
	}
	var users []dto.User

	res := u.BaseService.GetList(&users,"true")

	if err := res.Error; err !=nil {
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
	return &userController{
		service: services.NewUserService(),
		BaseService: services.NewBaseService(),
	}
}
