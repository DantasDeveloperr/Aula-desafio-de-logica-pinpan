//Esta API ainda não vai ter uma finalidade de algo espeficico, vou apenas deixar a estrutura desta API pronta.
/* Esta API terá:
GET
POST
DELETE
PUT
PATH
*/

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("Por enquanto é isso!!")

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()
}
