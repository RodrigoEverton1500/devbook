package main

import 	(
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
	)
	
func main() {
  config.Carregar()
  
  fmt.Println(config.Porta)
  fmt.Println("rodando API")
  
  r := router.Gerar()
  log.Fatal(http.ListenAndServe(":5000", r))
}
