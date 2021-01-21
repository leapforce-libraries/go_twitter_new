package twitter

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	errortools "github.com/leapforce-libraries/go_errortools"
	oauth2 "github.com/leapforce-libraries/go_oauth2"
)

const (
	APIURL         string = "https://api.twitter.com/2"
	AccessTokenURL string = "https://api.twitter.com/oauth2/token?grant_type=client_credentials"
)

// type
//
type Service struct {
	basicAuthorization string
	oAuth2             *oauth2.OAuth2
}

type ServiceConfig struct {
	ConsumerKey           string
	ConsumerSecret        string
	MaxRetries            *uint
	SecondsBetweenRetries *uint32
}

func NewService(serviceConfig ServiceConfig) (*Service, *errortools.Error) {
	if serviceConfig.ConsumerKey == "" {
		return nil, errortools.ErrorMessage("ConsumerKey not provided")
	}

	if serviceConfig.ConsumerSecret == "" {
		return nil, errortools.ErrorMessage("ConsumerSecret not provided")
	}

	service := Service{
		basicAuthorization: fmt.Sprintf("Basic %s", base64.URLEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", serviceConfig.ConsumerKey, serviceConfig.ConsumerSecret)))),
	}

	tokenFunction := func() (*oauth2.Token, *errortools.Error) {
		return service.GetAccessToken()
	}

	oAuth2Config := oauth2.OAuth2Config{
		NewTokenFunction:      &tokenFunction,
		MaxRetries:            serviceConfig.MaxRetries,
		SecondsBetweenRetries: serviceConfig.SecondsBetweenRetries,
	}
	service.oAuth2 = oauth2.NewOAuth(oAuth2Config)
	return &service, nil
}

/*
func (service *Service) ValidateToken() (*oauth2.Token, *errortools.Error) {
	return service.oAuth2.ValidateToken()
}*/

// generic Get method
//
func (service *Service) get(requestConfig *oauth2.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	return service.httpRequest(http.MethodGet, requestConfig)
}

// generic Post method
//
func (service *Service) post(requestConfig *oauth2.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	return service.httpRequest(http.MethodPost, requestConfig)
}

// generic Put method
//
func (service *Service) put(requestConfig *oauth2.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	return service.httpRequest(http.MethodPut, requestConfig)
}

// generic Patch method
//
func (service *Service) patch(requestConfig *oauth2.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	return service.httpRequest(http.MethodPatch, requestConfig)
}

// generic Delete method
//
func (service *Service) delete(requestConfig *oauth2.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	return service.httpRequest(http.MethodDelete, requestConfig)
}

func (service *Service) url(path string) string {
	return fmt.Sprintf("%s/%s", APIURL, path)
}

func (service *Service) httpRequest(httpMethod string, requestConfig *oauth2.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	errorResponse := ErrorResponse{}
	(*requestConfig).ErrorModel = &errorResponse

	request, response, e := service.oAuth2.HTTP(httpMethod, requestConfig)

	if e != nil {
		if errorResponse.Detail != "" {
			e.SetMessage(errorResponse.Detail)
		}

		b, _ := json.Marshal(errorResponse)
		e.SetExtra("error", string(b))
	}

	return request, response, e
}
