package arbitrage

// CalculateArbitrage will calculate arbitrage with given bid and ask price.
// if bid is higher than ask => arbitrage!
func CalculateArbitrage(bid, ask int) float64 {
	delta := calculateDelta(bid, ask)
	return delta
}

func calculateDelta(a, b int) float64 {
	diff := float64(a - b)
	return (diff / float64(b))
}
