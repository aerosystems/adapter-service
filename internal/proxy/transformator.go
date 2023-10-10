package proxy

type ResponseInspectCheckmail struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type ResponseCheckData struct {
	Error   int    `json:"error"`
	Result  bool   `json:"result"`
	Unknown *bool  `json:"unknown,omitempty"`
	Message string `json:"message,omitempty"`
}

func (s *Service) TransformCheckmailToProxy(responsePayload ResponseInspectCheckmail) (ResponseCheckData, error) {
	var payload ResponseCheckData

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

	return payload, nil
}
