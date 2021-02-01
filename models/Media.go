package models

type Media struct {
	MediaKey         string                 `json:"media_key"`
	Type             string                 `json:"type"`
	DurationMS       int64                  `json:"duration_ms"`
	Height           int64                  `json:"height"`
	NonPublicMetrics *MediaNonPublicMetrics `json:"non_public_metrics"`
	OrganicMetrics   *MediaOrganicMetrics   `json:"organic_metrics"`
	PreviewImageURL  string                 `json:"preview_image_url"`
	PromotedMetrics  *MediaPromotedMetrics  `json:"promoted_metrics"`
	PublicMetrics    *MediaPublicMetrics    `json:"public_metrics"`
	Width            int64                  `json:"width"`
}

type MediaNonPublicMetrics struct {
	Playback0Count   int `json:"playback_0_count"`
	Playback25Count  int `json:"playback_25_count"`
	Playback50Count  int `json:"playback_50_count"`
	Playback75Count  int `json:"playback_75_count"`
	Playback100Count int `json:"playback_100_count"`
}

type MediaOrganicMetrics struct {
	Playback0Count   int `json:"playback_0_count"`
	Playback25Count  int `json:"playback_25_count"`
	Playback50Count  int `json:"playback_50_count"`
	Playback75Count  int `json:"playback_75_count"`
	Playback100Count int `json:"playback_100_count"`
	ViewCount        int `json:"view_count"`
}

type MediaPromotedMetrics struct {
	Playback0Count   int `json:"playback_0_count"`
	Playback25Count  int `json:"playback_25_count"`
	Playback50Count  int `json:"playback_50_count"`
	Playback75Count  int `json:"playback_75_count"`
	Playback100Count int `json:"playback_100_count"`
	ViewCount        int `json:"view_count"`
}

type MediaPublicMetrics struct {
	ViewCount int `json:"view_count"`
}
