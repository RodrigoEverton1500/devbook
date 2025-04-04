package repositorios

import  (
        "database/sql"
        "api/src/modelos"
        "fmt"
        )
        
type Usuarios struct {
  db *sql.DB
}

func NovoRepositorioUsuarios(db *sql.DB) *Usuarios {
  return &Usuarios{db}
}

func (repositorio Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
  statement, erro := repositorio.db.Prepare("insert into usuarios (nome, nick, email, senha) values(?, ?, ?, ?)",)
  if erro != nil {
    return 0, erro
  }
  defer statement.Close()
  
  resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
  if erro != nil {
    return 0, erro
  }
  
  ultimoIDInserido, erro := resultado.LastInsertId()
  if erro != nil {
    return 0, erro
  }
  
  return uint64(ultimoIDInserido), nil
}

func (repositorio Usuarios) Buscar(nomeOuNick string) ([]modelos.Usuario, error) {
  nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick)
  
  linhas, erro := repositorio.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where nome LIKE ? or nick LIKE ?",
		nomeOuNick, nomeOuNick,
	)
  
  if erro != nil {
    return nil, erro
  }
  defer linhas.Close()
  
  var usuarios []modelos.Usuario
  
  for linhas.Next() {
    var usuario modelos.Usuario
  
    if erro = linhas.Scan (
      &usuario.ID,
      &usuario.Nome,
      &usuario.Nick,
      &usuario.Email,
      &usuario.CriadoEm,
    ); erro != nil {
      return nil, erro
    }
    
    usuarios = append(usuarios, usuario)
  }
  
  return usuarios, nil
}
