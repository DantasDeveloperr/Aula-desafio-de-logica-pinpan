//Recriar uma antiga brincadeira, escrever um codigo que pega um numero de 1 a 100, que nos numeros que são divisiveis por 3 reportanar pin e nos numeros divisiveis por 5 retornar pan, e nos numeros divisiveis por ambos retornar pinpan.

package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("pinpan")	
		} else if i%3 == 0 {
			fmt.Println("pin")
		}
		 else if i%5 == 0 {
			fmt.Println("pan")
		} else {
			fmt.Println(i)
		}	

	}
}

	