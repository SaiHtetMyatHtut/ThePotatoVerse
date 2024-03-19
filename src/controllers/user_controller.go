package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/SaiHtetMyatHtut/potatoverse/src/core/domain/services"
	authschemas "github.com/SaiHtetMyatHtut/potatoverse/src/schemas/auth_schemas"
	userschemas "github.com/SaiHtetMyatHtut/potatoverse/src/schemas/user_schemas"
	"go.uber.org/dig"
)

type UserController struct {
	userService services.UserService
}

type UserControllerDependencies struct {
	dig.In

	UserService services.UserService `name:"UserService"`
}

func NewUserController(deps UserControllerDependencies) UserController {
	return UserController{
		userService: deps.UserService,
	}
}

func (us *UserController) Exec(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if r.PathValue("id") != "" {
			us.getUserById(w, r)
		} else {
			us.getAllUsers(w, r)
		}
	case http.MethodPost:
		us.createUser(w, r)
	case http.MethodPut:
		us.updateUser(w, r)
	case http.MethodDelete:
		us.deleteUser(w, r)
	case http.MethodPatch:
		us.updatePartialUser(w, r)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func (us *UserController) getAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := us.userService.GetAllUsers(r.Context())
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

func (us *UserController) getUserById(w http.ResponseWriter, r *http.Request) {
	idPathValue := r.PathValue("id")
	id, err := strconv.ParseInt(idPathValue, 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := us.userService.GetUserById(r.Context(), id)
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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (us *UserController) createUser(w http.ResponseWriter, r *http.Request) {
	var body authschemas.UserSignUpSchema
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := us.userService.CreateUser(r.Context(), body)

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

func (us *UserController) updateUser(w http.ResponseWriter, r *http.Request) {
	var body userschemas.UpdateUserSchema
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (us *UserController) deleteUser(w http.ResponseWriter, r *http.Request) {}

func (us *UserController) updatePartialUser(w http.ResponseWriter, r *http.Request) {}
