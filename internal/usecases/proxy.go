package usecases

import "github.com/aerosystems/adapter-service/internal/models"

type ProxyUsecase struct {
	verifireRepo VerifireRepository
}

func NewProxyUsecase(verifireRepo VerifireRepository) *ProxyUsecase {
	return &ProxyUsecase{
		verifireRepo: verifireRepo,
	}
}

func (uc *ProxyUsecase) InspectData(data string, token string, clientIp string, serverIp string) (*models.InspectResult, error) {
	result, err := uc.verifireRepo.InspectData(data, token, clientIp, serverIp)
	if err != nil {
		return nil, err
	}
	return transformResponse(result), nil
}

func transformResponse(responsePayload *models.InspectDTO) *models.InspectResult {
	var payload models.InspectResult

	switch responsePayload.Code {
	case 400001:
		payload.Error = 31
		payload.Message = "Invalid email address"
	case 400002:
		payload.Error = 33
		payload.Message = "Invalid domain"
	case 404001:
		payload.Error = 34
		payload.Message = "No mail server is attached to this domain"
	default:
		switch responsePayload.Data {
		case "whitelist":
			payload.Result = true
			payload.Message = "This domain is in Whitelist"
		case "blacklist":
			payload.Result = false
			payload.Message = "This domain is in Blacklist"
		default:
			payload.Result = true
			payload.Unknown = new(bool)
			*payload.Unknown = true
			payload.Message = "Unknown domain. We will classify this domain shortly"
		}
	}
	return &payload
}
