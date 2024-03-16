package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type InspectHandler struct {
	ProxyUsecase ProxyUsecase
}

func NewInspectHandler(proxyUsecase ProxyUsecase) *InspectHandler {
	return &InspectHandler{
		ProxyUsecase: proxyUsecase,
	}

}

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
// @Success 200 {object} models.InspectResult
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /domain/check [get]
func (ih InspectHandler) CheckData(c echo.Context) error {
	res, err := ih.ProxyUsecase.InspectData(
		c.QueryParam("data"),
		c.Get("token").(string),
		c.QueryParam("ip"),
		c.RealIP(),
	)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error inspecting domain")
	}
	return c.JSON(http.StatusOK, res)
}
