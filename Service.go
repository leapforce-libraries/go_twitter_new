package twitter

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	oauth2 "github.com/leapforce-libraries/go_oauth2"
	utilities "github.com/leapforce-libraries/go_utilities"
)

const (
	APIURL            string = "https://api.twitter.com/2"
	AccessTokenURL    string = "https://api.twitter.com/oauth2/token?grant_type=client_credentials"
	DateLayoutISO8601 string = "2006-01-02T15:04:05Z"
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

func (service *Service) urlParams(model interface{}) (*string, *errortools.Error) {
	if utilities.IsNil(model) {
		return nil, nil
	}

	if reflect.TypeOf(model).Kind() != reflect.Ptr {
		return nil, errortools.ErrorMessage("The interface is not a pointer.")
	}

	p := reflect.ValueOf(model) //pointer
	s := p.Elem()               //interface

	if s.Kind() != reflect.Struct {
		s = s.Elem()
	}

	if s.Kind() != reflect.Struct {
		return nil, errortools.ErrorMessage("The interface is not a pointer to a struct.")
	}

	values := url.Values{}

	for j := 0; j < s.NumField(); j++ {
		fieldName := s.Type().Field(j).Tag.Get("tw")

		if fieldName == "" {
			continue
		}

		field := s.Field(j)

		if field.Kind() == reflect.Ptr {
			if field.IsNil() {
				continue
			}

			field = field.Elem()
		}

		value := field.Interface()

		switch v := value.(type) {
		case int:
			values.Set(fieldName, strconv.Itoa(v))
		case int64:
			values.Set(fieldName, strconv.FormatInt(v, 10))
		case string:
			values.Set(fieldName, v)
		case bool:
			values.Set(fieldName, strconv.FormatBool(v))
		case []string:
			s := []string{}
			for _, v1 := range v {
				s = append(s, v1)
			}
			values.Set(fieldName, strings.Join(s, ","))
		case time.Time:
			values.Set(fieldName, v.Format(DateLayoutISO8601))
		}
	}

	params := ""
	if len(values) > 0 {
		params = fmt.Sprintf("?%s", values.Encode())
	}

	return &params, nil
}
