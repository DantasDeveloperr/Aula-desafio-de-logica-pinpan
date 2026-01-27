//recursão é a capaciadade de uma função chamar ela mesma.
//calular fatorial é um exemplo clássico de recursão.

package main

import "fmt"

func fatorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * fatorial(n-1)
}

func main() {

	fmt.Println(fatorial(10))
	fmt.Println(fatorial(11))
	fmt.Println(fatorial(8))
	fmt.Println(fatorial(6))
	fmt.Println(fatorial(4))
	fmt.Println(fatorial(22))
	fmt.Println(fatorial(99))
	fmt.Println(fatorial(25))

}
