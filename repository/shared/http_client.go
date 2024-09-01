package shared

import "github.com/go-resty/resty/v2"

func CreateClient(hostURL string, apiKey string) *resty.Client {
	client := resty.New().
		SetBaseURL(hostURL)

	return client
}
