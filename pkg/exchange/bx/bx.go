package bx

import (
	"github.com/parnurzeal/gorequest"
)

// Orderbook contains information of the current orders at the exchange
type Orderbook struct {
	Asks [][2]string `json:"asks"`
	Bids [][2]string `json:"bids"`
}

const (
	exchangeURL = "https://bx.in.th/api"
)

var (
	client *gorequest.SuperAgent
)

func init() {
	client = gorequest.New()
}

// GetOrderbook gets the orderbook from coinone
func GetOrderbook(pairingID string) *Orderbook {
	orderbook := &Orderbook{}
	url := exchangeURL + "/orderbook/?pairing=" + pairingID
	// Im not sure i agree with gorequest on returning an array of errors as it gets hard to deal with
	// would like to move to a different library ...
	client.Get(url).EndStruct(orderbook)
	return orderbook
}
