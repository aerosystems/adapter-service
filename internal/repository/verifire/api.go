package verifire

type Api struct {
	BaseURL string
}

func NewApi(baseURL string) *Api {
	return &Api{
		BaseURL: baseURL,
	}
}
