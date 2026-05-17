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

func TestCalculateTotalEmpty(t *testing.T) {
	orders := []float64{}
	expected := 0.0

	result := CalculateTotal(orders)

	if result != expected {
		t.Errorf("got %v, want %v", result, expected)
	}
}

func TestCalculateTotalSingle(t *testing.T) {
	orders := []float64{50.0}
	expected := 50.0

	result := CalculateTotal(orders)

	if result != expected {
		t.Errorf("got %v, want %v", result, expected)
	}
}