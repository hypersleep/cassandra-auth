package main

import(
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Cassandra config section
const(
	cassandraHost string = "cassandra"
	cassandraKeyspace string = "auth"
)

func main() {

	// Config section
	port := "8080"
	cassandraUsers := make(map[string]string)
	cassandraUsers["john@example.com"] = "password"

	// Migrations section
	err := migrateCassandra(cassandraHost, cassandraKeyspace, cassandraUsers)
	if err != nil {
		log.Fatal("Migration failed! Reason:", err)
	}

	// HTTP Server section
	r := mux.NewRouter()

	users := r.PathPrefix("/users").Subrouter()

	users.Methods("POST").
		  Path("/create").
		  HandlerFunc(RegistrationHandler)
	users.Methods("POST").
		  Path("/signin").
		  HandlerFunc(SignInHandler)
	users.Methods("DELETE").
		  Path("/logout").
		  HandlerFunc(LogOutHandler)
	users.Methods("POST").
		  Path("/check").
		  HandlerFunc(CheckHandler)

	fmt.Println("Server running on port:", port)
	http.ListenAndServe(":" + port, r)
}