package orders

import (
	"reflect"
	"testing"
)

func TestFilterByMinAmount(t *testing.T) {
	orders := []float64{50.0, 200.0, 30.0, 150.0}
	result := FilterByMinAmount(orders, 100.0)
	expected := []float64{200.0, 150.0}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("got %v, want %v", result, expected)
	}
}

func TestFilterByMinAmountEmpty(t *testing.T) {
	result := FilterByMinAmount([]float64{}, 100.0)
	expected := []float64{}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("got %v, want %v", result, expected)
	}
}

func TestFilterByMinAmountAllFiltered(t *testing.T) {
	result := FilterByMinAmount([]float64{10.0, 20.0, 30.0}, 100.0)
	expected := []float64{}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("got %v, want %v", result, expected)
	}
}