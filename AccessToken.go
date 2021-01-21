package twitter

import (
	"net/http"

	errortools "github.com/leapforce-libraries/go_errortools"
	oauth2 "github.com/leapforce-libraries/go_oauth2"
)

// Token stures Token object
//
type AccessToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

func (service *Service) GetAccessToken() (*oauth2.Token, *errortools.Error) {
	accessToken := AccessToken{}

	skipAccessToken := true

	// basic authentication header
	header := http.Header{}
	header.Set("Authorization", service.basicAuthorization)

	requestConfig := oauth2.RequestConfig{
		URL:               AccessTokenURL,
		ResponseModel:     &accessToken,
		SkipAccessToken:   &skipAccessToken,
		NonDefaultHeaders: &header,
	}

	_, _, e := service.post(&requestConfig)
	if e != nil {
		return nil, e
	}

	token := oauth2.Token{
		AccessToken: &accessToken.AccessToken,
	}

	return &token, nil
}
