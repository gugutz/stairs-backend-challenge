package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	. "github.com/gugutz/stairs-backend-challenge/config"
	. "github.com/gugutz/stairs-backend-challenge/dao"

	router "github.com/gugutz/stairs-backend-challenge/router"
)

var dao = DAO{}
var config = Config{}

func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Collection = config.Collection
	dao.Connect()
}


func main() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/wines/new", router.Create).Methods("POST")
	r.HandleFunc("/wines", router.GetAll).Methods("GET")
	r.HandleFunc("/wines", router.GetAll).Methods("POST")
	r.HandleFunc("/wines/{id}", router.Get).Methods("GET")
	r.HandleFunc("/wines/{id}/edit", router.Update).Methods("PUT")
	r.HandleFunc("/wines/{id}", router.DeleteOne).Methods("DELETE")
	r.HandleFunc("/", router.RootHandler).Methods("GET")
	
	var port = ":3000"
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, r))
}

