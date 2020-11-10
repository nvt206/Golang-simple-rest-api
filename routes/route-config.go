package routes

import (
	"demo/controllers"
	"demo/controllers/dto"
	"demo/middlewares"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)
var (
	UserController = dto.NewUserController()
	CategoryController = dto.NewCategoryController()
	PostController = dto.NewPostController()
)
func ConfigRoute()  {
	r := gin.Default()
	r.POST("/login",controllers.Login)
	r.POST("/register",UserController.Register)
	api := r.Group("/api",middlewares.AuthorizeJWT())
	{
		users := api.Group("/users",middlewares.AuthorizeJWT())
		{
			users.GET("/",UserController.GetAll)
			users.GET("/info",UserController.GetById)
			users.PUT("/update",UserController.Update)

		}
		categories := api.Group("/categories")
		{
			categories.GET("/",CategoryController.GetAll)
			categories.GET("/query",CategoryController.GetById)
			categories.POST("/create",CategoryController.Post)
			categories.DELETE("/delete/:id",CategoryController.Delete)
		}
		posts := api.Group("/posts")
		{
			posts.GET("/",PostController.GetPosts)
			posts.GET("/query",PostController.GetPostByUserAndCategory)// query?categoryId=1&userId=1
			posts.POST("/create",PostController.Post)
			posts.DELETE("/delete/:id",PostController.Delete)
		}
	}

	// should check only dev can see this
	//if(os.Getenv(constants.ENVIRONMENT) == constants.Development)
	r.GET("/swagger/*any",ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}
