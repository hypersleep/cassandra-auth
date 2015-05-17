package main

import(
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// Config section

	port := "8080"
	cassandraHost := "cassandra"
	cassandraKeyspace := "auth"
	cassandraUsers := make(map[string]string)
	cassandraUsers["john@example.com"] = "password"

	// Cassandra keyspace section

	cassandraSession, err := newCassandraSession(cassandraHost, "")
	if err != nil {
		log.Fatal(err)
	}

	defer cassandraSession.Close()

	err = addKeyspace(cassandraKeyspace, cassandraSession)
	if err != nil {
		log.Fatal("Can't create keyspace:", err)
	}

	cassandraSession.Close()

	// Cassandra migrations section

	cassandraSession, err = newCassandraSession(cassandraHost, cassandraKeyspace)
	if err != nil {
		log.Fatal(err)
	}

	err = addUsers(cassandraUsers, cassandraSession)
	if err != nil {
		log.Fatal("Can't add users:", err)
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