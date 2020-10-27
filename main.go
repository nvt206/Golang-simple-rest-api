package main

import (
	"demo/common"
	"demo/routes"
)

func main() {


	common.ConnectData()
	routes.ConfigRoute()
}
