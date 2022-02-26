package twitter

import (
	"encoding/json"
	"fmt"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	models "github.com/leapforce-libraries/go_twitter_new/models"
)

const (
	maximumNumberOfTweetIDsPerCall int = 100
)

type TweetsResponse struct {
	Data     *[]models.Tweet  `json:"data"`
	Includes *models.Includes `json:"includes"`
	Meta     *models.Meta     `json:"meta"`
	Errors   *[]models.Error  `json:"errors"`
}

type Exclude string

const (
	ExcludeRetweets Exclude = "retweets"
	ExcludeReplies  Exclude = "replies"
)

type TweetExpansion string

const (
	ExpansionAttachmentsPollIDs         TweetExpansion = "attachments.poll_ids"
	ExpansionAttachmentsMediaKeys       TweetExpansion = "attachments.media_keys"
	ExpansionAuthorID                   TweetExpansion = "author_id"
	ExpansionEntitiesMentionsUsername   TweetExpansion = "entities.mentions.username"
	ExpansionGeoPlaceID                 TweetExpansion = "geo.place_id"
	ExpansionInReplyToUserID            TweetExpansion = "in_reply_to_user_id"
	ExpansionReferencedTweetsID         TweetExpansion = "referenced_tweets.id"
	ExpansionReferencedTweetsIDAuthorID TweetExpansion = "referenced_tweets.id.author_id"
)

type MediaField string

const (
	MediaFieldDurationMS       MediaField = "duration_ms"
	MediaFieldHeight           MediaField = "height"
	MediaFieldMediaKey         MediaField = "media_key"
	MediaFieldPreviewImageUrl  MediaField = "preview_image_url"
	MediaFieldType             MediaField = "type"
	MediaFieldUrl              MediaField = "url"
	MediaFieldWidth            MediaField = "width"
	MediaFieldPublicMetrics    MediaField = "public_metrics"
	MediaFieldNonPublicMetrics MediaField = "non_public_metrics"
	MediaFieldOrganicMetrics   MediaField = "organic_metrics"
	MediaFieldPromotedMetrics  MediaField = "promoted_metrics"
)

type PlaceField string

const (
	PlaceFieldContainedWithin PlaceField = "contained_within"
	PlaceFieldCountry         PlaceField = "country"
	PlaceFieldCountryCode     PlaceField = "country_code"
	PlaceFieldFullName        PlaceField = "full_name"
	PlaceFieldGeo             PlaceField = "geo"
	PlaceFieldID              PlaceField = "id"
	PlaceFieldName            PlaceField = "name"
	PlaceFieldPlaceType       PlaceField = "place_type"
)

type PollField string

const (
	PollFieldDurationMinutes PollField = "duration_minutes"
	PollFieldEndDateTime     PollField = "end_datetime"
	PollFieldID              PollField = "id"
	PollFieldOptions         PollField = "options"
	PollFieldVotingStatus    PollField = "voting_status"
)

type TweetField string

const (
	TweetFieldAttachments        TweetField = "attachments"
	TweetFieldAuthorID           TweetField = "author_id"
	TweetFieldContextAnnotations TweetField = "context_annotations"
	TweetFieldConversationID     TweetField = "conversation_id"
	TweetFieldCreatedAt          TweetField = "created_at"
	TweetFieldEntities           TweetField = "entities"
	TweetFieldGeo                TweetField = "geo"
	TweetFieldID                 TweetField = "id"
	TweetFieldInReplyToUserID    TweetField = "in_reply_to_user_id"
	TweetFieldLanguage           TweetField = "lang"
	TweetFieldNonPublicMetrics   TweetField = "non_public_metrics"
	TweetFieldPublicMetrics      TweetField = "public_metrics"
	TweetFieldOrganicMetrics     TweetField = "organic_metrics"
	TweetFieldPromotedMetrics    TweetField = "promoted_metrics"
	TweetFieldPossiblySensitive  TweetField = "possibly_sensitive"
	TweetFieldReferencedTweets   TweetField = "referenced_tweets"
	TweetFieldReplySettings      TweetField = "reply_settings"
	TweetFieldSource             TweetField = "source"
	TweetFieldText               TweetField = "text"
	TweetFieldWithheld           TweetField = "withheld"
)

