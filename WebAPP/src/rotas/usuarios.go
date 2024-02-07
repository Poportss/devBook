package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var rotasUsarios = []Rota{

	{
		URI:                "/",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeCadastroDeUsuario,
		RequerAutenticacao: false,
	},
}
