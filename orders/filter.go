package orders

func FilterByMinAmount(orders []float64, min float64) []float64 {
	result := []float64{}
	for _, order := range orders {
		if order >= min {
			result = append(result, order)
		}
	}
	return result
}