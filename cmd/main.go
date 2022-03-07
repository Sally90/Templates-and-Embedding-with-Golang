package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/Sally90/Embedding2/routes"
)

func main(){
	r := mux.NewRouter()
	routes.RegisterRoutes(r)
	log.Fatal(http.ListenAndServe(":5000", r))
}


