package main

import(
	"fmt"
	"log"
	"time"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gocql/gocql"
)

var cassandraSession *gocql.Session

func main() {
	var cassandraSessionChannel chan *gocql.Session = make(chan *gocql.Session)

	go func() {
		CassandraLoop:
			for {
				cluster := gocql.NewCluster("cassandra")
				session, err := cluster.CreateSession()

				if err != nil {
					log.Println("Can't connect cassandra! Try Again", err)
					time.Sleep(500 * time.Millisecond)
				} else {
					cassandraSessionChannel <- session
					break CassandraLoop
				}
			}
	}()

	select {
		case cassandraSession = <- cassandraSessionChannel:
			log.Println("Cassandra connection established!")
		case <- time.After(10 * time.Second):
			log.Fatal("Can't connect to Cassadra! Quit by timeout!")
	}

	defer cassandraSession.Close()

	port := "8080"

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