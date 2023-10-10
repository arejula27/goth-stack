package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func PublicHandler(c echo.Context) error {
	// User ID from path `users/:id`

	return c.HTML(http.StatusOK, "<div><a href=http://localhost:3000/admin>Login</a>")
}

func PrivateHandler(c echo.Context) error {
	// User ID from path `users/:id`
	return c.String(http.StatusOK, "Hello, Admin!")
}
