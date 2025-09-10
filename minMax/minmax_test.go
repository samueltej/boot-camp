package main

import (
	"reflect"
	"testing"
)

func TestAllValues(t *testing.T) {
	min, max := 30.0, 70.0
	values := []float64{10, 30, 40, 50, 60, 80}

	want := []float64{30, 40, 50, 60}
	got, err := rangeFilter(min, max, values)

	if !reflect.DeepEqual(got, want) && err != nil {
		t.Errorf("TestAllValues failed: values do not match (expected %v, got %v) or an unexpected error occurred: %v", want, got, err)
	}
}

func TestSingleValue(t *testing.T) {
	min, max := 10.0, 20.0
	values := []float64{15}

	want := []float64{15}
	got, err := rangeFilter(min, max, values)

	if !reflect.DeepEqual(got, want) && err != nil {
		t.Errorf("TestSingleValue failed: values do not match (expected %v, got %v) or an unexpected error occurred: %v", want, got, err)
	}
}

func TestGreaterMax(t *testing.T) {
	min, max := 70.0, 30.0
	values := []float64{10, 30, 40, 50, 70, 80}

	got, err := rangeFilter(min, max, values)

	if got == nil && err == nil {
		t.Errorf("TestGreaterMax failed: values were accepted as correct even though min (%v) > max (%v)", min, max)
	}

}

func TestNoValuesInRange(t *testing.T) {
	min, max := 10.0, 20.0
	values := []float64{30, 40, 50}
	want := []float64{}
	got, err := rangeFilter(min, max, values)

	if !reflect.DeepEqual(got, want) && err != nil {
		t.Errorf("TestNoValuesInRange failed: values do not match (expected %v, got %v) or an unexpected error occurred: %v", want, got, err)
	}

}

func TestNegativeLimit(t *testing.T) {
	min, max := -20.0, -10.0
	values := []float64{-30, -25, -15, -10, 0}

	want := []float64{-15, -10}
	got, err := rangeFilter(min, max, values)

	if !reflect.DeepEqual(got, want) && err != nil {
		t.Errorf("TestNegativeLimite failed: values do not match (expected %v, got %v) or an unexpected error occurred: %v", want, got, err)
	}
}

func TestEqualMinAndMax(t *testing.T) {
	min, max := 50.0, 50.0
	values := []float64{10, 50, 70}

	want := []float64{50}
	got, err := rangeFilter(min, max, values)

	if !reflect.DeepEqual(got, want) && err != nil {
		t.Errorf("TestEqualMinAndMax failed: values do not match (expected %v, got %v) or an unexpected error occurred: %v", want, got, err)
	}
}
