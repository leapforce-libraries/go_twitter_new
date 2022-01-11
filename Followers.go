package twitter

import (
	"encoding/json"
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	models "github.com/leapforce-libraries/go_twitter_new/models"
)

type FollowersResponse struct {
	Data   *[]models.User  `json:"data"`
	Meta   *models.Meta    `json:"meta"`
	Errors *[]models.Error `json:"errors"`
}

type GetFollowersCall struct {
	service         *Service
	userID          string
	Expansions      *[]string `tw:"expansions"`
	MaxResults      *int      `tw:"max_results"`
	PaginationToken *string   `tw:"pagination_token"`
	TweetFields     *[]string `tw:"tweet.fields"`
	UserFields      *[]string `tw:"user.fields"`
}

func (service *Service) NewGetFollowersCall(userID string) *GetFollowersCall {
	return &GetFollowersCall{
		service: service,
		userID:  userID,
	}
}

func (call *GetFollowersCall) SetExpansions(expansions ...UserExpansion) *GetFollowersCall {
	elems := []string{}

	for _, elem := range expansions {
		elems = append(elems, string(elem))
	}
	(*call).Expansions = &elems

	return call
}

func (call *GetFollowersCall) SetMaxResults(maxResults int) *GetFollowersCall {
	(*call).MaxResults = &maxResults

	return call
}

func (call *GetFollowersCall) SetTweetFields(tweetFields ...TweetField) *GetFollowersCall {
	return call.setTweetFields(false, tweetFields)
}

func (call *GetFollowersCall) AddTweetFields(tweetFields ...TweetField) *GetFollowersCall {
	return call.setTweetFields(true, tweetFields)
}

func (call *GetFollowersCall) setTweetFields(add bool, tweetFields []TweetField) *GetFollowersCall {
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

func (call *GetFollowersCall) SetUserFields(userFields ...UserField) *GetFollowersCall {
	return call.setUserFields(false, userFields)
}

func (call *GetFollowersCall) AddUserFields(userFields ...UserField) *GetFollowersCall {
	return call.setUserFields(true, userFields)
}

func (call *GetFollowersCall) setUserFields(add bool, userFields []UserField) *GetFollowersCall {
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

func (call *GetFollowersCall) Do() (*[]models.User, *errortools.Error) {
	followers := []models.User{}

	rowCount := 0

	for {
		params, e := call.service.urlParams(call)
		if e != nil {
			return nil, e
		}

		urlPath := fmt.Sprintf("users/%s/followers%s", call.userID, *params)
		//fmt.Println(urlPath)

		followersResponse := FollowersResponse{}
		requestConfig := go_http.RequestConfig{
			URL:           call.service.url(urlPath),
			ResponseModel: &followersResponse,
		}

		endpoint := "followers"
		call.service.rateLimitService.Check(endpoint)

		request, response, e := call.service.get(&requestConfig)
		if e != nil {
			return nil, e
		}

		if followersResponse.Errors != nil {
			e := new(errortools.Error)
			e.SetRequest(request)
			e.SetResponse(response)

			b, err := json.Marshal(followersResponse.Errors)
			if err == nil {
				e.SetExtra("errors", string(b))
			}

			e.SetMessage(fmt.Sprintf("%v errors found", len(*followersResponse.Errors)))
			errortools.CaptureError(e)
		}

		if followersResponse.Data == nil {
			break
		}

		rowCountCall := len(*followersResponse.Data)

		if rowCountCall > 0 {
			rowCount += rowCountCall
		}

		followers = append(followers, (*followersResponse.Data)...)

		call.service.rateLimitService.Set(endpoint, response)

		if followersResponse.Meta == nil {
			break
		}

		if followersResponse.Meta.NextToken == nil {
			break
		}

		call.PaginationToken = followersResponse.Meta.NextToken
	}

	return &followers, nil
}
