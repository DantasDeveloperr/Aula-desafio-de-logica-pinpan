// método é uma função associada a um tipo específico em Go.
package main

import "fmt"

type retangulo struct {
	largura, altura int
}

// Área é um método que calcula a área do retângulo.
func (r retangulo) area() int {
	return r.largura * r.altura
}

// Perímetro é um método que calcula o perímetro do retângulo.
func (r retangulo) perímetro() int {
	return 2 * (r.largura + r.altura)
}

func main() {
	r := retangulo{largura: 10, altura: 5}
	fmt.Println("area: ", r.area())
	fmt.Println("perímetro: ", r.perímetro())
}
