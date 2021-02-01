package models

type Meta struct {
	ResultCount   int     `json:"result_count"`
	NewestID      *string `json:"newest_id"`
	OldestID      *string `json:"oldest_id"`
	NextToken     *string `json:"next_token"`
	PreviousToken *string `json:"previous_token"`
}
