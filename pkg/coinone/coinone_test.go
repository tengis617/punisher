package coinone

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestGetOrderbook(t *testing.T) {
	assert := assert.New(t)

	t.Run("should return the orderbook", func(t *testing.T) {
		orderbook := GetOrderbook("btc")
		assert.NotNil(orderbook)
		assert.Equal("success", orderbook.Result)
	})
	t.Run("should return xrp orderbook", func(t *testing.T) {
		orderbook := GetOrderbook("xrp")
		assert.Equal("xrp", orderbook.Currency)
		assert.Equal("success", orderbook.Result)
	})
}
