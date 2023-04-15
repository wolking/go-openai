package openai

import (
	"fmt"
	"net/http"
	"net/url"
)

const (
	openaiAPIURLv1                 = "https://api.openai.com/v1"
	defaultEmptyMessagesLimit uint = 300

	azureAPIPrefix         = "openai"
	azureDeploymentsPrefix = "deployments"
)

type APIType string

const (
	APITypeOpenAI  APIType = "OPEN_AI"
	APITypeAzure   APIType = "AZURE"
	APITypeAzureAD APIType = "AZURE_AD"
)

const AzureAPIKeyHeader = "api-key"

// ClientConfig is a configuration of a client.
type ClientConfig struct {
	authToken string

	BaseURL    string
	OrgID      string
	APIType    APIType
	APIVersion string // required when APIType is APITypeAzure or APITypeAzureAD
	Engine     string // required when APIType is APITypeAzure or APITypeAzureAD

	HTTPClient *http.Client

	EmptyMessagesLimit uint
}

func DefaultConfig(authToken string) ClientConfig {
	return ClientConfig{
		authToken: authToken,
		BaseURL:   openaiAPIURLv1,
		APIType:   APITypeOpenAI,
		OrgID:     "",

		HTTPClient: &http.Client{},

		EmptyMessagesLimit: defaultEmptyMessagesLimit,
	}
}

func DefaultAzureConfig(apiKey, baseURL, engine string) ClientConfig {
	return ClientConfig{
		authToken:  apiKey,
		BaseURL:    baseURL,
		OrgID:      "",
		APIType:    APITypeAzure,
		APIVersion: "2023-03-15-preview",
		Engine:     engine,

		HTTPClient: &http.Client{},

		EmptyMessagesLimit: defaultEmptyMessagesLimit,
	}
}

func (ClientConfig) String() string {
	return "<OpenAI API ClientConfig>"
}

// UseProxyClient 使用代理
func (c ClientConfig) UseProxyClient(proxy *http.Client) {
	c.HTTPClient = proxy
	return
}

// UseProxy 使用代理，指定代理地址
func (c ClientConfig) UseProxy(proxyUrl string) {
	proxyURL, err := url.Parse(proxyUrl)
	if err != nil {
		panic(err)
	}

	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}

	client := &http.Client{
		Transport: transport,
	}
	c.HTTPClient = client
	return
}

// UseSocket5Proxy 使用socket5代理
func (c ClientConfig) UseSocket5Proxy(ip string, port int) {
	c.UseProxy(fmt.Sprintf("socks5://%s:%d", ip, port))
}
