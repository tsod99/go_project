package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/tsod99/go_project"
)

var port string = os.Getenv("PORT")

func init() {
	// initialize the database.
	conn, err := db.NewDB()
	if err != nil {
		log.Fatalf("Failed to get postgres connection %v", err)
	}
	defer conn.Close()

	if err := conn.InitialDatabase(); err != nil {
		log.Fatalf("Failed to initialize the databse %v", err)
	}
}

func main() {

	// http.HandleFunc("/list/users")
	// http.HandleFunc("/list/groups")
	// //
	// http.HandleFunc("/add/user")
	// http.HandleFunc("/add/group")
	// //
	// http.HandleFunc("/update/user")
	// http.HandleFunc("/update/group")
	// //
	// http.HandleFunc("/delete/user")
	// http.HandleFunc("/delete/group")

	log.Printf("listening on %s", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatal(err)
	}
}
