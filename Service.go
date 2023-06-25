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
	apiName                string = "Twitter"
	apiUrl                 string = "https://api.twitter.com/2"
	apiUrlV1               string = "https://api.twitter.com/1.1"
	accessTokenUrl2        string = "https://api.twitter.com/oauth2/token?grant_type=client_credentials"
	authorizeUrl           string = "https://api.twitter.com/oauth/authorize"
	requestTokenUrl        string = "https://api.twitter.com/oauth/request_token"
	accessTokenUrl         string = "https://api.twitter.com/oauth/access_token"
	requestTokenHttpMethod string = http.MethodPost
	accessTokenHttpMethod  string = http.MethodPost
	dateLayoutIso8601      string = "2006-01-02T15:04:05Z"
	redirectUrl            string = "http://localhost:8080/oauth/redirect"
)

const (
	_OauthCallback          string = "oauth_callback"
	_OauthCallbackConfirmed string = "oauth_callback_confirmed"
	_OauthConsumerKey       string = "oauth_consumer_key"
	_OauthToken             string = "oauth_token"
	_OauthTokenSecret       string = "oauth_token_secret"
	_OauthVerifier          string = "oauth_verifier"
)

// type
type Service struct {
	consumerKey string
	//basicAuthorization string
	httpService      *go_http.Service
	oAuth2Service    *oauth2.Service
	rateLimitService *ratelimit.Service
	oauthToken       string
	oauthVerifier    string
	accessToken      string
	accessSecret     string
}

type ServiceConfigOAuth1 struct {
	ConsumerKey    string
	ConsumerSecret string
	AccessToken    string
	AccessSecret   string
}

type ServiceConfigOAuth2 struct {
	ConsumerKey    string
	ConsumerSecret string
}

func NewServiceNoOAuth(consumerKey string) (*Service, *errortools.Error) {
	if consumerKey == "" {
		return nil, errortools.ErrorMessage("ConsumerKey not provided")
	}

	httpService, e := go_http.NewService(&go_http.ServiceConfig{})
	if e != nil {
		return nil, e
	}

	return &Service{
		consumerKey: consumerKey,
		httpService: httpService,
	}, nil
}

func NewServiceOAuth1(serviceConfig *ServiceConfigOAuth1) (*Service, *errortools.Error) {
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

	// create Service
	config := oauth1.NewConfig(serviceConfig.ConsumerKey, serviceConfig.ConsumerSecret)
	token := oauth1.NewToken(serviceConfig.AccessToken, serviceConfig.AccessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	httpServiceConfig := go_http.ServiceConfig{
		HttpClient: httpClient,
	}

	httpService, e := go_http.NewService(&httpServiceConfig)
	if e != nil {
		return nil, e
	}

	headerRemaining := "x-rate-limit-remaining"
	headerReset := "x-rate-limit-reset"
	rateLimitServiceConfig := ratelimit.ServiceConfig{
		HeaderRemaining: &headerRemaining,
		HeaderReset:     &headerReset,
	}

	return &Service{
		consumerKey:      serviceConfig.ConsumerKey,
		httpService:      httpService,
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

	basicAuthorization := fmt.Sprintf("Basic %s", base64.URLEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", serviceConfig.ConsumerKey, serviceConfig.ConsumerSecret))))

	oAuth2ServiceBasic, e := oauth2.NewService(&oauth2.ServiceConfig{})
	if e != nil {
		return nil, e
	}
	tokenSource, e := NewTokenSource(oAuth2ServiceBasic, basicAuthorization)
	if e != nil {
		return nil, e
	}

	oAuth2ServiceConfig := oauth2.ServiceConfig{
		TokenSource: tokenSource,
	}
	oAuth2Service, e := oauth2.NewService(&oAuth2ServiceConfig)
	if e != nil {
		return nil, e
	}

	return &Service{
		oAuth2Service: oAuth2Service,
	}, nil
}

// generic Get method
func (service *Service) get(requestConfig *go_http.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	return service.httpRequest(http.MethodGet, requestConfig)
}

func (service *Service) url(path string) string {
	return fmt.Sprintf("%s/%s", apiUrl, path)
}

func (service *Service) urlV1(path string) string {
	return fmt.Sprintf("%s/%s", apiUrlV1, path)
}

func (service *Service) httpRequest(httpMethod string, requestConfig *go_http.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	errorResponse := ErrorResponse{}
	(*requestConfig).ErrorModel = &errorResponse

	var request *http.Request = nil
	var response *http.Response = nil
	var e *errortools.Error = nil

	if service.httpService != nil {
		request, response, e = service.httpService.HttpRequest(requestConfig)
	} else {
		request, response, e = service.oAuth2Service.HttpRequest(requestConfig)
	}

	if response != nil {
		if response.StatusCode == 429 {
			// handle rate limit
			rateLimitReset, err := strconv.ParseInt(response.Header.Get("x-rate-limit-reset"), 10, 64)
			if err == nil {
				rateLimitResetUnix := time.Unix(rateLimitReset, 0)
				duration := time.Until(rateLimitResetUnix)

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
			values.Set(fieldName, strings.Join(v, ","))
		case time.Time:
			values.Set(fieldName, v.Format(dateLayoutIso8601))
		}
	}

	params := ""
	if len(values) > 0 {
		params = fmt.Sprintf("?%s", values.Encode())
	}

	return &params, nil
}

func (service *Service) InitToken() *errortools.Error {
	if service == nil {
		return errortools.ErrorMessage("Service is nil pointer")
	}

	// STEP 1: Create a request for a consumer application to obtain a request token
	e := service.GetOauthToken(redirectUrl)
	if e != nil {
		return e
	}

	// STEP 2: Let the user authenticate and send the consumer application a request token
	fmt.Printf("Go to this url to get new access token:\n\n%s\n\n", service.AuthorizeUrl1())

	// Create a new redirect route
	http.HandleFunc("/oauth/redirect", func(w http.ResponseWriter, r *http.Request) {

		// STEP 3: exchange request token by access token
		e := service.GetAccessToken(r)
		if e != nil {
			fmt.Println(e.Message())
			w.WriteHeader(http.StatusBadRequest)
		}

		w.WriteHeader(http.StatusFound)
	})

	http.ListenAndServe(":8080", nil)

	return nil
}

func (service *Service) ApiName() string {
	return apiName
}

func (service *Service) ApiKey() string {
	return service.consumerKey
}

func (service *Service) ApiCallCount() int64 {
	if service.httpService != nil {
		return service.httpService.RequestCount()
	} else {
		return service.oAuth2Service.ApiCallCount()
	}
}

func (service *Service) ApiReset() {
	if service.httpService != nil {
		service.httpService.ResetRequestCount()
	} else {
		service.oAuth2Service.ApiReset()
	}
}
