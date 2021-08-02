package main

import (
	"fmt"
	"net/http"
	"os"

	"templ/handlers/user"

	conn "templ/repositories"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middleware ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Health Check ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	e.GET("/active", func(c echo.Context) error {
		return c.String(http.StatusOK, "")
	})

	// CRUD ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	e.POST("/users", user.SaveHandler)
	e.GET("/users/:id", user.GetHandler)
	e.PUT("/users/:id", user.UpdateHandler)
	e.DELETE("/users/:id", user.DeleteHandler)

	//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))

	defer conn.Db.Close()

}
