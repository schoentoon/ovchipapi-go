package ovchipapi

import (
	"net/url"
	"strconv"
)

type OVTransaction struct {
	CheckInInfo            string `json:"checkInInfo"`
	CheckInText            string `json:"checkInText"`
	Fare                   float64 `json:"fare"`
	FareCalculation        string `json:"fareCalculation"`
	FareText               string `json:"fareText"`
	ModalType              string `json:"modalType"`
	ProductInfo            string `json:"productInfo"`
	ProductText            string `json:"productText"`
	Carrier                string `json:"pto"`
	TransactionDateTime    OVTime `json:"transactionDateTime"`
	TransactionInfo        string `json:"transactionInfo"`
	TransactionName        string `json:"transactionName"`
	EPurseMut              float64 `json:"ePurseMut"`
	EPurseMutInfo          string `json:"ePurseMutInfo"`
	TransactionExplanation string `json:"transactionExplanation"`
	TransactionPriority    string `json:"transactionPriority"`
}

type transactionsResponse struct {
	TotalSize int `json:"totalSize"`
	Records   []*OVTransaction `json:"records"`
	NextRequestContext struct {
		StartDate string `json:"startDate"`
		EndDate   string `json:"endDate"`
		Offset    int `json:"offset"`
	} `json:"nextRequestContext"`
}

// Get all transactions from startDate to endDate. This will potentially fire off a lot of requests.
// It will retrieve all transactions in requests of 20/request because this is a limit in the API.
func Transactions(authorizationToken string, locale Locale, mediumId string, startDate string, endDate string) ([]*OVTransaction, error) {
	firstSlice, err := transactionsSlice(authorizationToken, locale, mediumId, 0, startDate, endDate)
	if err != nil {
		return nil, err
	}

	transactions := make([]*OVTransaction, 0, firstSlice.TotalSize)
	transactions = append(transactions, firstSlice.Records...)

	offset := firstSlice.NextRequestContext.Offset

	for offset < firstSlice.TotalSize {
		slice, err := transactionsSlice(authorizationToken, locale, mediumId, offset, startDate, endDate)
		if err != nil {
			return nil, err
		}

		transactions = append(transactions, slice.Records...)
		offset += 20
	}

	return transactions, nil
}

func transactionsSlice(authorizationToken string, locale Locale, mediumId string, offset int, startDate string, endDate string) (*transactionsResponse, error) {
	object := &transactionsResponse{}

	err := postAndResponse(transactionsUrl, url.Values{
		"authorizationToken": {authorizationToken},
		"locale":             {string(locale)},
		"mediumId":           {mediumId},
		"offset":             {strconv.Itoa(offset)},
		"startDate":          {startDate},
		"endDate":            {endDate},
	}, object)

	return object, err
}
