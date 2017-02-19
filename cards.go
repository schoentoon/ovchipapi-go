package ovchipapi

import (
	"net/url"
)

type OVCard struct {
	Alias              string `json:"alias"`
	MediumID           string `json:"mediumId"`
	Balance            int `json:"balance"`
	BalanceDate        OVTime `json:"balanceDate"`
	DefaultCard        bool `json:"defaultCard"`
	Status             string `json:"status"`
	ExpiryDate         OVTime `json:"expiryDate"`
	AutoReloadEnabled  bool `json:"autoReloadEnabled"`
	Type               string `json:"type"`
	StatusAnnouncement string `json:"statusAnnouncement"`
}

// Get all OV chipkaarten registered to the account
func Cards(authorizationToken string, locale Locale) ([]*OVCard, error) {
	cards := make([]*OVCard, 0)

	err := postAndResponse(cardsUrl, url.Values{
		"authorizationToken": {authorizationToken},
		"locale":             {string(locale)},
	}, &cards)

	return cards, err
}
