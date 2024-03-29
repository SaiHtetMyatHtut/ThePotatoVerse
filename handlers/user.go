package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/SaiHtetMyatHtut/potatoverse/data/models"
	repo "github.com/SaiHtetMyatHtut/potatoverse/data/repositories"
	"github.com/SaiHtetMyatHtut/potatoverse/db"
	authschemas "github.com/SaiHtetMyatHtut/potatoverse/schemas/auth_schemas"
	userschemas "github.com/SaiHtetMyatHtut/potatoverse/schemas/user_schemas"
	"github.com/SaiHtetMyatHtut/potatoverse/utils"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if r.URL.Query().Get("id") == "" {
			GetAllUsers(w, r)
		} else {
			GetUserByID(w, r)
		}
	case http.MethodPost:
		CreateUser(w, r)
	case http.MethodPut:
		UpdateUser(w, r)
	case http.MethodDelete:
		DeleteUser(w, r)
	case http.MethodPatch:
		UpdatePartialUser(w, r)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	// users, err := repo.ReadAll(r.Context())
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// res, err := json.Marshal(users)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// w.WriteHeader(http.StatusOK)
	// w.Write(res)
	userRepository := repo.NewUserRepository(db.NewRedisClient())
	users, err := userRepository.ReadAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userRepository := repo.NewUserRepository(db.NewRedisClient())
	user, err := userRepository.ReadByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// TODO: this is not the best practice
	var userSchema userschemas.UserSchema
	userSchema.ID = user.ID
	userSchema.Username = user.Username
	userSchema.CreatedAt = user.CreatedAt
	userSchema.LastLogin = user.LastLogin

	res, err := json.Marshal(userSchema)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Only For Admin Creation, Other User Creation is in SignUp
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var body authschemas.UserSignUpSchema
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newHashPassword, err := utils.HashPassword(body.Password)

	// TODO: isn't this error should be handled under hashpassword or not?
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	newUser := models.User{
		Username:       body.Username,
		HashedPassword: newHashPassword,
		CreatedAt:      time.Now().UTC(),
		LastLogin:      time.Now().UTC(),
	}
  
	userRepository := repo.NewUserRepository(db.NewRedisClient())
	user, err := userRepository.Insert(r.Context(), newUser)
  
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var body userschemas.UpdateUserSchema
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userRepository := repo.NewUserRepository(db.NewRedisClient())
	user, err := userRepository.ReadAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for _, u := range user {
		if u.ID == body.ID {
			var newUser struct {
				ID             int64     `json:"id"`
				Username       string    `json:"username"`
				HashedPassword string    `json:"password"`
				CreatedAt      time.Time `json:"created_at"`
				LastLogin      time.Time `json:"last_login"`
			}
			newUser.ID = u.ID
			if body.Username != "" {
				newUser.Username = body.Username
			} else {
				newUser.Username = u.Username
			}
			if body.Password != "" {
				newHashPassword, err := utils.HashPassword(body.Password)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				newUser.HashedPassword = newHashPassword
			} else {
				newUser.HashedPassword = u.HashedPassword
			}
			err = userRepository.Update(r.Context(), newUser)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			var response struct {
				ID        int64     `json:"id"`
				Username  string    `json:"username"`
				CreatedAt time.Time `json:"created_at"`
				LastLogin time.Time `json:"last_login"`
			}
			response.ID = newUser.ID
			response.Username = newUser.Username
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

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Implement delete user logic here
	var body struct {
		ID       int64  `json:"id"`
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userRepository := repo.NewUserRepository(db.NewRedisClient())
	user, err := userRepository.ReadAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for _, u := range user {
		if u.ID == body.ID && u.Username == body.Username && utils.VerifyPassword(body.Password, u.HashedPassword) {
			err = userRepository.Delete(r.Context(), body.ID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write([]byte("User deleted successfully"))
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "User not found", http.StatusNotFound)
}

// TODO
func UpdatePartialUser(w http.ResponseWriter, r *http.Request) {
	// Implement update partial user logic here
	var user struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Update the user partially in the database
	w.WriteHeader(http.StatusOK)
}
