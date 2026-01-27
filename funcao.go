// função yambém é chamada de procedimento ou sub-rotina
// parte fundamental da programação estruturada
// ela pega um bloco de código e o encapsula em uma unidade lógica

package main

import "fmt"

//vamos criar um programa que caucula a media de uma sala de aula.
// iremos precisar de uma lista.

/*

func media(lista []float64) float64 {
	totalAlunos := 0.0

	for _, valor := range lista {
		totalAlunos += valor
	}
		return totalAlunos / float64(len(lista))
}

Tanto esta forma ou a de baixo vai me dar o mesmo resultado.
*/

func main() {
	lista := []float64{5.5, 6.5, 7.0, 8.0, 9.5, 10.0}

	totalAlunos := 0.0

	for _, valor := range lista {
		totalAlunos += valor
	}
	fmt.Println(totalAlunos / float64(len(lista)))

}
