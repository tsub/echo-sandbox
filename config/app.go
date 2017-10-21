package config

import (
	"os"

	"github.com/labstack/echo"
	"github.com/tsub/echo-sandbox/app/model"
)

// StartApp ...
func StartApp() {
	e := echo.New()
	Routes(e)

	db, err := model.NewDatabase();
	if err != nil {
		e.Logger.Fatal(err)
	}
	db.SetupDB()

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
