package handler

import (
	"net/http"

	"github.com/golang-jwt/jwt"
)

func JWTHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
	case http.MethodPost:
	case http.MethodPut:
	case http.MethodDelete:
	case http.MethodPatch:
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func GenerateJWTByClaim(w http.ResponseWriter, r *http.Request) {
	token := jwt.New(jwt.GetSigningMethod("RS256"))
	// token.Claims = &
}
