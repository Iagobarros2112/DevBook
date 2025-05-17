package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"encoding/json"
	"io"
	"log"

	"net/http"
)

//insere um usuario novo no database

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}

	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		log.Fatal(erro)
	}

	db, erro := banco.Conectar()
	if erro != nil {
		log.Fatal(erro)
	}

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuarioID, erro := repositorio.Criar(usuario)
	if erro != nil {
		log.Fatal(erro)
	}

	w.Write([]byte(fmt.Sprintf("id inserido: %d", usuarioID)))

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
