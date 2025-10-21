package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	min, max, values := getInput()
	output, err := rangeFilter(min, max, values)
	if err != nil {
		log.Fatal(err)
	}

	// Output
	fmt.Println("Values in range -->", output)
}

func getInput() (float64, float64, []float64) {
	// Variable declaration
	var min, max float64
	var values []float64

	// Input for min & max
	fmt.Println("Enter the values (minimum, maximum, list of values), separated by spaces:")
	fmt.Scan(&min, &max)

	// Input for values
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()

	parts := strings.Fields(line)

	for _, str := range parts {
		num, err := strconv.ParseFloat(str, 64)
		if err != nil {
			fmt.Println("Conversion error")
			continue
		}
		values = append(values, num)
	}
	return min, max, values
}

func rangeFilter(min float64, max float64, values []float64) ([]float64, error) {
	if min > max {
		return nil, fmt.Errorf("invalid limits")
	}

	result := []float64{}

	for _, v := range values {
		if v >= min && v <= max {
			result = append(result, v)
		}
	}

	return result, nil
}
