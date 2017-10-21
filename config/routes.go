package config

import (
	"github.com/labstack/echo"
	"github.com/tsub/echo-sandbox/app/controller"
)

// Routes ...
func Routes(e *echo.Echo) {
	e.GET("/api/v1/posts", controller.PostsIndex)
	e.GET("/api/v1/posts/:id", controller.PostShow)
	e.POST("/api/v1/posts", controller.PostCreate)
	e.PATCH("/api/v1/posts/:id", controller.PostUpdate)
	e.DELETE("/api/v1/posts/:id", controller.PostDelete)
}
