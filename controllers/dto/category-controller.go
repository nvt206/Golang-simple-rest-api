package dto

import (
	"demo/models"
	"demo/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CategoryController interface {

	GetAll(ctx *gin.Context)
	Post(ctx *gin.Context)

}

type categoryController struct {
	repo services.CategoryService
}

func (c categoryController) Post(ctx *gin.Context) {

	var category models.Category

	if err:= ctx.ShouldBindJSON(&category);err!=nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
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

func NewCategoryController() CategoryController{
	return &categoryController{repo: services.NewCategoryService()}
}
