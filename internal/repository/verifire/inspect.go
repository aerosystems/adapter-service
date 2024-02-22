package verifire

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/aerosystems/adapter-service/internal/models"
	"net/http"
)

type InspectResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type InspectRequestBody struct {
	Data     string `json:"data"`
	ClientIp string `json:"clientIp"`
}

func (a Api) InspectData(data, token, clientIp, serverIp string) (*models.InspectDTO, error) {
	var responsePayload InspectResponse

	requestPayload := InspectRequestBody{
		Data:     data,
		ClientIp: clientIp,
	}

	requestPayloadJSON, err := json.Marshal(requestPayload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/v1/data/inspect", a.BaseURL), bytes.NewBuffer(requestPayloadJSON))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Forwarded-For", serverIp)
	req.Header.Set("X-Api-Key", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&responsePayload); err != nil {
		return nil, err
	}
	return &models.InspectDTO{
		Code:    responsePayload.Code,
		Message: responsePayload.Message,
		Data:    requestPayload.Data,
	}, nil
}
