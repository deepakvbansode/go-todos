package main

import (
	"context"
	"log"
	"net/http"

	"go-todos/config"
	"go-todos/database"
	"go-todos/handlers"
	"go-todos/middlewares"

	muxHandler "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	conf := config.GetConfig()
	ctx := context.TODO()
	db := database.ConnectDB(ctx, conf.Mongo)
	collection := db.Collection(conf.Mongo.Collection)
	client := &database.TodoClient{
		Ctx: ctx,
		Col: collection,
	}

	r := mux.NewRouter()
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(middlewares.SessionMiddleware)
	r.Use(middlewares.LoginMiddleware)

	r.HandleFunc("/todos", handlers.SearchTodo(client)).Methods(http.MethodGet)
	r.HandleFunc("/todos/{id}", handlers.GetTodo(client)).Methods(http.MethodGet)
	r.HandleFunc("/todos", handlers.InsertTodo(client)).Methods(http.MethodPost)
	r.HandleFunc("/todos/{id}", handlers.UpdateTodo(client)).Methods(http.MethodPatch)
	r.HandleFunc("/todos/{id}", handlers.DeleteTodo(client)).Methods(http.MethodDelete)
	log.Fatal(http.ListenAndServe(":8080", muxHandler.CORS(muxHandler.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), muxHandler.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), muxHandler.AllowedOrigins([]string{"*"}))(r)))
}