type UserField string

const (
	UserFieldCreatedAt       UserField = "created_at"
	UserFieldDescription     UserField = "description"
	UserFieldEntities        UserField = "entities"
	UserFieldID              UserField = "id"
	UserFieldLocation        UserField = "location"
	UserFieldName            UserField = "name"
	UserFieldPinnedTweetID   UserField = "pinned_tweet_id"
	UserFieldProfileImageUrl UserField = "profile_image_url"
	UserFieldProtected       UserField = "protected"
	UserFieldPublicMetrics   UserField = "public_metrics"
	UserFieldUrl             UserField = "url"
	UserFieldUsername        UserField = "username"
	UserFieldVerified        UserField = "verified"
	UserFieldWithheld        UserField = "withheld"
)

type GetUserTweetsCall struct {
	service         *Service
	userID          string
	EndTime         *time.Time `tw:"end_time"`
	Exclude         *[]string  `tw:"exclude"`
	Expansions      *[]string  `tw:"expansions"`
	MaxResults      *int       `tw:"max_results"`
	MediaFields     *[]string  `tw:"media.fields"`
	PaginationToken *string    `tw:"pagination_token"`
	PlaceFields     *[]string  `tw:"place.fields"`
	PollFields      *[]string  `tw:"poll.fields"`
	SinceID         *string    `tw:"since_id"`
	StartTime       *time.Time `tw:"start_time"`
	TweetFields     *[]string  `tw:"tweet.fields"`
	UntilID         *string    `tw:"until_id"`
	UserFields      *[]string  `tw:"user.fields"`
}

func (service *Service) NewGetUserTweetsCall(userID string) *GetUserTweetsCall {
	return &GetUserTweetsCall{
		service: service,
		userID:  userID,
	}
}

func (call *GetUserTweetsCall) SetEndTime(endTime time.Time) *GetUserTweetsCall {
	(*call).EndTime = &endTime

	return call
}

func (call *GetUserTweetsCall) SetExclude(excludes ...Exclude) *GetUserTweetsCall {
	elems := []string{}

	for _, elem := range excludes {
		elems = append(elems, string(elem))
	}
	(*call).Exclude = &elems

	return call
}

func (call *GetUserTweetsCall) SetExpansions(expansions ...TweetExpansion) *GetUserTweetsCall {
	elems := []string{}

	for _, elem := range expansions {
		elems = append(elems, string(elem))
	}
	(*call).Expansions = &elems

	return call
}

func (call *GetUserTweetsCall) SetMaxResults(maxResults int) *GetUserTweetsCall {
	(*call).MaxResults = &maxResults

	return call
}

func (call *GetUserTweetsCall) SetMediaFields(mediaFields ...MediaField) *GetUserTweetsCall {
	if call.MediaFields == nil {
		call.MediaFields = &[]string{}
	}
	//return call.setMediaFields(false, mediaFields)
	setMediaFields(true, call.MediaFields, mediaFields)
	return call
}

func (call *GetUserTweetsCall) AddMediaFields(mediaFields ...MediaField) *GetUserTweetsCall {
	if call.MediaFields == nil {
		call.MediaFields = &[]string{}
	}
	//return call.setMediaFields(true, mediaFields)
	setMediaFields(true, call.MediaFields, mediaFields)
	return call
}

/*
func (call *GetUserTweetsCall) setMediaFields(add bool, mediaFields []MediaField) *GetUserTweetsCall {
	elems := []string{}

	if (*call).MediaFields != nil && add {
		elems = *(*call).MediaFields
	}

	for _, mediaField := range mediaFields {
		for _, _elem := range elems {
			if _elem == string(mediaField) {
				goto next
			}
		}
		elems = append(elems, string(mediaField))
	next:
	}
	(*call).MediaFields = &elems

	return call
}*/

func (call *GetUserTweetsCall) SetPaginationToken(paginationToken string) *GetUserTweetsCall {
	(*call).PaginationToken = &paginationToken

	return call
}

func (call *GetUserTweetsCall) SetPlaceFields(placeFields ...PlaceField) *GetUserTweetsCall {
	if call.PlaceFields == nil {
		call.PlaceFields = &[]string{}
	}
	//return call.setPlaceFields(false, placeFields)
	setPlaceFields(true, call.PlaceFields, placeFields)
	return call
}

