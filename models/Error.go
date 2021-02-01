package models

// Error stores general Ridder API error response
//
type Error struct {
	Detail       string `json:"detail"`
	Parameter    string `json:"parameter"`
	ResourceType string `json:"resource_type"`
	Title        string `json:"title"`
	Type         string `json:"type"`
	Value        string `json:"value"`
}
