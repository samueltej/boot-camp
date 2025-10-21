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

	if err != nil {
		t.Errorf("TestAllValues failed: unexpected error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("TestAllValues failed: values do not match (expected %v, got %v)", want, got)
	}
}

func TestSingleValue(t *testing.T) {
	min, max := 10.0, 20.0
	values := []float64{15}

	want := []float64{15}
	got, err := rangeFilter(min, max, values)

	if err != nil {
		t.Errorf("TestSingleValue failed: unexpected error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("TestSingleValue failed: values do not match (expected %v, got %v)", want, got)
	}
}

func TestGreaterMax(t *testing.T) {
	min, max := 70.0, 30.0
	values := []float64{10, 30, 40, 50, 70, 80}

	got, err := rangeFilter(min, max, values)

	if err == nil {
		t.Errorf("TestGreaterMax failed: expected an error because min (%v) > max (%v)", min, max)
	}

	if got != nil {
		t.Errorf("TestGreaterMax failed: expected nil values, but got %v", got)
	}
}

func TestNoValuesInRange(t *testing.T) {
	min, max := 10.0, 20.0
	values := []float64{30, 40, 50}

	want := []float64{}
	got, err := rangeFilter(min, max, values)

	if err != nil {
		t.Errorf("TestNoValuesInRange failed: unexpected error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("TestNoValuesInRange failed: values do not match (expected %v, got %v)", want, got)
	}
}

func TestNegativeLimit(t *testing.T) {
	min, max := -20.0, -10.0
	values := []float64{-30, -25, -15, -10, 0}

	want := []float64{-15, -10}
	got, err := rangeFilter(min, max, values)

	if err != nil {
		t.Errorf("TestNegativeLimit failed: unexpected error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("TestNegativeLimit failed: values do not match (expected %v, got %v)", want, got)
	}
}

func TestEqualMinAndMax(t *testing.T) {
	min, max := 50.0, 50.0
	values := []float64{10, 50, 70}

	want := []float64{50}
	got, err := rangeFilter(min, max, values)

	if err != nil {
		t.Errorf("TestEqualMinAndMax failed: unexpected error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("TestEqualMinAndMax failed: values do not match (expected %v, got %v)", want, got)
	}
}
