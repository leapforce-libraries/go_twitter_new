package twitter

import (
	"net/http"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	oauth2 "github.com/leapforce-libraries/go_oauth2"
)

// Token stures Token object
//
type AccessToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

func (service *Service) GetOauth2AccessToken() (*oauth2.Token, *errortools.Error) {
	if service.oAuth2Service == nil {
		return nil, errortools.ErrorMessage("OAuth2 not initialized")
	}

	accessToken := AccessToken{}

	// basic authentication header
	header := http.Header{}
	header.Set("Authorization", service.basicAuthorization)

	requestConfig := go_http.RequestConfig{
		Method:            http.MethodPost,
		URL:               accessTokenURL2,
		ResponseModel:     &accessToken,
		NonDefaultHeaders: &header,
	}

	_, _, e := service.oAuth2Service.HTTPRequestWithoutAccessToken(&requestConfig)
	if e != nil {
		return nil, e
	}

	token := oauth2.Token{
		AccessToken: &accessToken.AccessToken,
	}

	return &token, nil
}
