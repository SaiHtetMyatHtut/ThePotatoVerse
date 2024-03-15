package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/SaiHtetMyatHtut/potatoverse/src/core/data/models"
	repo "github.com/SaiHtetMyatHtut/potatoverse/src/core/data/repositories"
	"github.com/SaiHtetMyatHtut/potatoverse/src/db"
	authschemas "github.com/SaiHtetMyatHtut/potatoverse/src/schemas/auth_schemas"
	"github.com/SaiHtetMyatHtut/potatoverse/src/utils"
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
			log.Printf("User %s has logged in", u.Username)
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

func RefreshJwt(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid Request Method", http.StatusMethodNotAllowed)
		return
	}

	bearerToken := r.Header.Get("Authorization")

	token, err := jwt.Parse(strings.TrimPrefix(bearerToken, "Bearer "), func(token *jwt.Token) (interface{}, error) {
		// TODO Validate the your signing algorithm here
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrInvalidKeyType
		}
		log.Println("The Token is ", token)
		log.Println("The Token is 2 ", string([]byte{160, 246, 104, 83, 166, 14, 72, 239, 225, 39, 211, 178, 29, 191, 23, 242, 150, 172, 200, 197, 93, 36, 195, 7, 98, 39, 123, 219, 155, 32, 190, 83}))
		return []byte("PotatoSecret"), nil
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	if oldToken, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check if the signature is different
		// if oldToken["signature"] != "PotatoSecret" {
		// 	http.Error(w, "Invalid Signature", http.StatusUnauthorized)
		// 	return
		// }
		newToken := jwt.New(jwt.SigningMethodHS256)
		claims := newToken.Claims.(jwt.MapClaims)
		claims["id"] = oldToken["id"] // claims["id"]
		claims["username"] = oldToken["username"]
		claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
		tokenString, err := newToken.SignedString([]byte("PotatoSecret"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var response struct {
			Username string `json:"username"`
			Jwt      string `json:"jwt"`
		}

		response.Username = claims["username"].(string)
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
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
}
