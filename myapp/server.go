package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"os"
	"time"

	echo "github.com/labstack/echo/v4"
)

type User struct {
	Name  string `json:"name" param:"name" query:"name"`
	Email string `json:"email" param:"email" query:"email"`
}

type CustomContext struct {
	echo.Context
}

func (c *CustomContext) Foo() {
	println("foo")
}

func (c *CustomContext) Bar() {
	println("bar")
}

func main() {
	e := echo.New()

	// custom context
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &CustomContext{c}
			return next(cc)
		}
	})

	/*
		e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				// Extract the credentials from HTTP request header and perform a security
				// check

				// For invalid credentials
				return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")

				// For valid credentials call next
				// return next(c)
			}
		})
	*/

	e.Static("/", "assets")

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "<strong>Hello, World!</strong>")
	})
	e.GET("/json", func(c echo.Context) error {
		u := &User{
			Name:  "Jon",
			Email: "jon@labstack.com",
		}
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().WriteHeader(http.StatusOK)
		return json.NewEncoder(c.Response()).Encode(u)
	})

	e.GET("/xml", func(c echo.Context) error {
		u := &User{
			Name:  "Jon",
			Email: "jon@labstack.com",
		}
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationXMLCharsetUTF8)
		c.Response().WriteHeader(http.StatusOK)
		return xml.NewEncoder(c.Response()).Encode(u)
	})

	e.GET("/file", func(c echo.Context) error {
		return c.File("img/toka1.jpg")
	})

	e.GET("/sendStreamFile", func(c echo.Context) error {
		f, err := os.Open("img/toka1.jpg")
		if err != nil {
			return err
		}
		return c.Stream(http.StatusOK, "image/jpg", f)
	})

	e.GET("/redirect", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "https://www.inter.it/jp")
	})

	e.Any("/users/*", func(c echo.Context) error {
		return c.String(http.StatusOK, "any!!")
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

	e.GET("/customContext", func(c echo.Context) error {
		cc := c.(*CustomContext)
		cc.Foo()
		cc.Bar()
		return cc.String(200, "OK")
	})

	e.GET("/writeCookie", func(c echo.Context) error {
		return writeCookie(c)
	})

	e.GET("/readCookie", func(c echo.Context) error {
		return readCookie(c)
	})

	e.Logger.Fatal(e.Start(":1323"))
}

func writeCookie(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "username"
	cookie.Value = "jon"
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)
	return c.String(http.StatusOK, "write a cookie")
}

func readCookie(c echo.Context) error {
	cookie, err := c.Cookie("username")
	if err != nil {
		return err
	}
	fmt.Println(cookie.Name)
	fmt.Println(cookie.Value)
	return c.String(http.StatusOK, "read a cookie")
}
