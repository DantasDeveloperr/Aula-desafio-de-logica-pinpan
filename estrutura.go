// Estruturas são coleções de "campos". Elas são usadas para agrupar dados relacionados juntos para formar registros.
// São usadas para agrupar dados e formar registros.

package main

import "fmt"

type Pessoa struct {
	name string
	age  int
}

func main() {
	p1 := Pessoa{name: "Alice", age: 30}
	p2 := Pessoa{name: "Bob", age: 25}
	fmt.Println(p1)
	fmt.Println(p2)
}
