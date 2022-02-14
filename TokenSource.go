package twitter

import (
	"net/http"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	oauth2 "github.com/leapforce-libraries/go_oauth2"
	token "github.com/leapforce-libraries/go_oauth2/token"
)

type TokenSource struct {
	token              *token.Token
	basicAuthorization string
	oAuth2Service      *oauth2.Service
}

func NewTokenSource(oAuth2Service *oauth2.Service, basicAuthorization string) (*TokenSource, *errortools.Error) {
	if oAuth2Service == nil {
		return nil, errortools.ErrorMessage("oAuth2Service is a nil pointer")
	}

	return &TokenSource{
		basicAuthorization: basicAuthorization,
		oAuth2Service:      oAuth2Service,
	}, nil
}

func (ts *TokenSource) Token() *token.Token {
	return ts.token
}

func (ts *TokenSource) NewToken() (*token.Token, *errortools.Error) {
	type AccessToken struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
	}

	accessToken := AccessToken{}

	// basic authentication header
	header := http.Header{}
	header.Set("Authorization", ts.basicAuthorization)

	requestConfig := go_http.RequestConfig{
		Method:            http.MethodPost,
		URL:               accessTokenURL2,
		ResponseModel:     &accessToken,
		NonDefaultHeaders: &header,
	}

	_, _, e := ts.oAuth2Service.HTTPRequestWithoutAccessToken(&requestConfig)
	if e != nil {
		return nil, e
	}

	token := token.Token{
		AccessToken: &accessToken.AccessToken,
	}

	return &token, nil
}

func (ts *TokenSource) SetToken(*token.Token, bool) *errortools.Error {
	return nil
}

func (ts *TokenSource) RetrieveToken() *errortools.Error {
	return nil
}

func (ts *TokenSource) SaveToken() *errortools.Error {
	return nil
}