func (call *GetUserTweetsCall) AddPlaceFields(placeFields ...PlaceField) *GetUserTweetsCall {
	if call.PlaceFields == nil {
		call.PlaceFields = &[]string{}
	}
	//return call.setPlaceFields(true, placeFields)
	setPlaceFields(true, call.PlaceFields, placeFields)
	return call
}

/*
func (call *GetUserTweetsCall) setPlaceFields(add bool, placeFields []PlaceField) *GetUserTweetsCall {
	elems := []string{}

	if (*call).PlaceFields != nil && add {
		elems = *(*call).PlaceFields
	}

	for _, placeField := range placeFields {
		for _, _elem := range elems {
			if _elem == string(placeField) {
				goto next
			}
		}
		elems = append(elems, string(placeField))
	next:
	}
	(*call).PlaceFields = &elems

	return call
}*/

func (call *GetUserTweetsCall) SetPollFields(pollFields ...PollField) *GetUserTweetsCall {
	if call.PollFields == nil {
		call.PollFields = &[]string{}
	}
	//return call.setPollFields(false, pollFields)
	setPollFields(true, call.PollFields, pollFields)
	return call
}

func (call *GetUserTweetsCall) AddPollFields(pollFields ...PollField) *GetUserTweetsCall {
	if call.PollFields == nil {
		call.PollFields = &[]string{}
	}
	//return call.setPollFields(true, pollFields)
	setPollFields(true, call.PollFields, pollFields)
	return call
}

/*
func (call *GetUserTweetsCall) setPollFields(add bool, pollFields []PollField) *GetUserTweetsCall {
	elems := []string{}

	if (*call).PollFields != nil && add {
		elems = *(*call).PollFields
	}

	for _, pollField := range pollFields {
		for _, _elem := range elems {
			if _elem == string(pollField) {
				goto next
			}
		}
		elems = append(elems, string(pollField))
	next:
	}
	(*call).PollFields = &elems

	return call
}*/

func (call *GetUserTweetsCall) SetSinceID(sinceID string) *GetUserTweetsCall {
	(*call).SinceID = &sinceID

	return call
}

func (call *GetUserTweetsCall) SetStartTime(startTime time.Time) *GetUserTweetsCall {
	(*call).StartTime = &startTime

	return call
}

func (call *GetUserTweetsCall) SetTweetFields(tweetFields ...TweetField) *GetUserTweetsCall {
	if call.TweetFields == nil {
		call.TweetFields = &[]string{}
	}
	//return call.setTweetFields(false, tweetFields)
	setTweetFields(false, call.TweetFields, tweetFields)
	return call
}

func (call *GetUserTweetsCall) AddTweetFields(tweetFields ...TweetField) *GetUserTweetsCall {
	if call.TweetFields == nil {
		call.TweetFields = &[]string{}
	}
	//return call.setTweetFields(true, tweetFields)
	setTweetFields(true, call.TweetFields, tweetFields)
	return call
}

/*
func (call *GetUserTweetsCall) setTweetFields(add bool, tweetFields []TweetField) *GetUserTweetsCall {
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
}*/

func (call *GetUserTweetsCall) SetUntilID(untilID string) *GetUserTweetsCall {
	(*call).UntilID = &untilID

	return call
}

func (call *GetUserTweetsCall) SetUserFields(userFields ...UserField) *GetUserTweetsCall {
	if call.UserFields == nil {
		call.UserFields = &[]string{}
	}
	//return call.setUserFields(false, userFields)
	setUserFields(false, call.UserFields, userFields)
	return call
}

func (call *GetUserTweetsCall) AddUserFields(userFields ...UserField) *GetUserTweetsCall {
	if call.UserFields == nil {
		call.UserFields = &[]string{}
	}
	//return call.setUserFields(true, userFields)
	setUserFields(true, call.UserFields, userFields)
	return call
}

/*
func (call *GetUserTweetsCall) setUserFields(add bool, userFields []UserField) *GetUserTweetsCall {
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
}*/

