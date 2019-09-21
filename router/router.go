package router


import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/gugutz/stairs-backend-challenge/dao"
	. "github.com/gugutz/stairs-backend-challenge/models"
)

var dao = DAO{}

func ArticlesCategoryHandler(w http.ResponseWriter, r *http.Request) {
	// mux.Vars(r) gets the variables (params) from the request
	// it passes the variables used in the routes
    vars := mux.Vars(r)
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Category: %v\n", vars["category"])
}
