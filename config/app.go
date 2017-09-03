package config

import (
	"os"

	"github.com/labstack/echo"
)

// StartApp ...
func StartApp() {
	e := echo.New()
	Routes(e)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
