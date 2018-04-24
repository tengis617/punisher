package coinone

import (
	"log"

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
	coinoneURL = "https://api.coinone.co.kr/"
)

var (
	client *gorequest.SuperAgent
)

func init() {
	client = gorequest.New()
}

// GetOrderbook gets the orderbook from coinone
func GetOrderbook() (*Orderbook, error) {
	orderbook := &Orderbook{}
	_, _, errs := client.Get("https://api.coinone.co.kr/orderbook/").EndStruct(orderbook)

	if errs != nil {
		log.Printf("could not fetch: %v", errs)
	}
	return orderbook, nil
}
