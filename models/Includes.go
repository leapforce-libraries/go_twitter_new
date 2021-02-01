package models

type Includes struct {
	Tweets *[]Tweet `json:"tweets"`
	Users  *[]User  `json:"users"`
	Places *[]Place `json:"places"`
	Media  *[]Media `json:"media"`
	Polls  *[]Poll  `json:"polls"`
}
