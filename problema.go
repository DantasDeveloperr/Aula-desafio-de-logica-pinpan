//Objetico do programa: criar um codigo que exiba todos os numeros de 1 a 100, que divisivel por 3.

package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}
