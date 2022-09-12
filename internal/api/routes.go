package api

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
)

func (s *Server) initRoutes() {
	s.GET("/swagger/*", echoSwagger.WrapHandler)
	s.GET("/health_check", healthCheck)

	apiV1 := s.Group("/api/v1")
	{
		apiAuth := apiV1.Group("/short")
		{
			apiAuth.POST("/create", s.handler.makeCreateShorteningHandler(s.ss.shorteningURL))
			apiAuth.GET("/get/:short_url", s.handler.makeGetShorteningHandler(s.ss.shorteningURL))
		}
	}
}

// healthCheck godoc
// @Summary healthCheck
// @Description health check server auth
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /health_check [get]
func healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}
