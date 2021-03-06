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
	cassandraUsers := map[string]string{
		"john@example.com": "password",
		"tom@example.com": "password",
		"mary@example.com": "password",
	}

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
		  HandlerFunc(registrationHandler)
	users.Methods("POST").
		  Path("/signin").
		  HandlerFunc(signInHandler)
	users.Methods("DELETE").
		  Path("/logout").
		  HandlerFunc(logOutHandler)
	users.Methods("POST").
		  Path("/check").
		  HandlerFunc(checkHandler)

	fmt.Println("Server running on port:", port)
	http.ListenAndServe(":" + port, r)
}
