package twitter

import (
	"encoding/base64"
	"fmt"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	oauth2 "github.com/leapforce-libraries/go_oauth2"
	"github.com/leapforce-libraries/go_oauth2/tokensource"
	ratelimit "github.com/leapforce-libraries/go_ratelimit"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	/*apiName                string = "Twitter"
	apiUrl                 string = "https://api.twitter.com/2"
	apiUrlV1               string = "https://api.twitter.com/1.1"
	accessTokenUrl2        string = "https://api.twitter.com/2/oauth2/token"
	authorizeUrl           string = "https://twitter.com/i/oauth2/authorize"
	requestTokenUrl        string = "https://api.twitter.com/oauth/request_token"
	accessTokenUrl         string = "https://api.twitter.com/oauth/access_token"
	requestTokenHttpMethod string = http.MethodPost
	accessTokenHttpMethod  string = http.MethodPost
	dateLayoutIso8601      string = "2006-01-02T15:04:05Z"*/
	defaultRedirectUrl string = "http://localhost:8080/oauth/redirect"
)

type Service2 struct {
	consumerKey  string
	clientId     string
	clientSecret string
	//basicAuthorization string
	httpService      *go_http.Service
	oAuth2Service    *oauth2.Service
	redirectUrl      *string
	errorResponse    *ErrorResponse
	rateLimitService *ratelimit.Service
	oauthToken       string
	oauthVerifier    string
	accessToken      string
	accessSecret     string
}

type Service2ConfigOAuth1 struct {
	ConsumerKey    string
	ConsumerSecret string
	AccessToken    string
	AccessSecret   string
}

type Service2ConfigOAuth2 struct {
	ClientId     string
	ClientSecret string
	TokenSource  tokensource.TokenSource
	RedirectUrl  *string
}

func (service *Service2) getTokenRequest(r *http.Request) (*http.Request, *errortools.Error) {
	err := r.ParseForm()
	if err != nil {
		return nil, errortools.ErrorMessage(err)
	}
	code := r.FormValue("code")

	data := url.Values{}
	data.Set("code", code)
	data.Set("grant_type", "authorization_code")
	data.Set("client_id", service.clientId)
	if service.redirectUrl != nil {
		data.Set("redirect_uri", *service.redirectUrl)
	}
	data.Set("code_verifier", "1234")

	encoded := data.Encode()
	body := strings.NewReader(encoded)

	req, err := http.NewRequest(http.MethodPost, accessTokenUrl2, body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Content-Length", strconv.Itoa(len(encoded)))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", service.clientId, service.clientSecret)))))
	if err != nil {
		return nil, errortools.ErrorMessage(err)
	}

	return req, nil
}

func NewService2OAuth2(serviceConfig *Service2ConfigOAuth2) (*Service2, *errortools.Error) {
	if serviceConfig == nil {
		return nil, errortools.ErrorMessage("ServiceConfig must not be a nil pointer")
	}

	redirectUrl := defaultRedirectUrl
	if serviceConfig.RedirectUrl != nil {
		redirectUrl = *serviceConfig.RedirectUrl
	}

	var service = Service2{
		clientId:     serviceConfig.ClientId,
		clientSecret: serviceConfig.ClientSecret,
		redirectUrl:  &redirectUrl,
	}

	var getTokenRequestFunc = service.getTokenRequest

	oAuth2ServiceConfig := oauth2.ServiceConfig{
		ClientId:                serviceConfig.ClientId,
		ClientSecret:            serviceConfig.ClientSecret,
		RedirectUrl:             redirectUrl,
		AuthUrl:                 authorizeUrl,
		TokenUrl:                accessTokenUrl2,
		TokenHttpMethod:         accessTokenHttpMethod,
		TokenSource:             serviceConfig.TokenSource,
		GetTokenFromRequestFunc: &getTokenRequestFunc,
	}
	oAuth2Service, e := oauth2.NewService(&oAuth2ServiceConfig)
	if e != nil {
		return nil, e
	}

	service.oAuth2Service = oAuth2Service

	return &service, nil
}
func (service *Service2) AuthorizeUrl(scope string, state string) string {
	t := &url.URL{Path: scope}
	var url = service.oAuth2Service.AuthorizeUrl(nil, nil, nil, &state)
	return fmt.Sprintf("%s&scope=%s&code_challenge=1234&code_challenge_method=plain", url, t.String())
}

func (service *Service2) ValidateToken() {
	t, _ := service.oAuth2Service.ValidateToken()
	fmt.Printf("%+v\n", *t)
}

func (service *Service2) GetTokenFromCode(r *http.Request, checkState *func(state string) *errortools.Error) *errortools.Error {
	return service.oAuth2Service.GetTokenFromCode(r, checkState)
}