func (call *GetUserTweetsCall) Do() (*[]models.Tweet, *models.Includes, *[]string, *errortools.Error) {
	tweets := []models.Tweet{}
	includes := models.Includes{
		Tweets: &[]models.Tweet{},
		Users:  &[]models.User{},
		Places: &[]models.Place{},
		Media:  &[]models.Media{},
		Polls:  &[]models.Poll{},
	}

	rowCount := 0

	var nonExistingTweetIDs []string

	for {
		params, e := call.service.urlParams(call)
		if e != nil {
			return nil, nil, nil, e
		}

		urlPath := fmt.Sprintf("users/%s/tweets%s", call.userID, *params)
		//fmt.Println(call.service.url(urlPath))

		tweetsResponse := TweetsResponse{}
		requestConfig := go_http.RequestConfig{
			Url:           call.service.url(urlPath),
			ResponseModel: &tweetsResponse,
		}

		endpoint := "users_tweets"
		call.service.rateLimitService.Check(endpoint)

		request, response, e := call.service.get(&requestConfig)
		if e != nil {
			return nil, nil, nil, e
		}

		call.service.rateLimitService.Set(endpoint, response)
		var errors []models.Error

		if tweetsResponse.Errors != nil {
			for _, tweetError := range *tweetsResponse.Errors {
				if tweetError.Title == "Not Found Error" {
					nonExistingTweetIDs = append(nonExistingTweetIDs, tweetError.Value)
				} else {
					errors = append(errors, tweetError)
				}
			}

			if len(errors) > 0 {
				e := new(errortools.Error)
				e.SetRequest(request)
				e.SetResponse(response)

				b, err := json.Marshal(errors)
				if err == nil {
					e.SetExtra("errors", string(b))
				}

				return nil, nil, nil, errortools.ErrorMessage(fmt.Sprintf("%v errors found", len(errors)))
			}
		}

		if tweetsResponse.Data == nil {
			break
		}

		rowCountCall := len(*tweetsResponse.Data)

		if rowCountCall > 0 {
			rowCount += rowCountCall

			tweets = append(tweets, *tweetsResponse.Data...)

			if tweetsResponse.Includes != nil {
				if tweetsResponse.Includes.Tweets != nil {
					(*includes.Tweets) = append(*includes.Tweets, (*tweetsResponse.Includes.Tweets)...)
				}
				if tweetsResponse.Includes.Users != nil {
					(*includes.Users) = append(*includes.Users, (*tweetsResponse.Includes.Users)...)
				}
				if tweetsResponse.Includes.Places != nil {
					(*includes.Places) = append(*includes.Places, (*tweetsResponse.Includes.Places)...)
				}
				if tweetsResponse.Includes.Media != nil {
					(*includes.Media) = append(*includes.Media, (*tweetsResponse.Includes.Media)...)
				}
				if tweetsResponse.Includes.Polls != nil {
					(*includes.Polls) = append(*includes.Polls, (*tweetsResponse.Includes.Polls)...)
				}
			}
		}

		if tweetsResponse.Meta == nil {
			break
		}

		if tweetsResponse.Meta.NextToken == nil {
			break
		}

		call.PaginationToken = tweetsResponse.Meta.NextToken
	}

	return &tweets, &includes, &nonExistingTweetIDs, nil
}

type GetTweetsCall struct {
	service     *Service
	userID      string
	Expansions  *[]string `tw:"expansions"`
	IDs         []string  `tw:"ids"`
	MediaFields *[]string `tw:"media.fields"`
	PlaceFields *[]string `tw:"place.fields"`
	PollFields  *[]string `tw:"poll.fields"`
	TweetFields *[]string `tw:"tweet.fields"`
	UserFields  *[]string `tw:"user.fields"`
}

func (service *Service) NewGetTweetsCall(userID string) *GetTweetsCall {
	return &GetTweetsCall{
		service: service,
		userID:  userID,
	}
}

func (call *GetTweetsCall) SetExpansions(expansions ...TweetExpansion) *GetTweetsCall {
	elems := []string{}

	for _, elem := range expansions {
		elems = append(elems, string(elem))
	}
	(*call).Expansions = &elems

	return call
}

func (call *GetTweetsCall) SetIDs(ids []string) *GetTweetsCall {
	(*call).IDs = ids

	return call
}

