package models

type Place struct {
	FullName        string    `json:"full_name"`
	ID              string    `json:"id"`
	ContainedWithin *[]string `json:"contained_within"`
	Country         string    `json:"country"`
	CountryCode     string    `json:"country_code"`
	//Geo             *Geo      `json:"geo"`
	Name      string `json:"name"`
	PlaceType string `json:"place_type"`
}
