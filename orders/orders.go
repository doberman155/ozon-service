package orders

func CalculateTotal(orders []float64) float64 {
	total := 0.0
	for _, order := range orders {
		total += order
	}
	return total
}