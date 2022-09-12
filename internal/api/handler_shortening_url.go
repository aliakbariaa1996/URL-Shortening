package api

import (
	"context"
	"github.com/aliakbariaa1996/URL-Shortening/internal/services/shorteningURL"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserId  string `json:"user_id" binding:"required"`
}
type UrlGetRequest struct {
	ShortURL string `json:"short_url"`
}

// makeCreateShorteningHandler godoc
// @Summary makeCreateShorteningHandler
// @Description Get the Shortening URL
// @Tags shortening
// @Accept json
// @Produce json
// @Param long_url formData string true "Long Url"
// @Param user_id formData string true "User ID"
// @Router /short/create [post]

func (h *Handler) makeCreateShorteningHandler(shService shorteningURL.UseService) func(_ echo.Context) error {
	return func(c echo.Context) error {
		var err error
		var ctx context.Context
		var creationRequest UrlCreationRequest
		if err := c.Bind(&creationRequest); err != nil {
			return c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{"error": err.Error()})
		}

		res, err := shService.GenerateShortLink(creationRequest.LongUrl, creationRequest.UserId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		}

		res, err = shService.SaveUrlMapping(ctx, res, creationRequest.LongUrl, creationRequest.UserId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{"result": res})
	}
}

// makeGetShorteningHandler godoc
// @Summary makeGetShorteningHandler
// @Description Get the Shortening URL
// @Tags shortening
// @Accept json
// @Produce json
// @Router /short/get [get]
func (h *Handler) makeGetShorteningHandler(shService shorteningURL.UseService) func(_ echo.Context) error {
	return func(c echo.Context) error {
		var err error
		var ctx context.Context
		var getRequest UrlGetRequest
		shortUrl := c.Param(getRequest.ShortURL)

		res, err := shService.RetrieveInitialUrl(ctx, shortUrl)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{"result": res})
	}
}
