package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/SaiHtetMyatHtut/potatoverse/graph"
	"github.com/SaiHtetMyatHtut/potatoverse/src/handlers"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.HandleFunc("/signin", handlers.SignIn)
	http.HandleFunc("POST /signup", handlers.SignUp)
	http.HandleFunc("/refresh", handlers.RefreshJwt)
	http.HandleFunc("/users", handlers.UserHandler)
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
