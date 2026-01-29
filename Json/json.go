package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Usuarios struct {
	Usuarios []Usuario `json:"usuarios"`
}

type Usuario struct {
	Nome   string `json:"nome"`
	Idade  int    `json:"idade"`
	Email  string `json:"email"`
	Social Social `json:"social"`
}

type Social struct {
	Twitter  string `json:"twitter"`
	Facebook string `json:"facebook"`
}

func main() {
	JsonFile, err := os.Open("usuarios.Json")
	if err != nil {
		fmt.Println(err)

	}
	defer JsonFile.Close()
	fmt.Println("Arquivo Json aberto com sucesso")

	byteValue, err := ioutil.ReadAll(JsonFile)
	if err != nil {
		fmt.Println(err)
	}

	var usuarios Usuarios
	json.Unmarshal(byteValue, &usuarios)
	fmt.Println(usuarios)

	for i := 0; i < len(usuarios.Usuarios); i++ {
		fmt.Println("Usuario Nome: " + usuarios.Usuarios[i].Nome)
	}

}
