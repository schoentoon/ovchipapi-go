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

type OVCardDetail struct {
	Card struct {
		Alias                     string `json:"alias"`
		Balance                   int `json:"balance"`
		BalanceDate               OVTime `json:"balanceDate"`
		MediumID                  string `json:"mediumId"`
		ExpiryDate                OVTime `json:"expiryDate"`
		DefaultCard               bool `json:"defaultCard"`
		AutoReloadEnabled         bool `json:"autoReloadEnabled"`
		AutoReloadAccountNumber   string `json:"autoReloadAccountNumber"`
		AutoReloadAmount          int `json:"autoReloadAmount"`
		AutoReloadPaymentMandate  string `json:"autoReloadPaymentMandate"`
		AutoReloadThresholdAmount int `json:"autoReloadThresholdAmount"`
		Status                    string `json:"status"`
		StatusAnnouncement        string `json:"statusAnnouncement"`
		Type                      string `json:"type"`
	} `json:"card"`
	ProductInfoList []struct {
		ProductTitle             string `json:"productTitle"`
		ProductTitleExplanation  string `json:"productTitleExplanation"`
		ProductStatus            string `json:"productStatus"`
		ProductStatusExplanation string `json:"productStatusExplanation"`
		PassengerClass           string `json:"passengerClass"`
		GeographicValidity       string `json:"geographicValidity"`
		ProductValidity          string `json:"productValidity"`
	} `json:"productInfoList"`
}

// Get a single OV Chipkaart
func Card(authorizationToken string, locale Locale, mediumId string) (*OVCardDetail, error) {
	card := &OVCardDetail{}

	err := postAndResponse(cardUrl, url.Values{
		"authorizationToken": {authorizationToken},
		"locale":             {string(locale)},
		"mediumId":           {mediumId},
	}, &card)

	return card, err
}
