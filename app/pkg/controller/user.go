package controller

import (
	"encoding/json"
	"net/http"

	"github.com/mahiro72/go-mvc-server/pkg/model"
	"github.com/mahiro72/go-mvc-server/pkg/view"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	idQuery := r.URL.Query().Get("id")
	user,err := model.GetUser(idQuery)
	if err != nil {
		http.Error(w,err.Error(),http.StatusBadRequest)
		return
	}
	userJson := view.UserToJSON(user)
	if err := json.NewEncoder(w).Encode(userJson);err!=nil{
		http.Error(w,err.Error(),http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var j view.UserJSON
	if err := json.NewDecoder(r.Body).Decode(&j);err!=nil{
		http.Error(w,err.Error(),http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	user,err := model.CreateUser(j.Name)
	if err != nil {
		http.Error(w,err.Error(),http.StatusBadRequest)
		return
	}
	userJson := view.UserToJSON(user)
	if err := json.NewEncoder(w).Encode(userJson);err!=nil{
		http.Error(w,err.Error(),http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