func (call *GetTweetsCall) SetMediaFields(mediaFields ...MediaField) *GetTweetsCall {
	if call.MediaFields == nil {
		call.MediaFields = &[]string{}
	}
	setMediaFields(false, call.MediaFields, mediaFields)
	return call
}

func (call *GetTweetsCall) AddMediaFields(mediaFields ...MediaField) *GetTweetsCall {
	if call.MediaFields == nil {
		call.MediaFields = &[]string{}
	}
	setMediaFields(true, call.MediaFields, mediaFields)
	return call
}

func setMediaFields(add bool, mediaFields *[]string, setMediaFields []MediaField) {
	elems := []string{}

	if mediaFields != nil && add {
		elems = *mediaFields
	}

	for _, mediaField := range setMediaFields {
		for _, _elem := range elems {
			if _elem == string(mediaField) {
				goto next
			}
		}
		elems = append(elems, string(mediaField))
	next:
	}
	(*mediaFields) = elems
}

func (call *GetTweetsCall) SetPlaceFields(placeFields ...PlaceField) *GetTweetsCall {
	if call.PlaceFields == nil {
		call.PlaceFields = &[]string{}
	}
	setPlaceFields(false, call.PlaceFields, placeFields)
	return call
}

func (call *GetTweetsCall) AddPlaceFields(placeFields ...PlaceField) *GetTweetsCall {
	if call.PlaceFields == nil {
		call.PlaceFields = &[]string{}
	}
	setPlaceFields(true, call.PlaceFields, placeFields)
	return call
}

func setPlaceFields(add bool, placeFields *[]string, setPlaceFields []PlaceField) {
	elems := []string{}

	if placeFields != nil && add {
		elems = *placeFields
	}

	for _, placeField := range setPlaceFields {
		for _, _elem := range elems {
			if _elem == string(placeField) {
				goto next
			}
		}
		elems = append(elems, string(placeField))
	next:
	}
	(*placeFields) = elems
}

func (call *GetTweetsCall) SetPollFields(pollFields ...PollField) *GetTweetsCall {
	if call.PollFields == nil {
		call.PollFields = &[]string{}
	}
	setPollFields(false, call.PollFields, pollFields)
	return call
}

func (call *GetTweetsCall) AddPollFields(pollFields ...PollField) *GetTweetsCall {
	if call.PollFields == nil {
		call.PollFields = &[]string{}
	}
	setPollFields(true, call.PollFields, pollFields)
	return call
}

func setPollFields(add bool, pollFields *[]string, setPollFields []PollField) {
	elems := []string{}

	if pollFields != nil && add {
		elems = *pollFields
	}

	for _, pollField := range setPollFields {
		for _, _elem := range elems {
			if _elem == string(pollField) {
				goto next
			}
		}
		elems = append(elems, string(pollField))
	next:
	}
	(*pollFields) = elems
}

func (call *GetTweetsCall) SetTweetFields(tweetFields ...TweetField) *GetTweetsCall {
	if call.TweetFields == nil {
		call.TweetFields = &[]string{}
	}
	setTweetFields(false, call.TweetFields, tweetFields)
	return call
}

func (call *GetTweetsCall) AddTweetFields(tweetFields ...TweetField) *GetTweetsCall {
	if call.TweetFields == nil {
		call.TweetFields = &[]string{}
	}
	setTweetFields(true, call.TweetFields, tweetFields)
	return call
}

func setTweetFields(add bool, tweetFields *[]string, setTweetFields []TweetField) {
	elems := []string{}

	if tweetFields != nil && add {
		elems = *tweetFields
	}

	for _, tweetField := range setTweetFields {
		for _, _elem := range elems {
			if _elem == string(tweetField) {
				goto next
			}
		}
		elems = append(elems, string(tweetField))
	next:
	}
	(*tweetFields) = elems
}

func (call *GetTweetsCall) SetUserFields(userFields ...UserField) *GetTweetsCall {
	if call.UserFields == nil {
		call.UserFields = &[]string{}
	}
	setUserFields(false, call.UserFields, userFields)
	return call
}

