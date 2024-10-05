package router

import (
	"github.com/gofiber/fiber/v2"
	"uri-shortener/internal/pkg/errors"
)

func SendError(c *fiber.Ctx, err error) error {
	resultErr, code := errors.GetErrorMessageAndCode(err)
	_ = c.SendStatus(code)
	return c.JSON(resultErr)
}
