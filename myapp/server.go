package main

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type User struct {
	Name  string `json:"name" param:"name" query:"name"`
	Email string `json:"email" param:"email" query:"email"`
}

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
	e.GET("/show", func(c echo.Context) error {
		// Get team and member from the query string
		team := c.QueryParam("team")
		member := c.QueryParam("member")
		return c.String(http.StatusOK, "team:"+team+", member:"+member)
	})

	e.GET("/bind/users/:name/:email", func(c echo.Context) error {
		u := new(User)
		if err := c.Bind(u); err != nil {
			// error handling
		}
		return c.JSON(http.StatusOK, u)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
