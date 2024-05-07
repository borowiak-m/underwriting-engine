package data

import "time"

type Currency int

func (c Currency) String() string {
	switch c {
	case EUR:
		return "EUR"
	case USD:
		return "USD"
	case GBP:
		return "GBP"
	default:
		return "INVALID"
	}
}

const (
	EUR Currency = iota
	USD
	GBP
	CAD
)

type Subscription struct {
	ExternalID  string    `json:"externalId"`
	StartedAt   time.Time `json:"startedAt"`
	CancelledAt time.Time `json:"cancelledAt"`
	Amount      int       `json:"amount"`
	Currency    Currency  `json:"currency"`
	Period      int       `json:"period"`
	VAT         int       `json:"vat"`
}
