package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("escribe la frase")

	var saltos string
	if len(os.Args) > 1 {
		saltos = os.Args[1]
	} else {
		saltos = ""
	}

	contadorPalabras := 0
	contadorSaltos := 0
	salida := false

	for !salida {
		texto, _ := reader.ReadString('\n')
		contadorSaltos++
		words := strings.Fields(texto)
		for _, word := range words {
			if word != "exit" {
				contadorPalabras++
			} else {
				salida = true
				contadorSaltos = contadorSaltos - 1
			}
		}
	}

	switch saltos {
	case "-l":
		fmt.Println("numero de saltos", contadorSaltos)
	case "":
		fmt.Println("numero de palabras", contadorPalabras)
	default:
		fmt.Println("Opción no reconocida")
	}

}
