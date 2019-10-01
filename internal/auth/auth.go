package auth

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"
)

const (
	// DefaultAuthUrl is the url used for creating auth tokens
	defaultAuthURL = "https://login.microsoftonline.com/botframework.com/oauth2/v2.0/token"
)

// Client is the underlying http client connecting to Azure's oauth token generator
type Client struct {
	*http.Client
}

// ResponseData is the http response containing our oauth token
type ResponseData struct {
	Token string `json:"auth_token"`
}

// NewClient creates a new default http client
func NewClient() *Client {
	return &Client{
		&http.Client{
			Transport: http.DefaultTransport,
		},
	}
}

// FetchToken generates an auth token from appId and secret.
func (c *Client) FetchToken(ctx context.Context, appID, secret string) (string, error) {
	formValues := url.Values{
		"grant_type":    []string{"client_credentials"},
		"client_id":     []string{appID},
		"client_secret": []string{secret},
		"scope":         []string{appID + "/.default"},
	}

	buf := bytes.NewBufferString(formValues.Encode())
	req, _ := http.NewRequestWithContext(ctx, "POST", defaultAuthURL, buf)

	resp, err := c.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	var data ResponseData
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return "", err
	}

	return data.Token, nil
}
