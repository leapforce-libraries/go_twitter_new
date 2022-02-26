package twitter

import (
	"encoding/json"
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	models "github.com/leapforce-libraries/go_twitter_new/models"
)

type UsersResponse struct {
	Data     *models.User     `json:"data"`
	Includes *models.Includes `json:"includes"`
	Errors   *[]models.Error  `json:"errors"`
}

type UserExpansion string

const (
	UserExpansionPinnedTweetID UserExpansion = "pinned_tweet_id"
)

type GetUsersCall struct {
	service     *Service
	id          string
	Expansions  *[]string `tw:"expansions"`
	TweetFields *[]string `tw:"tweet.fields"`
	UserFields  *[]string `tw:"user.fields"`
}

func (service *Service) NewGetUsersCall(id string) *GetUsersCall {
	return &GetUsersCall{
		service: service,
		id:      id,
	}
}

func (call *GetUsersCall) SetExpansions(expansions ...UserExpansion) *GetUsersCall {
	elems := []string{}

	for _, elem := range expansions {
		elems = append(elems, string(elem))
	}
	(*call).Expansions = &elems

	return call
}

func (call *GetUsersCall) SetTweetFields(tweetFields ...TweetField) *GetUsersCall {
	return call.setTweetFields(false, tweetFields)
}

func (call *GetUsersCall) AddTweetFields(tweetFields ...TweetField) *GetUsersCall {
	return call.setTweetFields(true, tweetFields)
}

func (call *GetUsersCall) setTweetFields(add bool, tweetFields []TweetField) *GetUsersCall {
	elems := []string{}

	if (*call).TweetFields != nil && add {
		elems = *(*call).TweetFields
	}

	for _, tweetField := range tweetFields {
		for _, _elem := range elems {
			if _elem == string(tweetField) {
				goto next
			}
		}
		elems = append(elems, string(tweetField))
	next:
	}
	(*call).TweetFields = &elems

	return call
}

func (call *GetUsersCall) SetUserFields(userFields ...UserField) *GetUsersCall {
	return call.setUserFields(false, userFields)
}

func (call *GetUsersCall) AddUserFields(userFields ...UserField) *GetUsersCall {
	return call.setUserFields(true, userFields)
}

func (call *GetUsersCall) setUserFields(add bool, userFields []UserField) *GetUsersCall {
	elems := []string{}

	if (*call).UserFields != nil && add {
		elems = *(*call).UserFields
	}

	for _, userField := range userFields {
		for _, _elem := range elems {
			if _elem == string(userField) {
				goto next
			}
		}
		elems = append(elems, string(userField))
	next:
	}
	(*call).UserFields = &elems

	return call
}

func (call *GetUsersCall) Do() (*models.User, *models.Includes, *errortools.Error) {
	params, e := call.service.urlParams(call)
	if e != nil {
		return nil, nil, e
	}

	urlPath := fmt.Sprintf("users/%s%s", call.id, *params)
	//fmt.Println(urlPath)

	usersResponse := UsersResponse{}
	requestConfig := go_http.RequestConfig{
		Url:           call.service.url(urlPath),
		ResponseModel: &usersResponse,
	}

	endpoint := "users"
	call.service.rateLimitService.Check(endpoint)

	request, response, e := call.service.get(&requestConfig)
	if e != nil {
		return nil, nil, e
	}

	if usersResponse.Errors != nil {
		e := new(errortools.Error)
		e.SetRequest(request)
		e.SetResponse(response)

		b, err := json.Marshal(usersResponse.Errors)
		if err == nil {
			e.SetExtra("errors", string(b))
		}

		e.SetMessage(fmt.Sprintf("%v errors found", len(*usersResponse.Errors)))
		errortools.CaptureError(e)
	}

	call.service.rateLimitService.Set(endpoint, response)

	return usersResponse.Data, usersResponse.Includes, nil
}
