package handlers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
	"uri-shortener/internal/models/models_api"
	"uri-shortener/internal/pkg/errors/router_errors"
	"uri-shortener/internal/pkg/errors/usecase_errors"
	"uri-shortener/internal/pkg/router"
	"uri-shortener/internal/services/usecases"
)

type linkHandler struct {
	linkUseCase usecases.LinkUseCase
}

func NewLinkHandler(router fiber.Router, linkUseCase usecases.LinkUseCase) {
	handler := &linkHandler{
		linkUseCase: linkUseCase,
	}

	router.Route("", func(router fiber.Router) {
		router.Post("", handler.CreateLink)
		router.Get("/:tail", handler.FindLink)
	})
}

func (h *linkHandler) CreateLink(c *fiber.Ctx) error {
	linkRequest := new(models_api.NewLinkDto)
	if err := c.BodyParser(linkRequest); err != nil {
		return router.SendError(c, router_errors.BadRequestNoLink)
	}

	ttlMinutesString := c.Query("ttlMinutes")
	var ttlMinutes = 0
	if ttlMinutesString != "" {
		var err error
		ttlMinutes, err = strconv.Atoi(ttlMinutesString)
		if err != nil {
			return router.SendError(c, router_errors.BadRequest)
		}
	}

	shortLink, err := h.linkUseCase.Create(c.UserContext(), linkRequest.FullLink, ttlMinutes)
	if err != nil {
		return router.SendError(c, err)
	}
	linkDto := &models_api.LinkDto{
		ShortLink: shortLink,
	}

	return c.JSON(linkDto)
}

func (h *linkHandler) FindLink(c *fiber.Ctx) error {
	tail := c.Params("tail")
	if tail == "" {
		return router.SendError(c, router_errors.BadRequestNoLink)
	}

	fullLink, err := h.linkUseCase.GetFullLink(c.UserContext(), tail)
	if err != nil {
		if errors.Is(err, usecase_errors.LinkNotFoundError) {
			return router.SendError(c, router_errors.NotFoundNoLink)
		}
		return router.SendError(c, err)
	}

	return c.Redirect(fullLink, http.StatusFound)
}
