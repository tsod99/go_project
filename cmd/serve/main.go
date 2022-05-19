package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/tsod99/go_project/db"
	_ "github.com/tsod99/go_project/docs"
	"github.com/tsod99/go_project/handlers"
)

var port string = os.Getenv("PORT")

// this is the port for the swagger documentation
const documentationPort string = "6060"

func init() {
	// initialize the database.
	conn, err := db.NewDB()
	if err != nil {
		log.Fatalf("Failed to get postgres connection %v", err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("failed to close connection %v", err)
		}
	}()

	if err := conn.InitialDatabase(); err != nil {
		log.Fatalf("Failed to initialize the database %v", err)
	}
}

// @title api
// @version 1.0
// @description api swagger documentation

// NOTE: for the port use the same port you use for the variable `port`

// @host localhost:8080
// @BasePath /

func main() {

	http.HandleFunc("/list/users", handlers.CorsMiddleware(handlers.HandleListUsers))
	http.HandleFunc("/list/groups", handlers.CorsMiddleware(handlers.HandleListGroups))
	//
	http.HandleFunc("/add/user", handlers.CorsMiddleware(handlers.HandleAddUser))
	http.HandleFunc("/add/group", handlers.CorsMiddleware(handlers.HandleAddGroup))
	//
	http.HandleFunc("/update/user", handlers.CorsMiddleware(handlers.HandleUpdateUser))
	http.HandleFunc("/update/group", handlers.CorsMiddleware(handlers.HandleUpdateGroup))
	//
	http.HandleFunc("/delete/user", handlers.CorsMiddleware(handlers.HandleDeleteUser))
	http.HandleFunc("/delete/group", handlers.CorsMiddleware(handlers.HandleDeleteGroup))

	{
		swaggerRouter := chi.NewRouter()
		swaggerRouter.Get("/swagger/*", httpSwagger.Handler(
			httpSwagger.URL(fmt.Sprintf("http://localhost:%s/swagger/doc.json", documentationPort)),
		))

		go func() {
			http.ListenAndServe(fmt.Sprintf(":%s", documentationPort), swaggerRouter)
		}()
	}

	log.Printf("listening on %s", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatal(err)
	}
}
