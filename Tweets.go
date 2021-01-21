package twitter

import (
	"fmt"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	oauth2 "github.com/leapforce-libraries/go_oauth2"
)

type TweetsResponse struct {
	Data     *[]Tweet   `json:"data"`
	Includes *[]Include `json:"includes"`
	Meta     *Meta      `json:"meta"`
}

type Tweet struct {
	ID                 string               `json:"id"`
	Text               string               `json:"text"`
	CreatedAt          string               `json:"created_at"`
	AuthorID           string               `json:"author_id"`
	ConversationID     string               `json:"conversation_id"`
	InReplyToUserID    string               `json:"in_reply_to_user_id"`
	ReferencedTweets   *[]ReferencedTweet   `json:"referenced_tweets"`
	Attachments        *[]Attachment        `json:"attachments"`
	Geo                *Geo                 `json:"geo"`
	ContextAnnotations *[]ContextAnnotation `json:"context_annotations"`
	Entities           *[]Entity            `json:"entities"`
	Withheld           *Withheld            `json:"withheld"`
	PublicMetrics      *PublicMetrics       `json:"public_metrics"`
	NonPublicMetrics   *NonPublicMetrics    `json:"non_public_metrics"`
	OrganicMetrics     *OrganicMetrics      `json:"organic_metrics"`
	PromotedMetrics    *PromotedMetrics     `json:"promoted_metrics"`
	PossiblySensitive  *bool                `json:"possibly_sensitive"`
	Language           *string              `json:"lang"`
	ReplySettings      *string              `json:"reply_settings"`
	Source             *string              `json:"source"`
}

type ReferencedTweet struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

type Attachment struct {
	MediaKeys []string `json:"media_keys"`
	PollIDs   []string `json:"poll_ids"`
}

type Geo struct {
	Coordinates Coordinates `json:"coordinates"`
	PlaceID     string      `json:"place_id"`
}

type Coordinates struct {
	Type        string     `json:"type"`
	Coordinates *[]float64 `json:"coordinates"`
}

type ContextAnnotation struct {
	Domain ContextAnnotationDomain `json:"domain"`
	Entity ContextAnnotationEntity `json:"entity"`
}

type ContextAnnotationDomain struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ContextAnnotationEntity struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Entity struct {
	Annotations []EntityAnnotation `json:"annotations"`
	URLs        []EntityURL        `json:"urls"`
	Hashtags    []EntityHashtag    `json:"hashtags"`
	Mentions    []EntityMention    `json:"mentions"`
	Cashtags    []EntityCashtag    `json:"cashtags"`
}

type EntityAnnotation struct {
	Start          int     `json:"start"`
	End            int     `json:"end"`
	Probability    float64 `json:"probability"`
	Type           string  `json:"type"`
	NormalizedText string  `json:"normalized_text"`
}

type EntityURL struct {
	Start       int    `json:"start"`
	End         int    `json:"end"`
	URL         string `json:"url"`
	ExpandedURL string `json:"expanded_url"`
	DisplayURL  string `json:"display_url"`
	UnwoundURL  string `json:"unwound_url"`
}

type EntityHashtag struct {
	Start int    `json:"start"`
	End   int    `json:"end"`
	Tag   string `json:"tag"`
}

type EntityMention struct {
	Start    int    `json:"start"`
	End      int    `json:"end"`
	Username string `json:"username"`
}

type EntityCashtag struct {
	Start int    `json:"start"`
	End   int    `json:"end"`
	Tag   string `json:"tag"`
}

type Withheld struct {
	Copyright    bool     `json:"copyright"`
	CountryCodes []string `json:"country_codes"`
	Scope        string   `json:"scope"`
}

type PublicMetrics struct {
	RetweetCount int64 `json:"retweet_count"`
	ReplyCount   int64 `json:"reply_count"`
	LikeCount    int64 `json:"like_count"`
	QuoteCount   int64 `json:"quote_count"`
}

type NonPublicMetrics struct {
	ImpressionCount   int64 `json:"impression_count"`
	URLLinkClicks     int64 `json:"url_link_clicks"`
	UserProfileClicks int64 `json:"user_profile_clicks"`
}

type OrganicMetrics struct {
	ImpressionCount   int64 `json:"impression_count"`
	URLLinkClicks     int64 `json:"url_link_clicks"`
	UserProfileClicks int64 `json:"user_profile_clicks"`
	RetweetCount      int64 `json:"retweet_count"`
	ReplyCount        int64 `json:"reply_count"`
	LikeCount         int64 `json:"like_count"`
}

