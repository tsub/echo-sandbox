package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

// PostsIndex ...
func PostsIndex(c echo.Context) error {
	return c.String(http.StatusOK, "posts/index")
}
