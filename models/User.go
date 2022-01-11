package models

type User struct {
	ID              string             `json:"id"`
	Name            string             `json:"name"`
	Username        string             `json:"username"`
	CreatedAt       string             `json:"created_at"`
	Description     string             `json:"description"`
	Entities        *Entities          `json:"entities"`
	Location        *string            `json:"location"`
	PinnedTweetID   *string            `json:"pinned_tweet_id"`
	ProfileImageURL *string            `json:"profile_image_url"`
	Protected       bool               `json:"protected"`
	PublicMetrics   *UserPublicMetrics `json:"public_metrics"`
	URL             *string            `json:"url"`
	Verified        bool               `json:"verified"`
	Withheld        *Withheld          `json:"withheld"`
}

type UserPublicMetrics struct {
	FollowersCount int `json:"followers_count"`
	FollowingCount int `json:"following_count"`
	TweetCount     int `json:"tweet_count"`
	ListedCount    int `json:"listed_count"`
}
