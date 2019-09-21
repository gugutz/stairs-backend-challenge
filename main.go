package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	// imports with the dot on tbhe begginging allow to use its methods without the package name. example: `Method` instead of `package.Method`
	. "github.com/gugutz/stairs-backend-challenge/config"
	. "github.com/gugutz/stairs-backend-challenge/dao"

	// here i declare the variable that holds the imported package right on the import line
	router "github.com/gugutz/stairs-backend-challenge/router"
)

// import database access object from project
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
	r := mux.NewRouter()
	r.HandleFunc("/wines/{id}", router.Get).Methods("GET")
	r.HandleFunc("/wines/{id}/edit", router.Update).Methods("POST")
	r.HandleFunc("/wines", router.GetAll).Methods("GET")
	http.Handle("/", r)
	
	var port = ":3000"
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, r))
}
