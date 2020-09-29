package main

// Importando nuestras bibliotecas.
import "fmt"

// Creamos una funcion principal.
func main() {
	// Declaramos una variable de tipo float64.
	var side uint64

	fmt.Scanf("%d", &side) // Tomamos de entrada un valor numérico.
	// Formula para sacar el área de un cuadrado.
	res := side * side
	fmt.Println(res)
}