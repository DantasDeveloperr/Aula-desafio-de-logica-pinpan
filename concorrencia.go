//Goroutines: é uma funç]ao de executar de modo concorrente em outra funçãos.

package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		go fmt.Println(n, ":", i)
	}

}
func main() {
	go f(0)
	var escreva string
	fmt.Scanln(&escreva)
}

// canal: modo de duas vias de comunicação entre goroutines.

/*
package main

import (
	"fmt"
	"time"
)

func pingar(c chan string) {
	for i := 0; ; i++{
		c <- "ping"
	}
}
func imprimir(c chan string) {
	for {
		msg := <- c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}
func main() {
	var c chan string = make(chan string)
	go pingar(c)
	go imprimir(c)

	var entrada string
	fmt.Scanln(&entrada)
}

*/