func (call *GetTweetsCall) AddUserFields(userFields ...UserField) *GetTweetsCall {
	if call.UserFields == nil {
		call.UserFields = &[]string{}
	}
	setUserFields(true, call.UserFields, userFields)
	return call
}

func setUserFields(add bool, userFields *[]string, setUserFields []UserField) {
	elems := []string{}

	if userFields != nil && add {
		elems = *userFields
	}

	for _, userField := range setUserFields {
		for _, _elem := range elems {
			if _elem == string(userField) {
				goto next
			}
		}
		elems = append(elems, string(userField))
	next:
	}
	(*userFields) = elems
}

func (call *GetTweetsCall) Do() (*[]models.Tweet, *models.Includes, *[]string, *errortools.Error) {
	if len(call.IDs) == 0 {
		return nil, nil, nil, errortools.ErrorMessage("No TweetIDs specified")
	}

	tweets := []models.Tweet{}
	includes := models.Includes{}
	nonExistingTweetIDs := []string{}

	ids := call.IDs
	for {
		_ids := ids
		if len(ids) > maximumNumberOfTweetIDsPerCall {
			_ids = ids[:maximumNumberOfTweetIDsPerCall]
		}

		call.IDs = _ids
		params, e := call.service.urlParams(call)
		if e != nil {
			return nil, nil, nil, e
		}

		urlPath := fmt.Sprintf("tweets%s", *params)
		//fmt.Println(call.service.url(urlPath))

		tweetsResponse := TweetsResponse{}
		requestConfig := go_http.RequestConfig{
			Url:           call.service.url(urlPath),
			ResponseModel: &tweetsResponse,
		}

		endpoint := "tweets"
		call.service.rateLimitService.Check(endpoint)

		request, response, e := call.service.get(&requestConfig)
		if e != nil {
			return nil, nil, nil, e
		}

		call.service.rateLimitService.Set(endpoint, response)

		var errors []models.Error

		if tweetsResponse.Errors != nil {
			for _, tweetError := range *tweetsResponse.Errors {
				if tweetError.Title == "Not Found Error" {
					nonExistingTweetIDs = append(nonExistingTweetIDs, tweetError.Value)
				} else {
					errors = append(errors, tweetError)
				}
			}

			if len(errors) > 0 {
				e := new(errortools.Error)
				e.SetRequest(request)
				e.SetResponse(response)

				b, err := json.Marshal(errors)
				if err == nil {
					e.SetExtra("errors", string(b))
				}

				e.SetMessage(fmt.Sprintf("%v errors found", len(errors)))
				errortools.CaptureError(e)
			}
		}

		if tweetsResponse.Data != nil {
			tweets = append(tweets, (*tweetsResponse.Data)...)
		}

		if tweetsResponse.Includes != nil {
			if includes.Tweets == nil {
				includes.Tweets = tweetsResponse.Includes.Tweets
			} else if tweetsResponse.Includes.Tweets != nil {
				(*includes.Tweets) = append(*includes.Tweets, (*tweetsResponse.Includes.Tweets)...)
			}

			if includes.Users == nil {
				includes.Users = tweetsResponse.Includes.Users
			} else if tweetsResponse.Includes.Users != nil {
				(*includes.Users) = append(*includes.Users, (*tweetsResponse.Includes.Users)...)
			}

			if includes.Places == nil {
				includes.Places = tweetsResponse.Includes.Places
			} else if tweetsResponse.Includes.Places != nil {
				(*includes.Places) = append(*includes.Places, (*tweetsResponse.Includes.Places)...)
			}

			if includes.Polls == nil {
				includes.Polls = tweetsResponse.Includes.Polls
			} else if tweetsResponse.Includes.Polls != nil {
				(*includes.Polls) = append(*includes.Polls, (*tweetsResponse.Includes.Polls)...)
			}

			if includes.Media == nil {
				includes.Media = tweetsResponse.Includes.Media
			} else if tweetsResponse.Includes.Media != nil {
				(*includes.Media) = append(*includes.Media, (*tweetsResponse.Includes.Media)...)
			}
		}

		if len(ids) <= maximumNumberOfTweetIDsPerCall {
			break
		}

		ids = ids[maximumNumberOfTweetIDsPerCall:]
	}

	return &tweets, &includes, &nonExistingTweetIDs, nil
}
