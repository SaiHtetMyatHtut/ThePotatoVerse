package server

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/SaiHtetMyatHtut/potatoverse/configs"
	"github.com/SaiHtetMyatHtut/potatoverse/graph"
	"github.com/SaiHtetMyatHtut/potatoverse/src/controllers"
	"github.com/SaiHtetMyatHtut/potatoverse/src/handlers"
	"go.uber.org/dig"
)

type Server struct {
	UserController controllers.UserController
}

type ServerDependencies struct {
	dig.In

	UserContoller controllers.UserController `name:"UserController"`
}

func NewServer(deps ServerDependencies) {
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	app := Server{
		UserController: deps.UserContoller,
	}

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.HandleFunc("/signin", handlers.SignIn)
	http.HandleFunc("/signup", handlers.SignUp)
	http.HandleFunc("/refresh", handlers.RefreshJwt)
	http.HandleFunc("/users", app.UserController.Exec)
	http.HandleFunc("/users/{id}", app.UserController.Exec)
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", configs.Env.Server.Port)
	log.Fatal(http.ListenAndServe(":"+configs.Env.Server.Port, nil))
}
