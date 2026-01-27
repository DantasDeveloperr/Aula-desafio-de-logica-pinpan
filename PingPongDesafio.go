// Este desafio é para criarmos um codigo em Go utilizando concorrencia, canais e goroutines para exibir em alternancia, as palavras "ping" e "pong" no console.

package main

import (
	"fmt"
	"time"
)

func ping(c chan string) {
	for {
		c <- "ping"
		time.Sleep(time.Second)
	}
}

func pong(c chan string) {
	for {
		c <- "pong"
		time.Sleep(time.Second)
	}
}
func main() {
	c := make(chan string)
	go ping(c)
	go pong(c)

	for i := 0; i < 10; i++ {
		msg := <-c
		fmt.Println(msg)
	}

}

// canal: modo de duas vias de comunicação entre goroutines.
