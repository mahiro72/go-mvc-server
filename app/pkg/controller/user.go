package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/mahiro72/go-mvc-server/pkg/model"
	"github.com/mahiro72/go-mvc-server/pkg/view"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	idParam, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := model.GetUser(idParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userJson := view.UserToJSON(user)
	if err := json.NewEncoder(w).Encode(userJson); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var j view.User
	if err := json.NewDecoder(r.Body).Decode(&j); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	user, err := model.CreateUser(j.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userJson := view.UserToJSON(user)
	if err := json.NewEncoder(w).Encode(userJson); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
