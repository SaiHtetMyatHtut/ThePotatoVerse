package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/SaiHtetMyatHtut/potatoverse/model"
)

func SignIn(w http.ResponseWriter, r *http.Request) {

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

	resp, err := http.Get("http://localhost:8081/user")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var bodyBytes []byte
	if resp.Body == nil {
		http.Error(w, "Empty Response Body", http.StatusInternalServerError)
		return
	}
	bodyBytes, err = io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var users []model.User

	err = json.Unmarshal(bodyBytes, &users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, u := range users {
		if u.Username == body.Username && CheckPasswordHash(body.Password, u.Password) {
			log.Printf("User %s has logged in", u.Username)
			var response struct {
				Username  string    `json:"username"`
				CreatedAt time.Time `json:"created_at"`
				LastLogin time.Time `json:"last_login"`
			}
			response.Username = u.Username
			response.CreatedAt = u.CreatedAt
			res, err := json.Marshal(response)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write(res)
			return
		}
	}
	http.Error(w, "Invalid Username or Password", http.StatusUnauthorized)
}
