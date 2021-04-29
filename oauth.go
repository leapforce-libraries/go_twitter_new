package twitter

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

func (service *Service) GetOauthToken(redirectURL string) *errortools.Error {
	params := url.Values{}
	params.Set(_OauthCallback, redirectURL)
	params.Set(_OauthConsumerKey, service.consumerKey)

	requestConfig := go_http.RequestConfig{
		URL: fmt.Sprintf("%s?%s", requestTokenURL, params.Encode()),
	}

	_, response, e := service.httpService.HTTPRequest(requestTokenHTTPMethod, &requestConfig)
	if e != nil {
		return e
	}

	if response == nil {
		return errortools.ErrorMessage("Response is nil")
	}
	if response.Body == nil {
		return errortools.ErrorMessage("Response body is nil")
	}

	defer response.Body.Close()
	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return errortools.ErrorMessage(err)
	}
	values, err := url.ParseQuery(string(b))
	if err != nil {
		return errortools.ErrorMessage(err)
	}

	confirmed := values.Get(_OauthCallbackConfirmed)
	if confirmed != "true" {
		return errortools.ErrorMessagef("oauth_callback_confirmed is '%s' (not 'true')", confirmed)
	}

	_oauthToken := values.Get(_OauthToken)
	if _oauthToken == "" {
		return errortools.ErrorMessagef("Response does not contain '%s' value", _OauthToken)
	}

	service.oauthToken = _oauthToken

	return nil
}

func (service *Service) AuthorizeURL() string {
	return fmt.Sprintf("%s?%s=%s", authorizeURL, _OauthToken, service.oauthToken)
}

func (service *Service) GetAccessToken(r *http.Request) *errortools.Error {
	if service.oauthToken != r.URL.Query().Get(_OauthToken) {
		return errortools.ErrorMessage("OAuth token verification failed")
	}

	service.oauthVerifier = r.URL.Query().Get(_OauthVerifier)

	// STEP 3: Convert the request token into a usable access token
	params := url.Values{}
	params.Set(_OauthConsumerKey, service.consumerKey)
	params.Set(_OauthToken, service.oauthToken)
	params.Set(_OauthVerifier, service.oauthVerifier)

	requestConfig := go_http.RequestConfig{
		URL: fmt.Sprintf("%s?%s", accessTokenURL, params.Encode()),
	}

	// use new http service (without OAuth1.0 applied)
	httpService, e := go_http.NewService(&go_http.ServiceConfig{
		HTTPClient: &http.Client{},
	})
	if e != nil {
		return errortools.ErrorMessage(e.Message())
	}

	_, response2, e := httpService.HTTPRequest(accessTokenHTTPMethod, &requestConfig)
	if e != nil {
		return errortools.ErrorMessage(e.Message())
	}

	defer response2.Body.Close()
	b, err := ioutil.ReadAll(response2.Body)
	if err != nil {
		return errortools.ErrorMessage(err)
	}

	values, err := url.ParseQuery(string(b))
	if err != nil {
		return errortools.ErrorMessage(err)
	}

	service.accessToken = values.Get(_OauthToken)
	service.accessSecret = values.Get(_OauthTokenSecret)

	return nil
}

func (service *Service) AccessTokenSecret() (string, string) {
	return service.accessToken, service.accessSecret
}
