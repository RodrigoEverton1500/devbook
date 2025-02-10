package banco

import  (
        "database/sql"
        "github.com/go-sql-driver/mysql"
        )
        
func Conectar() (*sql.DB, error) {
  db, erro := sq.Open("mysql", config.StringConexaoBanco)
  if erro != nil {
    return nil, erro
  }
  
  if erro := db.Ping(); erro != nil {
    db.Close()
    return nil, erro
  }
  
  return db, nil
}
