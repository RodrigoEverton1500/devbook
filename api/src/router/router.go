package router

import	(
	"api/src/router/rotas"
	"github.com/gorilla/mux"
	)
	
//Router configurado é gerado e retornado
func Gerar() *mux.Router {
  r := mux.NewRouter()
  return rotas.Configurar(r)
}
