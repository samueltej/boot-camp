package main

import (
	"reflect"
	"testing"
)

func TestAllValues(t *testing.T) {
	min, max := 30.0, 70.0
	values := []float64{10, 30, 40, 50, 60, 80}

	want := []float64{30, 40, 50, 60}
	got, _ := rangeFilter(min, max, values)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("failed TestAllValues: expected %v, got %v", want, got)
	}
}

func TestSingleValue(t *testing.T) {
	min, max := 10.0, 20.0
	values := []float64{15}

	want := []float64{15}
	got, _ := rangeFilter(min, max, values)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("failed TestSingleValue: expected %v, got %v", want, got)
	}
}

func TestGreaterMax(t *testing.T) {
	min, max := 70.0, 30.0
	values := []float64{10, 30, 40, 50, 70, 80}

	_, err := rangeFilter(min, max, values)

	if err == nil || err.Error() != "invalid limits" {
		t.Errorf("failed TestGreaterMax: expected 'invalid limits' error, got %v", err)
	}

}

func TestNoValuesInRange(t *testing.T) {
	min, max := 10.0, 20.0
	values := []float64{30, 40, 50}

	_, err := rangeFilter(min, max, values)

	if err == nil || err.Error() != "no values found in range" {
		t.Errorf("failed TestNoValuesInRange: expected 'no values found in range', got %v", err)
	}

}

func TestNegativeLimit(t *testing.T) {
	min, max := -20.0, -10.0
	values := []float64{-30, -25, -15, -10, 0}

	want := []float64{-15, -10}
	got, _ := rangeFilter(min, max, values)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("failed TestNegativeLimit: expected %v, got %v", want, got)
	}
}

func TestEqualMinAndMax(t *testing.T) {
	min, max := 50.0, 50.0
	values := []float64{10, 50, 70}

	want := []float64{50}
	got, _ := rangeFilter(min, max, values)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("failed TestEqualMinAndMax: expected %v, got %v", want, got)
	}
}
