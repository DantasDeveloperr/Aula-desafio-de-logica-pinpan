// pacotes são caixinhas de função.
// função contains: contains procura por uma substring dentro de uma string maior ou menor.
// vamos iniciar falando de strings.Contains.
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("computador", "dor"))
}

// outros exemplos é o pacote io e o bytes.
/*

import ("io"
        "log"
		"os")

func main() {
message := []byte("Hello, Gophers!")
err := ioutil.WriteFile("hello", message, 0644)
if err != nil {
	log.Fatal(err)
}


Também podemos falar um pouco do pacote bytes. vamos dar um exemplo:
import (
	"bytes"
	"fmt"
)

func main() {
	b := []byte("Hello, Gophers!")
	c := bytes.NewBuffer(b)
	fmt.Println(c.String())
}


Porém o principal mesmo é ler a documentação oficial do Go: https://pkg.go.dev/std pós só assim para entender e saber de fato todo o que a liguagem pode lhe proporcionar.


*/
