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

func connectCassandra(cassandraHost string, cassandraSessionChannel chan *gocql.Session) {
	for {
		cluster := gocql.NewCluster(cassandraHost)
		session, err := cluster.CreateSession()

		if err == nil {
			cassandraSessionChannel <- session
			return
		}

		log.Println("Can't connect cassandra! Try again after 1 second!")
		time.Sleep(1 * time.Second)
	}
}

func main() {
	port := "8080"
	cassandraHost := "cassandra"

	var cassandraSessionChannel chan *gocql.Session = make(chan *gocql.Session)
	go connectCassandra(cassandraHost, cassandraSessionChannel)

	select {
		case cassandraSession = <- cassandraSessionChannel:
			log.Println("Cassandra connection established!")
		case <- time.After(10 * time.Second):
			log.Fatal("Can't connect to Cassadra! Quit by timeout!")
	}

	defer cassandraSession.Close()

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