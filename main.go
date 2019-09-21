package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	. "github.com/gugutz/stairs-backend-challenge/router"
)


func YourHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Gorilla!\n"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/wines/{id}", WineHandler).Methods(GET)
	r.HandleFunc("/wines/{id}/edit", WineHandler)
	r.HandleFunc("/wines", ArticlesCategoryHandler)
	r.HandleFunc("/wines/{category}/{id:[0-9]+}", ArticleHandler)
	r.HandleFunc("/", HomeHandler)
	// http.Handle("/", r)
	http.Handle("/", YourHandler)
	
	var port = ":3000"
	fmt.Println("Server running in port:", port)
	log.Fatal(srv.ListenAndServe())

}
