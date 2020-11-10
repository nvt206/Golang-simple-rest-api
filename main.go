package main

import (
	"demo/common"
	"demo/docs"
	"demo/routes"
)

func main() {

	docs.SwaggerInfo.Title = "Simple Restful API"
	docs.SwaggerInfo.Description = "This is a server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost" + ":" + "8080"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	common.ConnectData()
	routes.ConfigRoute()
}
