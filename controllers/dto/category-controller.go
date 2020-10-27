package dto

import (
	"demo/models/dto"
	"demo/services"
	"demo/validations"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
)

type CategoryController interface {

	GetAll(ctx *gin.Context)
	GetById(ctx *gin.Context)
	Post(ctx *gin.Context)
	Delete(ctx *gin.Context)

}

type categoryController struct {
	services.BaseService
	service services.CategoryService
}

func (c categoryController) Post(ctx *gin.Context) {

	var category dto.Category
	if err := ctx.ShouldBindJSON(&category);err!=nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"errors":validations.GetErrors(err.(validator.ValidationErrors)),
		})
		return
	}

	if err := c.BaseService.Create(&category).Error;err!=nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK,gin.H{
		"category":category,
	})
}

func (c categoryController) GetById(ctx *gin.Context) {
	// get id

	id,_ := strconv.Atoi(ctx.Query("id"))
	var category dto.Category

	if err:= c.BaseService.GetOne(&category,"ID=?",uint(id)).Error;err!=nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{
		"category": category,
	})

}

func (c categoryController) Delete(ctx *gin.Context) {

	//get id
	id,err := strconv.Atoi(ctx.Param("id"))
	if err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}

	if !c.BaseService.IsAdmin(ctx){
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":"Can not access",
		})
		return
	}

	if err := c.BaseService.Delete(&dto.Category{},"ID=?",uint(id)).Error;err!=nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK,gin.H{
		"message":"complete delete",
	})

}
func (c categoryController) GetAll(ctx *gin.Context) {

	var categories []dto.Category

	err := c.BaseService.GetList(&categories,"true").Error
	if err!=nil{
		ctx.JSON(http.StatusOK,gin.H{
			"error":err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{
		"categories":categories,
	})
}

func NewCategoryController() CategoryController{
	return &categoryController{
		service: services.NewCategoryService(),
		BaseService:services.NewBaseService(),
	}
}
