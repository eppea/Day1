package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/users", GetUsers)
	e.Logger.Fatal(e.Start(":8080"))
}

func GetUsers(c echo.Context) error {
	return c.JSON(200, "Hello WORED")
}
