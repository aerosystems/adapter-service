package usecases

import "github.com/aerosystems/adapter-service/internal/models"

type VerifireRepository interface {
	InspectData(data string, token string, clientIp string, serverIp string) (*models.InspectDTO, error)
}
