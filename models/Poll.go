package models

type Poll struct {
	ID              string        `json:"id"`
	Options         *[]PollOption `json:"options"`
	DurationMinutes int64         `json:"duration_minutes"`
	EndDatetime     string        `json:"end_datetime"`
	VotingStatus    string        `json:"voting_status"`
}

type PollOption struct {
	Position int64  `json:"position"`
	Label    string `json:"label"`
	Votes    int64  `json:"votes"`
}
