package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type User struct {
	Nama string
	Age  int
}

type Response struct {
	Massage string
	Data    interface{}
}

var data []User = []User{
	{"alta", 10},
	{"angga", 11},
}

func main() {
	e := echo.New()
	e.GET("/users", GetUsers)
	e.GET("/users/:id", GetDetailUsers)
	e.POST("/login/", Login)
	e.Logger.Fatal(e.Start(":8080"))
}
func Login(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	return c.JSON(http.StatusOK, Response{
		"sukses", email + "" + password,
	})
}
func GetUsers(c echo.Context) error {
	page := c.QueryParam("page")
	limit := c.QueryParam("limit")
	// logic
	return c.JSON(http.StatusOK, Response{
		"sukses", page + "&" + limit,
	})
}
func GetDetailUsers(c echo.Context) error {
	id := c.Param("id")
	// logic
	return c.JSON(http.StatusOK, Response{
		"sukses", id,
	})
}
