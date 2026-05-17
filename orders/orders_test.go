package orders

import "testing"

func TestCalculateTotal(t *testing.T) {
	orders := []float64{100.50, 200.00, 50.75}
	expected := 351.25

	result := CalculateTotal(orders)

	if result != expected {
		t.Errorf("got %v, want %v", result, expected)
	}
}