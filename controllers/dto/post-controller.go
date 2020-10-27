package dto

import (
	"demo/models/dto"
	"demo/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type PostController interface {
	GetPostByUserAndCategory(ctx *gin.Context)
	Post(ctx *gin.Context)
}
type postController struct {
	service services.PostService
}

func (p postController) Post(ctx *gin.Context) {
	var post dto.Post
	if err := ctx.ShouldBindJSON(&post);err!=nil {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}
	res,err := p.service.Post(ctx,&post)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
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

	posts,err := p.service.FindListPost(&dto.Post{UserId: uint(user_id),CategoryId: uint(categoryid)})

	if err !=nil {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}


	ctx.JSON(http.StatusOK,gin.H{
		"posts":posts,
	})

}

func NewPostController() PostController{
	return &postController{service: services.NewPostService()}
}

