package main

import (
	"fmt"
	"net/http"

	"github.com/DjonatasBorges/api-user/routes"
)

func main() {
	fmt.Println("Iniciando o servidor Rest com Go")

	routes.HandleRequests()

	http.ListenAndServe(":8080", nil)
}
