package models

import (
	"time"
)

type Tweet struct {
	ID                 string                 `json:"id"`
	Text               string                 `json:"text"`
	CreatedAt          string                 `json:"created_at"`
	AuthorID           string                 `json:"author_id"`
	ConversationID     string                 `json:"conversation_id"`
	InReplyToUserID    *string                `json:"in_reply_to_user_id"`
	ReferencedTweets   *[]ReferencedTweet     `json:"referenced_tweets"`
	Attachments        *Attachments           `json:"attachments"`
	Geo                *Geo                   `json:"geo"`
	ContextAnnotations *[]ContextAnnotation   `json:"context_annotations"`
	Entities           *[]Entity              `json:"entities"`
	Withheld           *Withheld              `json:"withheld"`
	PublicMetrics      *TweetPublicMetrics    `json:"public_metrics"`
	NonPublicMetrics   *TweetNonPublicMetrics `json:"non_public_metrics"`
	OrganicMetrics     *TweetOrganicMetrics   `json:"organic_metrics"`
	PromotedMetrics    *TweetPromotedMetrics  `json:"promoted_metrics"`
	PossiblySensitive  *bool                  `json:"possibly_sensitive"`
	Language           *string                `json:"lang"`
	ReplySettings      *string                `json:"reply_settings"`
	Source             *string                `json:"source"`
}

func (tweet Tweet) CreatedAtTime() (*time.Time, error) {
	return parseTime(tweet.CreatedAt)
}

type ReferencedTweet struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

type Attachments struct {
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

type TweetPublicMetrics struct {
	RetweetCount int `json:"retweet_count"`
	ReplyCount   int `json:"reply_count"`
	LikeCount    int `json:"like_count"`
	QuoteCount   int `json:"quote_count"`
}

type TweetNonPublicMetrics struct {
	ImpressionCount   int `json:"impression_count"`
	URLLinkClicks     int `json:"url_link_clicks"`
	UserProfileClicks int `json:"user_profile_clicks"`
}

type TweetOrganicMetrics struct {
	ImpressionCount   int `json:"impression_count"`
	URLLinkClicks     int `json:"url_link_clicks"`
	UserProfileClicks int `json:"user_profile_clicks"`
	RetweetCount      int `json:"retweet_count"`
	ReplyCount        int `json:"reply_count"`
	LikeCount         int `json:"like_count"`
}

type TweetPromotedMetrics struct {
	ImpressionCount   int `json:"impression_count"`
	URLLinkClicks     int `json:"url_link_clicks"`
	UserProfileClicks int `json:"user_profile_clicks"`
	RetweetCount      int `json:"retweet_count"`
	ReplyCount        int `json:"reply_count"`
	LikeCount         int `json:"like_count"`
}
