package models

type InspectResult struct {
	Error   int    `json:"error"`
	Result  bool   `json:"result"`
	Unknown *bool  `json:"unknown,omitempty"`
	Message string `json:"message,omitempty"`
}

type InspectDTO struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}
