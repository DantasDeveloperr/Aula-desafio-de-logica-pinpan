package main

import "fmt"

func soma(a int, b int, c int) int {
	return a + b + c
}

func main() {
	resultado := soma(3, 5, 6)
	fmt.Println("A soma é:", resultado)
}

//Isto é um exemplo de teste unitário.
/*
package main

import "testing"

func TestSoma(t *testing.T) {
	teste := soma(3, 5, 6)
	result := 14 //se eu trocar este resultado esperado, vai me retornar um erro de  logica no teste.
                 //exemplo: 15 vai me dar um erro de valor esperado diferente do valor retornado.
	if teste != result {
		t.Errorf("Valor esperado: %d, Valor retornado: %d", result, teste)
		}
*/

// uma forma de padronizar é com o padrão AAA -> padrão triple A

/*
package main
import "testing"
func TestSoma(t *testing.T) {
	// Arrange
	a := 3
	b := 5
	c := 6
	expected := 14
	// Act
	result := soma(a, b, c)
	// Assert
	if result != expected {
		t.Errorf("Valor esperado: %d, Valor retornado: %d", expected, result)
	}
}
*/

//go test -v soma.go soma_test.go -> comando para rodar o teste no terminal.
