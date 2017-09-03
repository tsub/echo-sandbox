package config

import (
	"github.com/labstack/echo"
)

// StartApp ...
func StartApp() {
	e := echo.New()
	Routes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
