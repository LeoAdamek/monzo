package monzo

import "time"

// Merchant represents an invidivual payment merchant.
// Merchants may belong to a group, which will contain other merchants which are similar and can be grouped
// i.e. Different store loctions for a chain would belong to a group
type Merchant struct {
	CreatedAt time.Time `json:"created"`
	GroupID   string    `json:"group_id"`
	ID        string    `json:"id"`

	Address struct {
		Address   string  `json:"address"`
		City      string  `json:"city"`
		Country   string  `json:"country"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
		Postcode  string  `json:"postcode"`
		Region    string  `json:"region"`
	}

	Logo     string `json:"logo"`
	Emoji    string `json:"emoji"`
	Name     string `json:"name"`
	Category string `json:"category"`
}
