package rotas

import	(
	"net/http"
	"github.com/gorilla/mux"
	)

type Rota struct {
  URI			string
  metodo		string
  funcao		func(http.ResponseWriter, *http.Request)
  requerAutenticacao	bool
}

func Configurar(r *mux.Router) *mux.Router {
  rotas := rotasUsuarios

  for _, rota := range rotas{
    r.HandleFunc(rota.URI, rota.funcao).Methods(rota.metodo)
  }
  
  return r
}
