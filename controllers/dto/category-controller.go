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
	GetAll2(ctx *gin.Context)
	Post(ctx *gin.Context)
	Post2(ctx *gin.Context)
	Delete(ctx *gin.Context)

}

type categoryController struct {
	services.BaseService
	repo services.CategoryService
}

func (c categoryController) Post2(ctx *gin.Context) {

	var category dto.Category
	if err := ctx.ShouldBindJSON(&category);err!=nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"errors":validations.GetErrors(err.(validator.ValidationErrors)),
		})
		return
	}

	if err := c.BaseService.Create(&category);err!=nil{
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

	if err:= c.BaseService.GetOne(&category,"ID=?",uint(id));err!=nil{
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
	if err := c.repo.Delete(ctx,uint(id));err!=nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":err,
		})
		return
	}

	ctx.JSON(http.StatusOK,gin.H{
		"message":"complete delete",
	})

}

func (c categoryController) Post(ctx *gin.Context) {

	var category dto.Category

	if err:= ctx.ShouldBindJSON(&category);err!=nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":validations.GetErrors(err.(validator.ValidationErrors)),
		})
		return
	}
	res := c.repo.Post(&category)
	ctx.JSON(http.StatusOK,gin.H{
		"category":res,
	})


}

func (c categoryController) GetAll(ctx *gin.Context) {
	categories := c.repo.GetAll()
	ctx.JSON(http.StatusOK,gin.H{
		"categories":categories,
	})
}
func (c categoryController) GetAll2(ctx *gin.Context) {

	var categories []dto.Category

	err := c.BaseService.GetList(&categories,"true")
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
	return &categoryController{repo: services.NewCategoryService(),BaseService:services.NewBaseService()}


}
