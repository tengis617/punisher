package exchange

import (
	"fmt"
	"strconv"
	"time"

	"github.com/tengis617/punisher/pkg/exchange/bx"

	"github.com/tengis617/punisher/pkg/exchange/coinone"
)

// Watch watches for prices on given exchanges
func Watch() {
	c := time.Tick(2 * time.Second)
	for range c {
		ticker := "xrp"
		orderbookCoinone := coinone.GetOrderbook(ticker)
		AskOrderCoinoneKR, _ := strconv.ParseFloat(orderbookCoinone.Ask[0].Price, 64)
		BidOrderCoinoneKR, _ := strconv.ParseFloat(orderbookCoinone.Bid[0].Price, 64)
		AskOrderCoinoneTHB := AskOrderCoinoneKR / 34.94
		BidOrderCoinoneTHB := BidOrderCoinoneKR / 34.94

		orderbookBX := bx.GetOrderbook("25")
		AskOrderBX, _ := strconv.ParseFloat(orderbookBX.Asks[0][0], 64)
		BidOrderBX, _ := strconv.ParseFloat(orderbookBX.Bids[0][0], 64)
		// 34.94

		fmt.Printf("KR - ASK: %v - BID: %v - %s \n", AskOrderCoinoneTHB, BidOrderCoinoneTHB, ticker)
		fmt.Printf("TH - ASK: %v - BID: %v - %s \n", AskOrderBX, BidOrderBX, ticker)
	}
}
