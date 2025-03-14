// Importação do pacote principal
package main

import "fmt"

// Importação da biblioteca

// Declaração da constante do ponto de ebulição da água em Kelvin
const ebulicaoK float64 = 373

// Função principal
func main() {

	tempK := ebulicaoK     // Variável do valor da temperatura em graus Kelvin
	tempC := (tempK - 273) // Conversão de Kelvin para Celsius

	fmt.Printf("A temperatura de ebulição em graus Kelvin = %v é igual à temperatura em (°C) = %v.", tempK, tempC)
}
