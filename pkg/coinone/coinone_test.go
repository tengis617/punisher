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

	t.Run("should return no error", func(t *testing.T) {
		orderbook, err := GetOrderbook()
		assert.NotNil(orderbook)
		assert.Nil(err)
		assert.NotNil(orderbook)
	})
}
