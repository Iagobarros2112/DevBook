package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("rodando api")

	r := router.Gerar()

	log.Fatal(http.ListenAndServe(":5000", r))
}
