package main

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/users/:id", func(c echo.Context) error {
		// User ID from path `users/:id`
		id := c.Param("id")
		return c.String(http.StatusOK, id)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
