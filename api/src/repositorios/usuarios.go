package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

//usuarios representa um repositorio de usuarios

type Usuarios struct {
	db *sql.DB
}

// cria um novo repositorio de usuarios
func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

// insere um usuario no banco de dados
func (u Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	return 0, nil
}
