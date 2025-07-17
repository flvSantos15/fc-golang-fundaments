package struct_types

import "fmt"

type Cliente struct {
	Nome string
	Email string
	CPF int
}

type ClientInternacional struct {
	Cliente
	Pais string
}

func NovoCliente() {
	client := Cliente{
		Nome: "Flavio",
		Email: "test@email",
		CPF: 000000000,
	}
	fmt.Println(client)

	client1 := ClientInternacional{
		Cliente: Cliente{
			Nome: "Davi",
			Email: "doe@email",
			CPF: 12312312300,
		},
		Pais: "EUA",
	}
	fmt.Println(client1)
}