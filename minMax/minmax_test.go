package main

import (
	"reflect"
	"testing"
)

func TestAllValues(t *testing.T) {
	min, max := 30.0, 70.0
	values := []float64{10, 30, 40, 50, 60, 80}

	want := []float64{30, 40, 50, 60}
	got, _ := minMax(min, max, values)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("falló TestMinMaxWithValuesInsideRange: esperaba %v, obtuve %v", want, got)
	}
}

func TestSingleValue(t *testing.T) {
	min, max := 10.0, 20.0
	values := []float64{15}

	want := []float64{15}
	got, _ := minMax(min, max, values)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("falló TestMinMaxWithSingleValue: esperaba %v, obtuve %v", want, got)
	}
}

func TestGreaterMax(t *testing.T) {
	min, max := 70.0, 30.0
	values := []float64{10, 30, 40, 50, 70, 80}

	_, err := minMax(min, max, values)
	if err == nil {
		t.Errorf("falló TestGreaterMax, los limites min %f, y %f son correctos", min, max)
	}

}

func TestNoValuesInRange(t *testing.T) {
	min, max := 10.0, 20.0
	values := []float64{30, 40, 50}

	_, err := minMax(min, max, values)

	if err == nil {
		t.Errorf("falló TestNoValuesInRange, los valores %v, estan dentro de los limites", values)
	}
}

func TestNegativeLimit(t *testing.T) {
	min, max := -20.0, -10.0
	values := []float64{-30, -25, -15, -10, 0}

	want := []float64{-15, -10}
	got, _ := minMax(min, max, values)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("falló TestMinMaxWithNegativeRange: esperaba %v, obtuve %v", want, got)
	}
}


func TestMinMaxWithEqualMinAndMax(t *testing.T) {
	min, max := 50.0, 50.0
	values := []float64{10, 50, 70}

	want := []float64{50}
	got, _ := minMax(min, max, values)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("falló TestMinMaxWithEqualMinAndMax: esperaba %v, obtuve %v", want, got)
	}
}
