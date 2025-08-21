package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var saltos string
	if len(os.Args) > 1 {
		saltos = os.Args[1]
	} else {
		saltos = ""
	}
	if saltos == "-l" {
		fmt.Println("Modo conteo de líneas. Escribe tu texto (una o varias líneas). Escribe 'exit' para terminar:")
	} else {
		fmt.Println("Modo conteo de palabras. Escribe tu texto (una o varias líneas). Escribe 'exit' para terminar:")
	}

	contadorPalabras := 0
	contadorSaltos := 0
	salida := false

	for !salida {
		texto, _ := reader.ReadString('\n')
		frase := strings.TrimSpace(texto)

		if frase == "exit" {
			salida = true
			continue
		}

		if saltos == "-l" {
			contadorSaltos++
		} else {
			words := strings.Fields(texto)
			for _, word := range words {
				if word != "exit" {
					contadorPalabras++
				}
			}
		}
	}

	if saltos == "-l" {
		fmt.Println("Número de líneas:", contadorSaltos)
	} else {
		fmt.Println("Número de palabras:", contadorPalabras)
	}
}
