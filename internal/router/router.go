package router

import (
	"just-shopping-around-server/internal/handler"

	"github.com/labstack/echo/v4"
)

 

func CreateRouter() *echo.Echo{

	e := echo.New()
	e.GET("/news",handler.GetNews)
	e.POST("/auth", handler.Auth )
	return e
}