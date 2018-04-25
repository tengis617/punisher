package bx

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
	})
}
