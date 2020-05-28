package main

import (
	"context"
	"fmt"
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
	var (
		addr              = ":8080"
		readHeaderTimeout = 1 * time.Second
		writeTimeout      = 10 * time.Second
		idleTimeout       = 90 * time.Second
		maxHeaderBytes    = http.DefaultMaxHeaderBytes
	)

	config := config.LoadConfig(".")

	ctx := context.Background()
	log := handler.NewLogger(config)
	userService := service.NewUserService()

	ctx = context.WithValue(ctx, strContextKey("config"), config)
	ctx = context.WithValue(ctx, strContextKey("log"), log)
	ctx = context.WithValue(ctx, strContextKey("userService"), userService)

	var schema *graphql.Schema
	schemaString, err := handler.GetSchema()
	if err != nil {
		fmt.Printf("Error getting schema: %s", err)
	}
	schema = graphql.MustParseSchema(schemaString, &resolver.RootResolver{})

	mux := http.NewServeMux()

	graphiqlHandler, err := graphiql.NewGraphiqlHandler("/graphql")
	if err != nil {
		panic(err)
	}

	mux.Handle("/graphql", handler.AddContext(ctx, &relay.Handler{Schema: schema}))
	mux.Handle("/graphiql", graphiqlHandler)

	// Configure the HTTP server.
	server := &http.Server{
		Addr:              addr,
		Handler:           mux,
		ReadHeaderTimeout: readHeaderTimeout,
		WriteTimeout:      writeTimeout,
		IdleTimeout:       idleTimeout,
		MaxHeaderBytes:    maxHeaderBytes,
	}

	log.Fatal(server.ListenAndServe())
}
