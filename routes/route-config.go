package routes

import (
	"demo/controllers"
	"demo/controllers/dto"
	"demo/middlewares"
	"github.com/gin-gonic/gin"
)

var (
	UserController = dto.NewUserController()
	CategoryController = dto.NewCategoryController()
	PostController = dto.NewPostController()
)

func ConfigRoute()  {

	r := gin.Default()

	api := r.Group("/api")
	{
		api.POST("/login",controllers.Login)
		api.POST("/register",UserController.Register)
		users := api.Group("/users",middlewares.AuthorizeJWT())
		{
			users.GET("/",UserController.GetAll)
			users.GET("/:id",UserController.GetById)
			users.PUT("/update/:id",UserController.Update)

		}
		categories := api.Group("/categories")
		{
			categories.GET("/",CategoryController.GetAll)
			categories.POST("/create",CategoryController.Post)
		}
		posts := api.Group("/posts")
		{

			posts.GET("/query",PostController.GetPostByUserAndCategory)// query?categoryid=1&userid=1
			posts.POST("/create",middlewares.AuthorizeJWT(),PostController.Post)
		}
	}
	r.Run()

}
