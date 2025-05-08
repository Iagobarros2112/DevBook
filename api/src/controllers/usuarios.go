package controllers

import "net/http"

//cria um usuario novo no database

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("criando usuario"))
}

// busca usuarios no database
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("buscando usuarios"))
}

// busca um usuario em expecifico
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("buscando usuario"))
}

// atualiza as info de um usuario
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("atualizando usuario"))
}

// deleta um usuario em expecifico
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("deletando usuario"))
}
