package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/SaiHtetMyatHtut/potatoverse/graph/model"
	"github.com/SaiHtetMyatHtut/potatoverse/models"
	"github.com/SaiHtetMyatHtut/potatoverse/repo"
	"github.com/SaiHtetMyatHtut/potatoverse/utils"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	newHashPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		panic(err)
	}
	newUser := models.User{
		Username:       input.Username,
		HashedPassword: newHashPassword,
		CreatedAt:      time.Now().UTC(),
		LastLogin:      time.Now().UTC(),
	}
	user, _ := repo.Insert(ctx, newUser)

	resUser := &model.User{
		ID:        strconv.FormatInt(user.ID, 10), // Convert int64 to string
		Username:  user.Username,
		CreatedAt: user.CreatedAt.Format(time.RFC3339), // Convert time.Time to string
		LastLogin: user.LastLogin.Format(time.RFC3339),
	}

	return resUser, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	users, _ := repo.ReadAll(ctx)
	var resUsers []*model.User
	for _, u := range users {
		resUser := &model.User{
			ID:        strconv.FormatInt(u.ID, 10), // Convert int64 to string
			Username:  u.Username,
			CreatedAt: u.CreatedAt.Format(time.RFC3339), // Convert time.Time to string
			LastLogin: u.LastLogin.Format(time.RFC3339),
		}
		resUsers = append(resUsers, resUser)
	}
	return resUsers, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver {
	log.Println("QueryResolver")
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
