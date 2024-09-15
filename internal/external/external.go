package external

import (
	"fmt"

	"github.com/oneweerachai/service1/internal/logger"

	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
)

type Client struct {
	httpClient *resty.Client
	baseURL    string
	apiKey     string
}

func NewClient(baseURL, apiKey string) *Client {
	client := resty.New()
	client.SetBaseURL(baseURL)
	client.SetHeader("Authorization", "Bearer "+apiKey)
	client.SetHeader("Content-Type", "application/json")

	return &Client{
		httpClient: client,
		baseURL:    baseURL,
		apiKey:     apiKey,
	}
}

func (c *Client) GetSomeData(endpoint string) (interface{}, error) {
	var result interface{}
	resp, err := c.httpClient.R().
		SetResult(&result).
		Get(endpoint)

	if err != nil {
		logger.NewLogger().Error("External API request failed", zap.Error(err))
		return nil, err
	}

	if resp.IsError() {
		logger.NewLogger().Error("External API returned error", zap.Int("status", resp.StatusCode()))
		return nil, fmt.Errorf("external API error: %s", resp.Status())
	}

	return result, nil
}
