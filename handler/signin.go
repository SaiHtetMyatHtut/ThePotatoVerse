package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/SaiHtetMyatHtut/potatoverse/repo"
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

	user, err := repo.ReadAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for _, u := range user {
		if u.Username == body.Username && CheckPasswordHash(body.Password, u.HashedPassword) {
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
