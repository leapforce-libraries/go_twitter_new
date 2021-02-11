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

	"github.com/dghubble/oauth1"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	oauth2 "github.com/leapforce-libraries/go_oauth2"
	ratelimit "github.com/leapforce-libraries/go_ratelimit"
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
	httpService        *go_http.Service
	oAuth2             *oauth2.OAuth2
	rateLimitService   *ratelimit.Service
}

type ServiceConfigOAuth1 struct {
	ConsumerKey           string
	ConsumerSecret        string
	AccessToken           string
	AccessSecret          string
	MaxRetries            *uint
	SecondsBetweenRetries *uint32
}

type ServiceConfigOAuth2 struct {
	ConsumerKey           string
	ConsumerSecret        string
	MaxRetries            *uint
	SecondsBetweenRetries *uint32
}

func NewServiceOAuth1(serviceConfig ServiceConfigOAuth1) (*Service, *errortools.Error) {
	if serviceConfig.ConsumerKey == "" {
		return nil, errortools.ErrorMessage("ConsumerKey not provided")
	}

	if serviceConfig.ConsumerSecret == "" {
		return nil, errortools.ErrorMessage("ConsumerSecret not provided")
	}

	if serviceConfig.AccessToken == "" {
		return nil, errortools.ErrorMessage("AccessToken not provided")
	}

	if serviceConfig.AccessSecret == "" {
		return nil, errortools.ErrorMessage("AccessSecret not provided")
	}

	// create (organisation specific) Service
	config := oauth1.NewConfig(serviceConfig.ConsumerKey, serviceConfig.ConsumerSecret)
	token := oauth1.NewToken(serviceConfig.AccessToken, serviceConfig.AccessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	httpServiceConfig := go_http.ServiceConfig{
		HTTPClient:            httpClient,
		MaxRetries:            serviceConfig.MaxRetries,
		SecondsBetweenRetries: serviceConfig.SecondsBetweenRetries,
	}

	headerRemaining := "x-rate-limit-remaining"
	headerReset := "x-rate-limit-reset"
	rateLimitServiceConfig := ratelimit.ServiceConfig{
		HeaderRemaining: &headerRemaining,
		HeaderReset:     &headerReset,
	}

	return &Service{
		httpService:      go_http.NewService(httpServiceConfig),
		rateLimitService: ratelimit.NewService(&rateLimitServiceConfig),
	}, nil
}

func NewServiceOAuth2(serviceConfig ServiceConfigOAuth2) (*Service, *errortools.Error) {
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

// generic Get method
//
func (service *Service) get(requestConfig *go_http.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	return service.httpRequest(http.MethodGet, requestConfig)
}

// generic Post method
//
func (service *Service) post(requestConfig *go_http.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	return service.httpRequest(http.MethodPost, requestConfig)
}

// generic Put method
//
func (service *Service) put(requestConfig *go_http.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	return service.httpRequest(http.MethodPut, requestConfig)
}

// generic Patch method
//
func (service *Service) patch(requestConfig *go_http.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	return service.httpRequest(http.MethodPatch, requestConfig)
}

// generic Delete method
//
func (service *Service) delete(requestConfig *go_http.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	return service.httpRequest(http.MethodDelete, requestConfig)
}

func (service *Service) url(path string) string {
	return fmt.Sprintf("%s/%s", APIURL, path)
}

func (service *Service) httpRequest(httpMethod string, requestConfig *go_http.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	errorResponse := ErrorResponse{}
	(*requestConfig).ErrorModel = &errorResponse

	request := new(http.Request)
	response := new(http.Response)
	e := new(errortools.Error)

	if service.httpService != nil {
		request, response, e = service.httpService.HTTPRequest(httpMethod, requestConfig)
	} else {
		request, response, e = service.oAuth2.HTTPRequest(httpMethod, requestConfig, false)
	}

	if response != nil {
		if response.StatusCode == 429 {
			// handle rate limit
			rateLimitReset, err := strconv.ParseInt(response.Header.Get("x-rate-limit-reset"), 10, 64)
			if err == nil {
				rateLimitResetUnix := time.Unix(rateLimitReset, 0)
				duration := rateLimitResetUnix.Sub(time.Now())

				if duration > 0 {
					errortools.CaptureInfo(fmt.Sprintf("Rate limit exceeded, waiting %v ms.", duration.Milliseconds()))
					time.Sleep(duration)

					return service.httpRequest(httpMethod, requestConfig)
				}
			}
		}
	}

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
