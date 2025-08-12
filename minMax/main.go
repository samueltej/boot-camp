package main

import (
	"fmt"
	"bufio"
	"strconv"
	"strings"
	"os"
)
func main()  {

    min, max, values := getInput()
	resultado := minMax(min, max, values)

	//salida de datos 
	fmt.Println("EL rango es -->", resultado)
}

func getInput()(float64, float64, []float64){
	//declaracion de variables
	var min, max float64
	var values [] float64

	//entrada de min & max
	fmt.Println("Escribir los valores (mínimo, máximo, lista de valores), separado por espacios")
	fmt.Scan(&min, &max)

	//entrada de values
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	linea := scanner.Text()

	partes := strings.Fields(linea)

	for _, str := range partes{
		num, err := strconv.ParseFloat(str, 64)
		if err != nil {
			fmt.Println("error de conversion")
		} 
	  	values = append(values, num)
	}
	return min, max, values
}

func minMax (min float64, max float64, values[] float64) [] float64{
	var slide[] float64
	for _, s := range values{

		if s >= min && s <= max{
			slide = append(slide, s)
		}
	}	
	return slide
}