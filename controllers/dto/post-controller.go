package dto

import (
	"demo/models"
	"demo/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var (
	userRepo = services.NewUserService()
	categoryRepo = services.NewCategoryService()
)

type PostController interface {
	GetPostByUserAndCategory(ctx *gin.Context)
	Post(ctx *gin.Context)
}
type postController struct {
	service services.PostService
}

func (p postController) Post(ctx *gin.Context) {
	var post models.Post
	if err := ctx.ShouldBindJSON(&post);err!=nil {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}

	jwtService := services.NewJWTService()
	if ! jwtService.IsPermit(uint(post.UserId),ctx){
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error": "Can not access",
		})
		return
	}
	post.User = userRepo.FindById(post.UserId)
	post.Category = categoryRepo.GetById(post.CategoryId)

	if post.Category == nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error": fmt.Sprintf("Not exists category with id %v",post.CategoryId),
		})
		return
	}

	res := p.service.Post(&post)
	if res == nil {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":"can insert",
		})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{
		"post":res,
	})


}

func (p postController) GetPostByUserAndCategory(ctx *gin.Context) {

	user_id, _ := strconv.Atoi(ctx.Query("userid"))
	categoryid, _ := strconv.Atoi(ctx.Query("categoryid"))

	//posts := p.repo.GetByUserAndCategory(uint(user_id),uint(categoryid))
	posts := p.service.FindListPost(&models.Post{UserId: uint(user_id),CategoryId: uint(categoryid)})
	for i,_ :=range posts {
		posts[i].User = userRepo.FindById(posts[i].UserId)
		posts[i].Category = categoryRepo.GetById(posts[i].CategoryId)
	}


	ctx.JSON(http.StatusOK,gin.H{
		"posts":posts,
	})

}

func NewPostController() PostController{
	return &postController{service: services.NewPostService()}
}

