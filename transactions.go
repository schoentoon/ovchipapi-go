package ovchipapi

import (
	"math"
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

type resultError struct {
	Result      *transactionsResponse
	BatchOffset int
	Error       error
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

	offset := 20

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

// Get all transactions from startDate to endDate. This will potentially fire off a lot of requests.
// It will retrieve all transactions in requests of 20/request because this is a limit in the API.
// This method will execute all requests in parallel and is a lot faster compared to Transactions.
func TransactionsAsync(authorizationToken string, locale Locale, mediumId string, startDate string, endDate string) ([]*OVTransaction, error) {
	firstSlice, err := transactionsSlice(authorizationToken, locale, mediumId, 0, startDate, endDate)
	if err != nil {
		return nil, err
	}

	transactions := make([]*OVTransaction, firstSlice.TotalSize)

	copy(transactions, firstSlice.Records)

	totalBatches := int(math.Ceil(float64(firstSlice.TotalSize) / 20))

	results := make(chan resultError)

	for i := 1; i < totalBatches; i++ {
		batchOffset := i * 20

		go func() {
			slice, err := transactionsSlice(authorizationToken, locale, mediumId, batchOffset, startDate, endDate)
			if err != nil {
				results <- resultError{nil, batchOffset, err}
				return
			}

			results <- resultError{slice, batchOffset, nil}
		}()
	}

	var resultError error = nil

	for i := 1; i < totalBatches; i++ {
		result := <-results

		if resultError != nil {
			continue
		}

		if result.Error != nil { // we still need to receive all responses, so keep looping
			resultError = result.Error
			continue
		}

		for j, v := range result.Result.Records {
			transactions[result.BatchOffset+j] = v
		}
	}

	if resultError != nil {
		return nil, resultError
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
