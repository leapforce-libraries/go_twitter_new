package twitter

import (
	"encoding/json"
	"fmt"
	"strings"
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
	Count         int    `json:"count"`
	NewestID      string `json:"newest_id"`
	OldestID      string `json:"oldest_id"`
	NextToken     string `json:"next_token"`
	PreviousToken string `json:"previous_token"`
}

type Enums []string

func (enums *Enums) MarshalJSON() ([]byte, error) {
	if enums == nil {
		return nil, nil
	}

	_enums := strings.Join(*enums, ",")

	return json.Marshal(_enums)
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

type Excludes Enums
type Expansions Enums
type MediaFields Enums
type PlaceFields Enums
type PollFields Enums
type TweetFields Enums
type UserFields Enums

type Time struct {
	time time.Time
}

func (t *Time) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}

	return json.Marshal(t.time.Format(time.RFC3339))
}

type GetTweetsParams struct {
	UserID          int64
	EndTime         *Time        `json:"end_time"`
	Exclude         *Excludes    `json:"exclude"`
	Expansions      *Expansions  `json:"expansions"`
	MaxResults      *int         `json:"max_results"`
	MediaFields     *MediaFields `json:"media.fields"`
	PaginationToken *string      `json:"pagination_token"`
	PlaceFields     *PlaceFields `json:"place.fields"`
	PollFields      *PollFields  `json:"poll.fields"`
	SinceID         *string      `json:"since_id"`
	StartTime       *Time        `json:"start_time"`
	TweetFields     *TweetFields `json:"tweet.fields"`
	UntilID         *string      `json:"until_id"`
	UserFields      *UserFields  `json:"user.fields"`
}

func (service *Service) GetTweets(params *GetTweetsParams) (*[]Tweet, *errortools.Error) {
	top := 100
	skip := 0

	tweets := []Tweet{}

	rowCount := 0

	for skip == 0 || rowCount > 0 {
		urlPath := fmt.Sprintf("users/%v/tweets", params.UserID)

		tweetsResponse := TweetsResponse{}
		requestConfig := oauth2.RequestConfig{
			URL:           service.url(urlPath),
			ResponseModel: &tweetsResponse,
		}
		_, _, e := service.get(&requestConfig)
		if e != nil {
			return nil, e
		}

		rowCount = len(*tweetsResponse.Data)

		if rowCount > 0 {
			tweets = append(tweets, *tweetsResponse.Data...)
		}

		skip += top
	}

	return &tweets, nil
}
