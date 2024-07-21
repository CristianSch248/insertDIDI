package main

import (
	"fmt"
	"insertDiploma/utilitarios"
)

func main() {

	var user string
	var senha string
	var qtd int

	cpfs := []string{}

	// Captura o usuário
	fmt.Print("Digite seu usuário do Asten Diploma Digital: ")
	fmt.Scanln(&user)

	// Captura a senha
	fmt.Print("Digite sua senha do Asten Diploma Digital: ")
	fmt.Scanln(&senha)

	fmt.Print("Quantas inserções deseja fazer? ")
	fmt.Scanln(&qtd)

	for i := 0; i < qtd; i++ {
		cpfs = append(cpfs, utilitarios.GetCPF())
	}

	token := utilitarios.GetToken(user, senha)

	utilitarios.Insert(cpfs, token)
}
