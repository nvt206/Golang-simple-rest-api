package controllers

import (
	"demo/models"
	"demo/models/dto"
	"demo/services"
	"demo/validations"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/mitchellh/mapstructure"
	"net/http"
	"time"
)

var (
	UserService = services.NewUserService()
)
func Login(ctx *gin.Context) {

	var validation validations.LoginValidation
	if err:= ctx.ShouldBindJSON(&validation);err!=nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"code":http.StatusBadRequest,
			"errors":validations.GetErrors(err.(validator.ValidationErrors)),
		})
		return
	}
	email := validation.Email
	password := validation.Password
	user := UserService.FindOne(&dto.User{Email: email})
	if user == nil {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":"Incorrect username or password",
		})
		return
	}

	if user.Password != password {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":"Incorrect username or password",
		})
		return
	}
	claims := services.MyCustomClaims{
		user.ID,
		user.IsAdmin,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute*5).Unix(),
		},
	}

	var userInfo models.UserInfo
	mapstructure.Decode(user,&userInfo)

	jwtService := services.NewJWTService()
	token := jwtService.GenerateTokenWithClaims(claims)
	ctx.JSON(http.StatusOK,gin.H{
		"token":token,
		"info":userInfo,
	})
}