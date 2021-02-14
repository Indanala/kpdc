package routes

import (
	"template-api-gecho/controller"
	//"template-api-gecho/utils"

	"github.com/labstack/echo"
)

type Routing struct {
	example controller.Controller
}

func (Routing Routing) GetRoutes() *echo.Echo {
	e := echo.New()

	//e.GET("/posts/", Routing.example.PostsRegis)
	//e.POST("/posts/", Routing.example.PostsRegis, utils.Middleware)
	e.POST("/register/", Routing.example.PostsRegis)
	e.POST("/login/", Routing.example.GetLogin)
	e.GET("/login/", Routing.example.GetLogin)
	e.POST("/create/project", Routing.example.PostsCreate)
	e.POST("/create/usecase", Routing.example.PostsCreateForm)
	e.POST("/create/dfs", Routing.example.PostsCreatedfs)
	e.POST("/create/RP", Routing.example.PostsCreateRequestParam)
	e.POST("/create/Header", Routing.example.PostsCreateHeader)
	e.POST("/create/Res", Routing.example.PostsCreateRes)
	e.POST("/create/DO", Routing.example.PostsCreateDO)

	return e
}
