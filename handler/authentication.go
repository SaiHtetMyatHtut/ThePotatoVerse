package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/SaiHtetMyatHtut/potatoverse/model"
	"github.com/SaiHtetMyatHtut/potatoverse/repo"
)

func AuthenticationHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetUserByID(w, r)
	case http.MethodPost:
		CreateUser(w, r)
	case http.MethodPut:
	case http.MethodDelete:
	case http.MethodPatch:
	default:
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newUser := model.User{
		Username:  body.Username,
		Password:  body.Password,
		CreatedAt: time.Now().UTC(),
		LastLogin: time.Now().UTC(),
	}
	// TODO
	err := repo.Insert(r.Context(), newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	const base = 10
	bitSize := 64
	userID, err := strconv.ParseInt(idParam, base, bitSize)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO int to int64
	user, err := repo.ReadByID(r.Context(), int(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
