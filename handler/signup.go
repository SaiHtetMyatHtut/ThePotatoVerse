package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/SaiHtetMyatHtut/potatoverse/model"
	"github.com/SaiHtetMyatHtut/potatoverse/repo"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid Request Method", http.StatusMethodNotAllowed)
		return
	}
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hashedPassword, err := HashPassword(body.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newUser := model.User{
		Username:       body.Username,
		HashedPassword: hashedPassword,
		CreatedAt:      time.Now().UTC(),
		LastLogin:      time.Now().UTC(),
	}
	// TODO
	user, err := repo.Insert(r.Context(), newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var response struct {
		Username string `json:"username"`
		Message  string `json:"message"`
	}
	response.Username = user.Username
	response.Message = "User Created Successfully"

	res, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}
