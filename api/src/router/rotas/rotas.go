package rotas

import "net/http"

// rota eh uma rota generica
// que representa todas as rotas da aplicaçao
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}
