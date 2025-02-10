package controllers

import	(
        "api/src/banco"
        "api/src/modelos"
	"net/http"
	"io/ioutil"
	"log"
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
    log.Fatal(erro)
  }
  
  var usuario modelos.Usuario
  if erro = json.Unmarshall(corpoRequest, &usuario); erro != nil {
    log.Fatal(erro)
  }
  
  db, erro := banco.Conectar()
  if erro != nil {
    log.Fatal(erro)
  }
}

func PutUsuario(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("criando usuario"))
}

func DeleteUsuario(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("deletando usuario"))
}
