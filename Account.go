package twitter

import (
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type Account struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	ScreenName string `json:"screen_name"`
}

func (service *Service) GetAccount() (*Account, *errortools.Error) {
	urlPath := "account/verify_credentials.json"

	account := Account{}
	requestConfig := go_http.RequestConfig{
		URL:           service.urlV1(urlPath),
		ResponseModel: &account,
	}

	_, _, e := service.get(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &account, nil
}
