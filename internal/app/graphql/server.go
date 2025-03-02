package graphql

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/ThePotatoVerse/internal/app/graphql/resolver"
	"github.com/ThePotatoVerse/internal/app/middleware"
	"github.com/ThePotatoVerse/internal/app/service"
	"github.com/ThePotatoVerse/pkg/logger"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

// Server represents a GraphQL server
type Server struct {
	schema      *graphql.Schema
	log         logger.Logger
	authService service.AuthService
}

// NewServer creates a new GraphQL server
func NewServer(log logger.Logger, authService service.AuthService) (*Server, error) {
	// Load schema
	schemaString, err := loadSchema()
	if err != nil {
		return nil, err
	}

	// Create resolvers
	authResolver := resolver.NewAuthResolver(authService, log)
	rootResolver := &RootResolver{
		authResolver: authResolver,
	}

	// Parse schema
	schema := graphql.MustParseSchema(schemaString, rootResolver)

	return &Server{
		schema:      schema,
		log:         log,
		authService: authService,
	}, nil
}

// Handler returns an HTTP handler for the GraphQL server
func (s *Server) Handler() http.Handler {
	// Create GraphQL handler
	handler := &relay.Handler{Schema: s.schema}

	// Add authentication middleware
	authMiddleware := middleware.GraphQLAuthMiddleware(s.authService, s.log)

	return authMiddleware(handler)
}

// RootResolver combines all resolvers
type RootResolver struct {
	authResolver *resolver.AuthResolver
}

// Register delegates to the auth resolver
func (r *RootResolver) Register(args struct{ Input resolver.RegisterInput }) (*resolver.AuthResponse, error) {
	return r.authResolver.Register(nil, args.Input)
}

// Login delegates to the auth resolver
func (r *RootResolver) Login(args struct{ Input resolver.LoginInput }) (*resolver.AuthResponse, error) {
	return r.authResolver.Login(nil, args.Input)
}

// Me delegates to the auth resolver
func (r *RootResolver) Me() (*resolver.UserResolver, error) {
	return r.authResolver.Me(nil)
}

// loadSchema loads the GraphQL schema from files
func loadSchema() (string, error) {
	// Get schema directory
	schemaDir := filepath.Join("internal", "app", "graphql", "schema")

	// Read schema files
	var schema string
	err := filepath.Walk(schemaDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".graphql" {
			data, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			schema += string(data) + "\n"
		}
		return nil
	})

	return schema, err
}
