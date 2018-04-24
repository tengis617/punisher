package coinone

import (
	"github.com/parnurzeal/gorequest"
)

// Orderbook contains information of the current orders at the exchange
type Orderbook struct {
	Result    string  `json:"result"`
	ErrorCode string  `json:"errorCode"`
	Ask       []Order `json:"ask"`
	Bid       []Order `json:"bid"`
	Timestamp int     `json:"timestamp"`
	Currency  string  `json:"currency"`
}

// Order is a bid/ask order on the exchange
type Order struct {
	Price int     `json:"price"`
	Qty   float64 `json:"qty"`
}

const (
	coinoneURL = "https://api.coinone.co.kr"
)

var (
	client *gorequest.SuperAgent
)

func init() {
	client = gorequest.New()
}

// GetOrderbook gets the orderbook from coinone
func GetOrderbook(currency string) *Orderbook {
	orderbook := &Orderbook{}
	url := coinoneURL + "/orderbook/?currency=" + currency
	// Im not sure i agree with gorequest on returning an array of errors as it gets hard to deal with
	// would like to move to a different library ...
	client.Get(url).EndStruct(orderbook)
	return orderbook
}
