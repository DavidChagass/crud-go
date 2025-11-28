package models

import (
	"errors"
	"net/mail"
	"strings"
)

type User struct {
	//models para o projeto (colocar mais atributos para o user depois)
	ID    int
	Nome  string
	Email string
}

func (us *User) FormatarString() {
	//retirar espaços em branco para o nome
	us.Nome = strings.TrimSpace(us.Nome)
	//retirar espaços em branco para o email
	//lowercase para o email
	us.Email = strings.ToLower(strings.TrimSpace(us.Email))
}

// função para validar os dados de email e nome
// caso fique muito grande futuramente , criar um service somente para isto
func (us *User) ValidarModel() error {
	//rodar a formatação
	us.FormatarString()

	//teste de nome e email nulo
	if us.Nome == "" {
		return errors.New("o nome não pode ser vazio")
	}

	if us.Email == "" {
		return errors.New("o email não pode ser vazio")

	}
	//testar se o email está em um formato valido
	if _, err := mail.ParseAddress(us.Email); err != nil {
		return errors.New("insira um email valido")
	}

	//caso não dê erro, retorna nulo
	return nil
}
