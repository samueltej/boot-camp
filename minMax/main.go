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
	resultado, err := minMax(min, max, values)
	if err != nil {
		log.Fatal(err)
	}

	//salida de datos
	fmt.Println("El rango es -->", resultado)
}

func getInput() (float64, float64, []float64) {
	//declaracion de variables
	var min, max float64
	var values []float64

	//entrada de min & max
	fmt.Println("Escribir los valores (mínimo, máximo, lista de valores), separado por espacios")
	fmt.Scan(&min, &max)

	//entrada de values
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	linea := scanner.Text()

	partes := strings.Fields(linea)

	for _, str := range partes {
		num, err := strconv.ParseFloat(str, 64)
		if err != nil {
			fmt.Println("error de conversion")
		}
		values = append(values, num)
	}
	return min, max, values
}

func minMax(min float64, max float64, values []float64) ([]float64, error) {

	if min > max {
		return nil, fmt.Errorf("limites incorrectos")
	}

	var slide []float64
	for _, s := range values {

		if s >= min && s <= max {
			slide = append(slide, s)
		}
	}

	if len(slide) == 0 {
		return nil, fmt.Errorf("slide vacio")
	}
	return slide, nil
}
