package rest

import "github.com/aerosystems/adapter-service/internal/models"

type ProxyUsecase interface {
	InspectData(data string, token string, clientIp string, serverIp string) (*models.InspectResult, error)
}
