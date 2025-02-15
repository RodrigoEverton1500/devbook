package controllers

import	(
        "api/src/repositorios"
        "api/src/respostas"
        "api/src/banco"
        "api/src/modelos"
	"net/http"
	"io/ioutil"
	"log"
	"fmt"
	"encoding/json"
	)
	
func GetUsuarios(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("buscando usuarios"))
}

func GetUsuario(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("buscando usuario"))
}

func PostUsuario(w http.ResponseWriter, r *http.Request) {
  corpoRequest, erro := ioutil.ReadAll(r.Body)
  if erro != nil {
    respostas.Erro(w, http.StatusUnprocessableEntity, erro)
    return
  }
  
  var usuario modelos.Usuario
  if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
    respostas.Erro(w, http.StatusBadRequest, erro)
    return
  }
  
  if erro = usuario.Preparar(); erro != nil {
    respostas.Erro(w, http.StatusBadRequest, erro)
    return
  }
  
  db, erro := banco.Conectar()
  if erro != nil {
    respostas.Erro(w, http.StatusInternalServerError, erro)
    return
  }
  
  repositorio := repositorios.NovoRepositorioUsuarios(db)
  usuario.ID, erro := repositorio.Criar(usuario)
  if erro != nil {
    respostas.Erro(w, http.StatusInternalServerError, erro)
    return
  }
  
  respostas.JSON(w, http.StatusCreated, usuario)
}

func PutUsuario(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("criando usuario"))
}

func DeleteUsuario(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("deletando usuario"))
}
