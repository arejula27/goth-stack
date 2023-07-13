package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func NotFoundMiddleware(c *fiber.Ctx) error {
	return c.Status(http.StatusNotFound).Render("error/404", nil)
}

