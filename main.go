package main

import (
	"context"
	"net/http"
	"time"

	"github.com/friendsofgo/graphiql"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/mrtrom/go-graphql-example-api/config"
	"github.com/mrtrom/go-graphql-example-api/handler"
	"github.com/mrtrom/go-graphql-example-api/resolver"
	"github.com/mrtrom/go-graphql-example-api/service"
)

type strContextKey string

func main() {
	config := config.LoadConfig(".")

	ctx := context.Background()
	log := handler.NewLogger(config)

	userService := service.NewUserService()

	ctx = context.WithValue(ctx, strContextKey("config"), config)
	ctx = context.WithValue(ctx, strContextKey("logg"), log)
	ctx = context.WithValue(ctx, strContextKey("userService"), userService)

	var schema *graphql.Schema
	schemaString, err := handler.GetSchema()
	if err != nil {
		log.Fatal("Error getting schema: ", err)
	}
	schema = graphql.MustParseSchema(schemaString, &resolver.RootResolver{})

	mux := http.NewServeMux()

	graphiqlHandler, err := graphiql.NewGraphiqlHandler("/graphql")
	if err != nil {
		log.Panic(err)
	}

	mux.Handle("/graphql", handler.AddContext(ctx, &relay.Handler{Schema: schema}))
	mux.Handle("/graphiql", graphiqlHandler)

	// Configure the HTTP server.
	server := &http.Server{
		Addr:              config.Address,
		ReadHeaderTimeout: config.ReadHeaderTimeOut * time.Second,
		WriteTimeout:      config.WriteTimeout * time.Second,
		IdleTimeout:       config.IdleTimeout * time.Second,
		MaxHeaderBytes:    config.MaxHeaderBytes,
		Handler:           mux,
	}

	log.Fatal(server.ListenAndServe())
}
