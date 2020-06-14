package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	h "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
	cfg "github.com/mrtrom/go-graphql-example-api/config"
	"github.com/mrtrom/go-graphql-example-api/db"
	"github.com/mrtrom/go-graphql-example-api/graph/generated"
	"github.com/mrtrom/go-graphql-example-api/graph/resolver"
	"github.com/mrtrom/go-graphql-example-api/handler"
	"github.com/mrtrom/go-graphql-example-api/service"
	"github.com/rs/cors"
)

func main() {
	config := cfg.LoadConfig(".")

	ctx := context.Background()
	log := handler.NewLogger(config)

	db := db.CreateConnetion(config, log)

	userService := service.NewUserService(db, log)
	chatService := service.NewChatService(db, log)

	ctx = context.WithValue(ctx, cfg.CTXConfig, config)
	ctx = context.WithValue(ctx, cfg.CTXLog, log)
	ctx = context.WithValue(ctx, cfg.CTXUserService, userService)
	ctx = context.WithValue(ctx, cfg.CTXChatService, chatService)

	router := chi.NewRouter()

	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	router.Route("/graphql", func(r chi.Router) {
		// r.Use(dataloaders.NewMiddleware(db)...)

		schema := generated.NewExecutableSchema(generated.Config{
			Resolvers: &resolver.Resolver{
				UserChannels: map[string]chan string{},
				Mutex:        sync.Mutex{},
			},
			// Directives: generated.DirectiveRoot{},
			// Complexity: generated.ComplexityRoot{},
		})

		srv := h.New(schema)
		// srv.Use(extension.FixedComplexityLimit(300))

		srv.AddTransport(transport.POST{})
		srv.AddTransport(transport.Websocket{
			KeepAlivePingInterval: 10 * time.Second,
			Upgrader: websocket.Upgrader{
				CheckOrigin: func(r *http.Request) bool {
					return true
				},
				ReadBufferSize:  1024,
				WriteBufferSize: 1024,
			},
		})
		srv.Use(extension.Introspection{})

		r.Handle("/", handler.AddContext(ctx, srv))
	})

	gqlPlayground := playground.Handler("api-gateway", "/graphql")
	router.Get("/", gqlPlayground)

	server := &http.Server{
		Addr:              config.Address,
		ReadHeaderTimeout: config.ReadHeaderTimeOut * time.Second,
		WriteTimeout:      config.WriteTimeout * time.Second,
		IdleTimeout:       config.IdleTimeout * time.Second,
		MaxHeaderBytes:    config.MaxHeaderBytes,
		Handler:           router,
	}

	log.Info(fmt.Sprintf("Connect to http://localhost%s/ for GraphQL playground", config.Address))

	log.Fatal(server.ListenAndServe())
}
