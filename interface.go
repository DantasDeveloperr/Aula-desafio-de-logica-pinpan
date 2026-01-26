package main

import (
	"fmt"
	"math"
)

type geometrica interface {
	area() float64
	perimetro() float64
}

type quadrado struct {
	lado float64
}
type circulo struct {
	raio float64
}

func (q quadrado) area() float64 {
	return q.lado * q.lado
}
func (q quadrado) perimetro() float64 {
	return q.lado * 4
}
func (c circulo) area() float64 {
	return math.Pi * c.raio * c.raio
}
func (c circulo) perimetro() float64 {
	return 2 * math.Pi * c.raio
}
func medir(g geometrica) {
	fmt.Println("Área:", g.area())
	fmt.Println("Perímetro:", g.perimetro())
}
func testarGeometrica() {
	q := quadrado{lado: 3}
	c := circulo{raio: 5}

	medir(q)
	medir(c)
	fmt.Println("Programa encerrado.")
}
