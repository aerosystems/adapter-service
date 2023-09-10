package proxy

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

// CheckData godoc
// @Summary Get Data about domain/email address
// @Description Get Data about domain/email address
// @Tags data
// @Accept  json
// @Produce application/json
// @Security BearerAuth
// @Success 200 {object} ResponseCheckData
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /domain/check [get]
func (s *Service) CheckData(c echo.Context) error {
	token := c.Get("token").(string)
	data := c.QueryParam("data")
	clientIp := c.QueryParam("ip")
	serverIp := c.RealIP()

	res, err := s.InspectData(data, token, clientIp, serverIp)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error inspecting domain")
	}

	response, err := s.TransformCheckmailToProxy(*res)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error adapting response")
	}

	return c.JSON(http.StatusOK, response)
}
