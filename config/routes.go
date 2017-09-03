package config

import (
	"github.com/labstack/echo"
	"github.com/tsub/echo-sandbox/app/controller"
)

// Routes ...
func Routes(e *echo.Echo) {
	e.GET("/", controller.PostsIndex)
	e.GET("/posts", controller.PostsIndex)
}
