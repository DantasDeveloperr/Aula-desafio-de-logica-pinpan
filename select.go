//Select: ele funciona como switch para canais.

package main

import (
	"fmt"
	"time"
)

func main(){
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "mensagem do canal 1"
			time.Sleep(time.Second * 2)
		}
}()
	go func() {
		for {
			c2 <- "mensagem do canal 2"
			time.Sleep(time.Second * 3)
		}
	}()

	for i := 0; i < 2; i++ {
		slect {
		case msg1 := <- c1:
			fmt.Println(msg1)
		case msg2 := <- c2:
			fmt.Println(msg2)
		}
	}

}