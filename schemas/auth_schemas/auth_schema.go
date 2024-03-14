package authschemas

type UserSignInSchema struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserSignUpSchema struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserSignInResponseSchema struct {
	Username string `json:"username"`
	Jwt      string `json:"jwt"`
}
