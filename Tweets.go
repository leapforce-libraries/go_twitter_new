package twitter

import (
	"encoding/json"
	"fmt"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	models "github.com/leapforce-libraries/go_twitter_new/models"
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
	MediaFieldPreviewImageURL  MediaField = "preview_image_url"
	MediaFieldType             MediaField = "type"
	MediaFieldURL              MediaField = "url"
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
	UserFieldProfileImageURL UserField = "profile_image_url"
	UserFieldProtected       UserField = "protected"
	UserFieldPublicMetrics   UserField = "public_metrics"
	UserFieldURL             UserField = "url"
	UserFieldUsername        UserField = "username"
	UserFieldVerified        UserField = "verified"
	UserFieldWithheld        UserField = "withheld"
)

type GetTweetsCall struct {
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

func (service *Service) NewGetTweetsCall(userID string) *GetTweetsCall {
	return &GetTweetsCall{
		service: service,
		userID:  userID,
	}
}

func (call *GetTweetsCall) SetEndTime(endTime time.Time) *GetTweetsCall {
	(*call).EndTime = &endTime

	return call
}

func (call *GetTweetsCall) SetExclude(excludes ...Exclude) *GetTweetsCall {
	elems := []string{}

	for _, elem := range excludes {
		elems = append(elems, string(elem))
	}
	(*call).Exclude = &elems

	return call
}

func (call *GetTweetsCall) SetExpansions(expansions ...TweetExpansion) *GetTweetsCall {
	elems := []string{}

	for _, elem := range expansions {
		elems = append(elems, string(elem))
	}
	(*call).Expansions = &elems

	return call
}

func (call *GetTweetsCall) SetMaxResults(maxResults int) *GetTweetsCall {
	(*call).MaxResults = &maxResults

	return call
}

func (call *GetTweetsCall) SetMediaFields(mediaFields ...MediaField) *GetTweetsCall {
	return call.setMediaFields(false, mediaFields)
}

func (call *GetTweetsCall) AddMediaFields(mediaFields ...MediaField) *GetTweetsCall {
	return call.setMediaFields(true, mediaFields)
}

func (call *GetTweetsCall) setMediaFields(add bool, mediaFields []MediaField) *GetTweetsCall {
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
}

func (call *GetTweetsCall) SetPaginationToken(paginationToken string) *GetTweetsCall {
	(*call).PaginationToken = &paginationToken

	return call
}

func (call *GetTweetsCall) SetPlaceFields(placeFields ...PlaceField) *GetTweetsCall {
	return call.setPlaceFields(false, placeFields)
}

func (call *GetTweetsCall) AddPlaceFields(placeFields ...PlaceField) *GetTweetsCall {
	return call.setPlaceFields(true, placeFields)
}

func (call *GetTweetsCall) setPlaceFields(add bool, placeFields []PlaceField) *GetTweetsCall {
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
}

func (call *GetTweetsCall) SetPollFields(pollFields ...PollField) *GetTweetsCall {
	return call.setPollFields(false, pollFields)
}

func (call *GetTweetsCall) AddPollFields(pollFields ...PollField) *GetTweetsCall {
	return call.setPollFields(true, pollFields)
}

func (call *GetTweetsCall) setPollFields(add bool, pollFields []PollField) *GetTweetsCall {
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
}

func (call *GetTweetsCall) SetSinceID(sinceID string) *GetTweetsCall {
	(*call).SinceID = &sinceID

	return call
}

func (call *GetTweetsCall) SetStartTime(startTime time.Time) *GetTweetsCall {
	(*call).StartTime = &startTime

	return call
}

func (call *GetTweetsCall) SetTweetFields(tweetFields ...TweetField) *GetTweetsCall {
	return call.setTweetFields(false, tweetFields)
}

func (call *GetTweetsCall) AddTweetFields(tweetFields ...TweetField) *GetTweetsCall {
	return call.setTweetFields(true, tweetFields)
}

func (call *GetTweetsCall) setTweetFields(add bool, tweetFields []TweetField) *GetTweetsCall {
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

func (call *GetTweetsCall) SetUntilID(untilID string) *GetTweetsCall {
	(*call).UntilID = &untilID

	return call
}

func (call *GetTweetsCall) SetUserFields(userFields ...UserField) *GetTweetsCall {
	return call.setUserFields(false, userFields)
}

func (call *GetTweetsCall) AddUserFields(userFields ...UserField) *GetTweetsCall {
	return call.setUserFields(true, userFields)
}

func (call *GetTweetsCall) setUserFields(add bool, userFields []UserField) *GetTweetsCall {
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

func (call *GetTweetsCall) Do() (*[]models.Tweet, *models.Includes, *errortools.Error) {
	tweets := []models.Tweet{}
	includes := models.Includes{
		Tweets: &[]models.Tweet{},
		Users:  &[]models.User{},
		Places: &[]models.Place{},
		Media:  &[]models.Media{},
		Polls:  &[]models.Poll{},
	}

	rowCount := 0

	for true {
		params, e := call.service.urlParams(call)
		if e != nil {
			return nil, nil, e
		}

		urlPath := fmt.Sprintf("users/%s/tweets%s", call.userID, *params)
		//fmt.Println(call.service.url(urlPath))

		tweetsResponse := TweetsResponse{}
		requestConfig := go_http.RequestConfig{
			URL:           call.service.url(urlPath),
			ResponseModel: &tweetsResponse,
		}

		endpoint := "users_tweets"
		call.service.rateLimitService.Check(endpoint)

		request, response, e := call.service.get(&requestConfig)
		if e != nil {
			return nil, nil, e
		}

		call.service.rateLimitService.Set(endpoint, response)

		if tweetsResponse.Errors != nil {
			e := new(errortools.Error)
			e.SetRequest(request)
			e.SetResponse(response)

			b, err := json.Marshal(tweetsResponse.Errors)
			if err == nil {
				e.SetExtra("errors", string(b))
			}

			return nil, nil, errortools.ErrorMessage(fmt.Sprintf("%v errors found", len(*tweetsResponse.Errors)))
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

	return &tweets, &includes, nil
}
