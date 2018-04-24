package arbitrage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateArbitrage(t *testing.T) {
	assert := assert.New(t)
	// if someone sells btc for 10$
	// and someone will buy for 15$
	// you can potentially make 5$ per 10$
	// your profit will be 50%
	assert.Equal(0.5, CalculateArbitrage(15, 10), "should return 0.5 when bid=15 ask=10")
	assert.Equal(0.0, CalculateArbitrage(10, 10), "should return 0.0 when bid=10 ask=10")
	assert.Equal(0.01, CalculateArbitrage(101, 100), "should return 0.01 when bid=101 ask=100")
	assert.Equal(0.1, CalculateArbitrage(11, 10), "should return 0.1 when bid=11 ask=10")
	assert.Equal(1.0, CalculateArbitrage(20, 10), "should return 10 when bid=20 ask=10")
	// negative arbitrage cases
	assert.Equal(-0.9, CalculateArbitrage(10, 100), "should return 0.1 when bid=10 ask=100")
	assert.Equal(-1.0, CalculateArbitrage(0, 100), "should return 0.1 when bid=0 ask=100")
}
