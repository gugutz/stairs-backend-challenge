package router

import (
	"fmt"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/gugutz/stairs-backend-challenge/dao"
	"github.com/gugutz/stairs-backend-challenge/models"
)

var dao = DAO{}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	fmt.Println("HTTP Status Code for the request: ", code)
	fmt.Println("JSON Response Object: ", payload)
	w.Write(response)
}

func Create (w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var wine models.Wine
	err := json.NewDecoder(r.Body).Decode(&wine)
	if err != nil {
		fmt.Printf("Error generating payload: %+v \n", err)
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	inserted_wine, err := dao.Create(wine)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, inserted_wine)
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	wines, err := dao.GetAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, &wines)
}

func Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	result, err := dao.Get(id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, result)
}

func Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	id := params["id"]

	var wine models.Wine
	wine.Id = id
	err := json.NewDecoder(r.Body).Decode(&wine)
	if err != nil {
		fmt.Printf("Error generating payload: %+v", err)
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	_, err = dao.Update(id, wine)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": wine.Name + " atualizado com sucesso!"})
}

func DeleteOne(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	id := params["id"]

	_, err := dao.DeleteOne(id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "deleted successfully"})
}

func RootHandler(w http.ResponseWriter, r *http.Request){
	fmt.Printf("Response: %v\n", http.StatusOK)
	w.Write([]byte("Root Handler works!!\n"))
}