type PromotedMetrics struct {
	ImpressionCount   int64 `json:"impression_count"`
	URLLinkClicks     int64 `json:"url_link_clicks"`
	UserProfileClicks int64 `json:"user_profile_clicks"`
	RetweetCount      int64 `json:"retweet_count"`
	ReplyCount        int64 `json:"reply_count"`
	LikeCount         int64 `json:"like_count"`
}

type Include struct {
	Tweets *[]Tweet `json:"tweets"`
	//Users  *[]User  `json:"users"`
	//Places *[]Place `json:"places"`
	//Media  *[]Media `json:"media"`
	//Polls  *[]Poll  `json:"polls"`
}

type Meta struct {
	ResultCount   int     `json:"result_count"`
	NewestID      *string `json:"newest_id"`
	OldestID      *string `json:"oldest_id"`
	NextToken     *string `json:"next_token"`
	PreviousToken *string `json:"previous_token"`
}

type Exclude string

const (
	ExcludeRetweets Exclude = "retweets"
	ExcludeReplies  Exclude = "replies"
)

type Expansion string

const (
	ExpansionAttachmentsPollIDs         Expansion = "attachments.poll_ids"
	ExpansionAttachmentsMediaKeys       Expansion = "attachments.media_keys"
	ExpansionAuthorID                   Expansion = "author_id"
	ExpansionEntitiesMentionsUsername   Expansion = "entities.mentions.username"
	ExpansionGeoPlaceID                 Expansion = "geo.place_id"
	ExpansionInReplyToUserID            Expansion = "in_reply_to_user_id"
	ExpansionReferencedTweetsID         Expansion = "referenced_tweets.id"
	ExpansionReferencedTweetsIDAuthorID Expansion = "referenced_tweets.id.author_id"
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
	UserFieldreatedAt        UserField = "created_at"
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
	userID          int64
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

func (service *Service) NewGetTweetsCall(userID int64) *GetTweetsCall {
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

func (call *GetTweetsCall) SetExpansions(expansions ...Expansion) *GetTweetsCall {
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
	elems := []string{}

	for _, elem := range mediaFields {
		elems = append(elems, string(elem))
	}
	(*call).MediaFields = &elems

	return call
}

func (call *GetTweetsCall) SetPaginationToken(paginationToken string) *GetTweetsCall {
	(*call).PaginationToken = &paginationToken

	return call
}

func (call *GetTweetsCall) SetPlaceFields(placeFields ...PlaceField) *GetTweetsCall {
	elems := []string{}

	for _, elem := range placeFields {
		elems = append(elems, string(elem))
	}
	(*call).PlaceFields = &elems

	return call
}

func (call *GetTweetsCall) SetPollFields(pollFields ...PollField) *GetTweetsCall {
	elems := []string{}

	for _, elem := range pollFields {
		elems = append(elems, string(elem))
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
	elems := []string{}

	for _, elem := range tweetFields {
		elems = append(elems, string(elem))
	}
	(*call).TweetFields = &elems

	return call
}

func (call *GetTweetsCall) SetUntilID(untilID string) *GetTweetsCall {
	(*call).UntilID = &untilID

	return call
}

func (call *GetTweetsCall) SetUserFields(userFields ...UserField) *GetTweetsCall {
	elems := []string{}

	for _, elem := range userFields {
		elems = append(elems, string(elem))
	}
	(*call).UserFields = &elems

	return call
}

func (call *GetTweetsCall) Do() (*[]Tweet, *errortools.Error) {
	tweets := []Tweet{}

	rowCount := 0

	for true {
		params, e := call.service.urlParams(call)
		if e != nil {
			return nil, e
		}

		urlPath := fmt.Sprintf("users/%v/tweets%s", call.userID, *params)
		fmt.Println(urlPath)

		tweetsResponse := TweetsResponse{}
		requestConfig := oauth2.RequestConfig{
			URL:           call.service.url(urlPath),
			ResponseModel: &tweetsResponse,
		}
		_, _, e = call.service.get(&requestConfig)
		if e != nil {
			return nil, e
		}

		if tweetsResponse.Data == nil {
			break
		}

		rowCountCall := len(*tweetsResponse.Data)

		if rowCountCall > 0 {
			tweets = append(tweets, *tweetsResponse.Data...)
			rowCount += rowCountCall
		}

		if tweetsResponse.Meta == nil {
			break
		}

		if tweetsResponse.Meta.NextToken == nil {
			break
		}

		call.PaginationToken = tweetsResponse.Meta.NextToken
	}

	return &tweets, nil
}
