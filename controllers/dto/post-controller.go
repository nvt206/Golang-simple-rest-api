package dto

import (
	"demo/models/dto"
	"demo/services"
	"demo/validations"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
)

type PostController interface {
	GetPostByUserAndCategory(ctx *gin.Context)
	Post(ctx *gin.Context)
	Delete(ctx *gin.Context)
	GetPosts(ctx *gin.Context)
}
type postController struct {
	services.BaseService
	service services.PostService
}

func (p postController) GetPosts(ctx *gin.Context) {

	userId := services.NewJWTService().GetClaims(ctx).ID

	var posts []dto.Post
	p.BaseService.GetList(&posts,"user_id=?",userId)
	for i,_ := range posts{
		p.service.MapUserAndCategory(&posts[i])
	}

	ctx.JSON(http.StatusBadRequest,gin.H{
		"posts":posts,
	})
}

func (p postController) Delete(ctx *gin.Context) {

	id,_ := strconv.Atoi(ctx.Param("id"))
	postDel := p.BaseService.GetOne(&dto.Post{},"ID=?",uint(id)).Value.(*dto.Post)

	if postDel == nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":"Invalid post id",
		})
		return
	}

	if !(p.BaseService.IsAdmin(ctx) || p.BaseService.IsRealUser(ctx,postDel.UserId)){
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":"Can not access",
		})
		return
	}

	err := p.BaseService.Delete(&dto.Post{},"ID=?",uint(id)).Error
	if err!=nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusBadRequest,gin.H{
		"message":"completed",
	})


}

func (p postController) Post(ctx *gin.Context) {
	var post dto.Post
	//check validate
	if err := ctx.ShouldBindJSON(&post);err!=nil {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"errors":validations.GetErrors(err.(validator.ValidationErrors)),
		})
		return
	}
	// check permit access
	//only real user can post a story
	if ! p.BaseService.IsRealUser(ctx,post.UserId){
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":"Can not access",
		})
		return
	}
	//just post a story when exist category

	if p.BaseService.GetOne(&dto.Category{},"ID=?",post.CategoryId).RowsAffected == 0{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":fmt.Sprintf("Not exists category with id %v",post.CategoryId),
		})
		return
	}
	//res,err := p.service.Post(ctx,&post)
	res := p.BaseService.Create(&post)
	if err := res.Error;err!=nil {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}

	//map user and category
	p.service.MapUserAndCategory(res.Value.(*dto.Post))

	ctx.JSON(http.StatusOK,gin.H{
		"post":res.Value,
	})


}

func (p postController) GetPostByUserAndCategory(ctx *gin.Context) {

	userId,_ := strconv.Atoi(ctx.Query("userId"))
	categoryId, _ := strconv.Atoi(ctx.Query("categoryId"))
	var posts []dto.Post
	if err := p.BaseService.GetList(&posts,"user_id=? and category_id=?",uint(userId),uint(categoryId)).Error;err!=nil {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}

	//map user and category by id
	for i,_ :=range posts {
		p.service.MapUserAndCategory(&posts[i])
	}


	ctx.JSON(http.StatusOK,gin.H{
		"posts":posts,
	})

}

func NewPostController() PostController{
	return &postController{
		service: services.NewPostService(),
		BaseService:services.NewBaseService(),
	}
}

