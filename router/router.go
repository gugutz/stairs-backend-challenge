package router

import (
	"fmt"
	// "encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/gugutz/stairs-backend-challenge/dao"
	// . "github.com/gugutz/stairs-backend-challenge/models"
)

var dao = DAO{}

func GetAll(w http.ResponseWriter, r *http.Request) {
	// mux.Vars(r) gets the variables (params) from the request
	// it passes the variables used in the routes
    vars := mux.Vars(r)
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Category: %v\n", vars["category"])
}

func Get(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Get works!!\n"))
}

func Update(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Update Called!\n"))
}
