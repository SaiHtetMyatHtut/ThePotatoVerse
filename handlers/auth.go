package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/SaiHtetMyatHtut/potatoverse/data/models"
	repo "github.com/SaiHtetMyatHtut/potatoverse/data/repositories"
	"github.com/SaiHtetMyatHtut/potatoverse/db"
	authschemas "github.com/SaiHtetMyatHtut/potatoverse/schemas/auth_schemas"
	"github.com/SaiHtetMyatHtut/potatoverse/utils"
	"github.com/golang-jwt/jwt/v5"
)

func SignIn(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid Request Method", http.StatusMethodNotAllowed)
		return
	}

	var body authschemas.UserSignInSchema
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userRepository := repo.NewUserRepository(db.NewRedisClient())
	users, err := userRepository.ReadAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, u := range users {
		if u.Username == body.Username && utils.VerifyPassword(u.HashedPassword, body.Password) {
			// log.Printf("User %s has logged in", u.Username)
			token := jwt.New(jwt.SigningMethodHS256)
			claims := token.Claims.(jwt.MapClaims)
			claims["id"] = u.ID
			claims["username"] = u.Username
			claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
			tokenString, err := token.SignedString([]byte("PotatoSecret"))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			var response authschemas.UserSignInResponseSchema

			response.Username = u.Username
			response.Jwt = tokenString
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

func SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid Request Method", http.StatusMethodNotAllowed)
		return
	}

	var body authschemas.UserSignUpSchema
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hashedPassword, err := utils.HashPassword(body.Password)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newUser := models.User{
		Username:       body.Username,
		HashedPassword: hashedPassword,
		CreatedAt:      time.Now().UTC(),
		LastLogin:      time.Now().UTC(),
	}

	// TODO: this should use DI. not manually inject.
	userRepository := repo.NewUserRepository(db.NewRedisClient())
	user, err := userRepository.Insert(r.Context(), newUser)
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

// func RefreshJwt(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodGet {
// 		http.Error(w, "Invalid Request Method", http.StatusMethodNotAllowed)
// 		return
// 	}
// 	cookie, err := r.Cookie("jwt")
// 	if err != nil {
// 		http.Error(w, "Unauthorized", http.StatusUnauthorized)
// 		return
// 	}
// 	token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
// 		// Don't forget to validate the alg is what you expect:
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, jwt.ErrInvalidKeyType
// 		}
// 		return []byte("PotatoSecret"), nil
// 	})
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusUnauthorized)
// 		return
// 	}
// 	if oldToken, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 		newToken := jwt.New(jwt.SigningMethodHS256)
// 		claims := newToken.Claims.(jwt.MapClaims)
// 		claims["id"] = oldToken["id"] // claims["id"]
// 		claims["username"] = oldToken["username"]
// 		claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
// 		tokenString, err := newToken.SignedString([]byte("PotatoSecret"))
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		var response struct {
// 			Username string `json:"username"`
// 			Jwt      string `json:"jwt"`
// 		}

// 		response.Username = claims["username"].(string)
// 		response.Jwt = tokenString
// 		res, err := json.Marshal(response)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		w.WriteHeader(http.StatusOK)
// 		w.Write(res)
// 		return
// 	}
// 	http.Error(w, "Unauthorized", http.StatusUnauthorized)
// }
