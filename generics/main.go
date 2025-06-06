package main

import "fmt"

type NormalUser struct {
	nome string
}

func (n NormalUser) GetNome() string {
	return n.nome
}

type SuperUser struct {
	nome string
}

func (s SuperUser) GetNome() string {
	return s.nome
}

func getName[T interface{ GetNome() string }](user T) string {
	return user.GetNome()
}

func main() {
	var superUser SuperUser = SuperUser{nome: "Super User"}
	var normalUser NormalUser = NormalUser{nome: "Normal User"}

	fmt.Println(getName(superUser))  // Output: Super User
	fmt.Println(getName(normalUser)) // Output: Normal User
}
