package proxy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Service struct {
	Endpoint string
}

func NewService(endpoint string) *Service {
	return &Service{
		Endpoint: endpoint,
	}
}

type RequestInspectCheckmail struct {
	Data     string `json:"data"`
	ClientIp string `json:"clientIp"`
}

func (s *Service) InspectData(data, token, clientIp, serverIp string) (*ResponseInspectCheckmail, error) {
	var responsePayload ResponseInspectCheckmail

	requestPayload := RequestInspectCheckmail{
		Data:     data,
		ClientIp: clientIp,
	}

	requestPayloadJSON, err := json.Marshal(requestPayload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/v1/data/inspect", s.Endpoint), bytes.NewBuffer(requestPayloadJSON))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Forwarded-For", serverIp)
	req.Header.Set("X-API-KEY", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&responsePayload); err != nil {
		return nil, err
	}
	return &responsePayload, nil
}
