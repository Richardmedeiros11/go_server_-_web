package models

import "github.com/Richardmedeiros11/go_server_-_web/db"

type Produtos struct {
	Id         int
	Nome       string
	Descrição  string
	Preco      float64
	Quantidade int
}

func BuscaProdutos() []Produtos {
	db := db.ConectaComBancoDeDados()

	selectDeTodosOsProdutos, err := db.Query("select * from produtos")

	if err != nil {
		panic(err.Error())
	}
	p := Produtos{}
	produtos := []Produtos{}

	for selectDeTodosOsProdutos.Next() {
		var id, quantidade int
		var nome, descrição string
		var preco float64

		err = selectDeTodosOsProdutos.Scan(&id, &nome, &descrição, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Nome = nome
		p.Descrição = descrição
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos
}
