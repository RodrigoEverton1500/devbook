package rotas

import	(
	"net/http"
	"api/src/controllers"
	)

var rotasUsuarios = []Rota {
  {
    URI: 		"/usuarios",
    metodo: 		http.MethodGet,
    funcao:		controllers.GetUsuarios,
    requerAutenticacao:	false,
  },
  {
    URI: 		"/usuarios/{usuarioID}",
    metodo: 		http.MethodGet,
    funcao:		controllers.GetUsuario,
    requerAutenticacao:	false,
  },
  {
    URI: 		"/usuarios/{usuarioID}",
    metodo: 		http.MethodPost,
    funcao:		controllers.PostUsuario,
    requerAutenticacao:	false,
  },
  {
    URI: 		"/usuarios",
    metodo: 		http.MethodPut,
    funcao:		controllers.PutUsuario,
    requerAutenticacao:	false,
  },
  {
    URI: 		"/usuarios/{usuarioID}",
    metodo: 		http.MethodDelete,
    funcao:		controllers.DeleteUsuario,
    requerAutenticacao:	false,
  },
}

