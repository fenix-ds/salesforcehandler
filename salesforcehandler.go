package salesforcehandler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func NewSalesForceHandler(param SalesForceParam) (*SalesForceHandler, error) {
	if err := param.checkdata(); err != nil {
		return nil, err
	}

	accessToken, err := _SalesForceLogin(&param.Urls.Autentication, &param.Autentication)
	if err != nil {
		return nil, err
	}

	return &SalesForceHandler{
		urls:        &param.Urls,
		accessToken: accessToken,
	}, nil
}

func _SalesForceLogin(urlAddress *string, param *SalesForceAutentication) (*string, error) {
	if urlAddress == nil || param == nil {
		return nil, fmt.Errorf("login details for salesforce not found.")
	} else if len(*urlAddress) == 0 {
		return nil, fmt.Errorf("url for authentication not found.")
	}

	formLogin := url.Values{}

	formLogin.Set("grant_type", param.GrantType)
	formLogin.Set("client_id", param.ClientId)
	formLogin.Set("client_secret", param.ClientSecret)
	formLogin.Set("username", param.UserName)
	formLogin.Set("password", param.Password)

	req, err := http.NewRequest(http.MethodPost, *urlAddress, strings.NewReader(formLogin.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	accessToken, ok := result["access_token"].(string)
	if ok {
		return &accessToken, nil
	} else {
		return nil, fmt.Errorf("error retrieving access token")
	}
}
